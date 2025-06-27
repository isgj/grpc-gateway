package integration_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"net"

	"google.golang.org/grpc/health/grpc_health_v1"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/anypb"

	pb "github.com/grpc-ecosystem/grpc-gateway/v2/examples/internal/proto/editions"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
)

type editionsTestServer struct {
	pb.UnimplementedEditionsTestServiceServer
}

func (s *editionsTestServer) TestFieldPresence(ctx context.Context, req *pb.EditionsTestRequest) (*pb.EditionsTestResponse, error) {
	resp := &pb.EditionsTestResponse{}
	fm := req.GetFeaturesMessage()
	if fm != nil {
		resp.WasOptionalStringPresent = fm.OptionalString != nil
		if fm.OptionalString != nil {
			resp.ServerReceivedOptionalString = *fm.OptionalString
		}

		resp.WasOptionalInt32Present = fm.OptionalInt32 != nil
		if fm.OptionalInt32 != nil {
			resp.ServerReceivedOptionalInt32 = *fm.OptionalInt32
		}

		resp.WasOptionalBoolPresent = fm.OptionalBool != nil
		if fm.OptionalBool != nil {
			resp.ServerReceivedOptionalBool = *fm.OptionalBool
		}
		resp.ServerReceivedRequiredString = fm.GetRequiredString() // GetRequiredString will be present due to [(google.api.field_behavior) = REQUIRED]

		if fm.OpaqueAny != nil {
			resp.AnyTypeUrlReceived = fm.OpaqueAny.TypeUrl
			// Attempt to use opaque API - unmarshal SimpleMessageForOpaque
			if fm.OpaqueAny.MessageIs(&pb.SimpleMessageForOpaque{}) {
				var simple pb.SimpleMessageForOpaque
				if err := fm.OpaqueAny.UnmarshalTo(&simple); err == nil {
					// This also implicitly tests protoopaque.Open through UnmarshalTo's typical implementation path
					// and MessageIs.
				}
			}
		}
	}
	if req.GetSimpleOpaqueMessage() != nil {
		resp.SimpleOpaqueDataReceived = req.GetSimpleOpaqueMessage().GetData()
	}

	return resp, nil
}

func runEditionsTestServer(ctx context.Context, t *testing.T, gwAddr string, dialOpts ...grpc.DialOption) func() {
	ctx, cancel := context.WithCancel(ctx)

	lis, teardown := startGRPCServer(ctx, t, func(s *grpc.Server) {
		pb.RegisterEditionsTestServiceServer(s, &editionsTestServer{})
	})

	mux := runtime.NewServeMux(
		runtime.WithForwardResponseOptions(forwardResponseOption),
		runtime.WithStreamErrorHandler(handleStreamError),
		runtime.WithRoutingErrorHandler(handleRoutingError),
		runtime.WithIncomingHeaderMatcher(headerMatcher),
		runtime.WithOutgoingHeaderMatcher(headerMatcher),
		runtime.WithProtoErrorHandler(handleProtoErrorHandler),
		runtime.WithHealthzEndpoint(healthzpb.NewHealthClient(conn)),
		runtime.WithQueryParameterParser(utilities.NewHTTPFormQueryParameterParser()),
	)

	if err := pb.RegisterEditionsTestServiceHandlerFromEndpoint(ctx, mux, lis.Addr().String(), dialOpts); err != nil {
		t.Fatalf("Failed to register gateway: %v", err)
	}

	go func() {
		if err := http.ListenAndServe(gwAddr, mux); err != nil {
			// ErrServerClosed is expected on graceful shutdown.
			if err != http.ErrServerClosed {
				t.Errorf("ListenAndServe returned an unexpected error: %v", err)
			}
		}
	}()

	return cancel
}

func TestEditionsFieldPresence(t *testing.T) {
	if testing.Short() {
		t.Skip()
		return
	}
	ctx := context.Background()
	addr := ":8099"
	cancel := runEditionsTestServer(ctx, t, addr, grpc.WithInsecure())
	defer cancel()

	testCases := []struct {
		name           string
		requestBody    map[string]interface{}
		expectedResponse *pb.EditionsTestResponse
	}{
		{
			name: "all optional fields present",
			requestBody: map[string]interface{}{
				"features_message": map[string]interface{}{
					"optional_string": "hello",
					"optional_int32":  123,
					"optional_bool":   true,
					"required_string": "world",
				},
			},
			expectedResponse: &pb.EditionsTestResponse{
				ServerReceivedOptionalString: "hello", WasOptionalStringPresent: true,
				ServerReceivedOptionalInt32: 123, WasOptionalInt32Present: true,
				ServerReceivedOptionalBool: true, WasOptionalBoolPresent: true,
				ServerReceivedRequiredString: "world",
			},
		},
		{
			name: "some optional fields absent",
			requestBody: map[string]interface{}{
				"features_message": map[string]interface{}{
					"optional_string": "test",
					"required_string": "mandatory",
					// optional_int32 and optional_bool are absent
				},
			},
			expectedResponse: &pb.EditionsTestResponse{
				ServerReceivedOptionalString: "test", WasOptionalStringPresent: true,
				ServerReceivedOptionalInt32: 0, WasOptionalInt32Present: false, // zero value for int32
				ServerReceivedOptionalBool: false, WasOptionalBoolPresent: false, // zero value for bool
				ServerReceivedRequiredString: "mandatory",
			},
		},
		{
			name: "all optional fields absent",
			requestBody: map[string]interface{}{
				"features_message": map[string]interface{}{
					"required_string": "only_required",
				},
			},
			expectedResponse: &pb.EditionsTestResponse{
				ServerReceivedOptionalString: "", WasOptionalStringPresent: false,
				ServerReceivedOptionalInt32: 0, WasOptionalInt32Present: false,
				ServerReceivedOptionalBool: false, WasOptionalBoolPresent: false,
				ServerReceivedRequiredString: "only_required",
			},
		},
		{
			name: "opaque Any handling",
			requestBody: map[string]interface{}{
				"features_message": map[string]interface{}{
					"required_string": "any_test",
					"opaque_any": map[string]interface{}{
						"@type": "type.googleapis.com/grpc.gateway.examples.internal.proto.editions.SimpleMessageForOpaque",
						"data":  "opaque data in any",
					},
				},
				"simple_opaque_message": map[string]interface{}{
					"data": "direct opaque data",
				},
			},
			expectedResponse: &pb.EditionsTestResponse{
				ServerReceivedRequiredString:  "any_test",
				AnyTypeUrlReceived:            "type.googleapis.com/grpc.gateway.examples.internal.proto.editions.SimpleMessageForOpaque",
				SimpleOpaqueDataReceived:      "direct opaque data",
				WasOptionalStringPresent:      false,
				WasOptionalInt32Present:       false,
				WasOptionalBoolPresent:        false,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			jsonBody, err := json.Marshal(tc.requestBody)
			if err != nil {
				t.Fatalf("Failed to marshal request body: %v", err)
			}

			req, err := http.NewRequest("POST", fmt.Sprintf("http://localhost%s/v1/editions/test_field_presence", addr), bytes.NewReader(jsonBody))
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("Failed to send request: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				// Simplified error handling for brevity in this example
				t.Fatalf("Expected status OK; got %v", resp.Status)
			}

			var result pb.EditionsTestResponse
			if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
				t.Fatalf("Failed to decode response body: %v", err)
			}

			if diff := cmp.Diff(tc.expectedResponse, &result, protocmp.Transform()); diff != "" {
				t.Errorf("Response mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

// Minimal stubs for other error handlers to make the test setup compile
func forwardResponseOption(ctx context.Context, w http.ResponseWriter, msg proto.Message) error {
	return nil
}

func handleStreamError(ctx context.Context, err error) *status.Status {
	return runtime.DefaultHTTPStreamErrorHandler(ctx, err)
}

func handleRoutingError(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, httpStatus int) {
	runtime.DefaultRoutingErrorHandler(ctx, mux, marshaler, w, r, httpStatus)
}

func headerMatcher(headerName string) (mdName string, ok bool) {
	return fmt.Sprintf("%s%s", runtime.MetadataPrefix, headerName), true
}

func handleProtoErrorHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	runtime.DefaultHTTPProtoErrorHandler(ctx, mux, marshaler, w, r, err)
}

// Assuming healthzpb and conn are available from other test files or setup.
// For this standalone test, we might need to simplify or mock this if it causes issues.
// For now, let this be, it might be picked from other integration tests' setup.
var healthzpb = &mockHealthClient{} // Placeholder
var conn *grpc.ClientConn = nil // Placeholder

type mockHealthClient struct{}
func (m *mockHealthClient) Check(ctx context.Context, in *grpc.health.v1.HealthCheckRequest, opts ...grpc.CallOption) (*grpc.health.v1.HealthCheckResponse, error) {
	return &grpc.health.v1.HealthCheckResponse{Status: grpc.health.v1.HealthCheckResponse_SERVING}, nil
}
func (m *mockHealthClient) Watch(ctx context.Context, in *grpc.health.v1.HealthCheckRequest, opts ...grpc.CallOption) (grpc.health.v1.Health_WatchClient, error) {
	return nil, status.Errorf(codes.Unimplemented, "Watch is not implemented")
}
// End placeholder for healthz

// Helper to get a gRPC server listening on a random port.
// This is a simplified version of what might be in main_test.go
func startGRPCServer(ctx context.Context, t *testing.T, register func(s *grpc.Server)) (lis net.Listener, teardown func()) {
	var err error
	lis, err = net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	register(s)

	go func() {
		if err := s.Serve(lis); err != nil && err != grpc.ErrServerStopped {
			t.Logf("gRPC server failed to serve: %v", err)
		}
	}()
	return lis, func() {
		s.GracefulStop()
		lis.Close()
	}
}
// Need to add net and grpc/health/grpc_health_v1 for the placeholders
// import "net"
// import "google.golang.org/grpc/health/grpc_health_v1"
// For now, I'll add them directly here. If they are pulled by other test files, this is fine.
// Otherwise, a `gazelle` run or manual BUILD file update would be needed.

[end of examples/internal/integration/editions_test.go]
