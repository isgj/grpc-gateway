// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: examples/internal/proto/examplepb/generate_unbound_methods.proto

/*
Package examplepb is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package examplepb

import (
	"context"
	"errors"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var (
	_ codes.Code
	_ io.Reader
	_ status.Status
	_ = errors.New
	_ = runtime.String
	_ = utilities.NewDoubleArray
	_ = metadata.Join
)

func request_GenerateUnboundMethodsEchoService_Echo_0(ctx context.Context, marshaler runtime.Marshaler, client GenerateUnboundMethodsEchoServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq GenerateUnboundMethodsSimpleMessage
		metadata runtime.ServerMetadata
	)
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && !errors.Is(err, io.EOF) {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if req.Body != nil {
		_, _ = io.Copy(io.Discard, req.Body)
	}
	msg, err := client.Echo(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

func local_request_GenerateUnboundMethodsEchoService_Echo_0(ctx context.Context, marshaler runtime.Marshaler, server GenerateUnboundMethodsEchoServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq GenerateUnboundMethodsSimpleMessage
		metadata runtime.ServerMetadata
	)
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && !errors.Is(err, io.EOF) {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := server.Echo(ctx, &protoReq)
	return msg, metadata, err
}

func request_GenerateUnboundMethodsEchoService_EchoBody_0(ctx context.Context, marshaler runtime.Marshaler, client GenerateUnboundMethodsEchoServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq GenerateUnboundMethodsSimpleMessage
		metadata runtime.ServerMetadata
	)
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && !errors.Is(err, io.EOF) {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if req.Body != nil {
		_, _ = io.Copy(io.Discard, req.Body)
	}
	msg, err := client.EchoBody(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

func local_request_GenerateUnboundMethodsEchoService_EchoBody_0(ctx context.Context, marshaler runtime.Marshaler, server GenerateUnboundMethodsEchoServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq GenerateUnboundMethodsSimpleMessage
		metadata runtime.ServerMetadata
	)
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && !errors.Is(err, io.EOF) {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := server.EchoBody(ctx, &protoReq)
	return msg, metadata, err
}

func request_GenerateUnboundMethodsEchoService_EchoDelete_0(ctx context.Context, marshaler runtime.Marshaler, client GenerateUnboundMethodsEchoServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq GenerateUnboundMethodsSimpleMessage
		metadata runtime.ServerMetadata
	)
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && !errors.Is(err, io.EOF) {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if req.Body != nil {
		_, _ = io.Copy(io.Discard, req.Body)
	}
	msg, err := client.EchoDelete(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

func local_request_GenerateUnboundMethodsEchoService_EchoDelete_0(ctx context.Context, marshaler runtime.Marshaler, server GenerateUnboundMethodsEchoServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq GenerateUnboundMethodsSimpleMessage
		metadata runtime.ServerMetadata
	)
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && !errors.Is(err, io.EOF) {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := server.EchoDelete(ctx, &protoReq)
	return msg, metadata, err
}

// RegisterGenerateUnboundMethodsEchoServiceHandlerServer registers the http handlers for service GenerateUnboundMethodsEchoService to "mux".
// UnaryRPC     :call GenerateUnboundMethodsEchoServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterGenerateUnboundMethodsEchoServiceHandlerFromEndpoint instead.
// GRPC interceptors will not work for this type of registration. To use interceptors, you must use the "runtime.WithMiddlewares" option in the "runtime.NewServeMux" call.
func RegisterGenerateUnboundMethodsEchoServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server GenerateUnboundMethodsEchoServiceServer) error {
	mux.Handle(http.MethodPost, pattern_GenerateUnboundMethodsEchoService_Echo_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/grpc.gateway.examples.internal.proto.examplepb.GenerateUnboundMethodsEchoService/Echo", runtime.WithHTTPPathPattern("/grpc.gateway.examples.internal.proto.examplepb.GenerateUnboundMethodsEchoService/Echo"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_GenerateUnboundMethodsEchoService_Echo_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_GenerateUnboundMethodsEchoService_Echo_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodPost, pattern_GenerateUnboundMethodsEchoService_EchoBody_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/grpc.gateway.examples.internal.proto.examplepb.GenerateUnboundMethodsEchoService/EchoBody", runtime.WithHTTPPathPattern("/grpc.gateway.examples.internal.proto.examplepb.GenerateUnboundMethodsEchoService/EchoBody"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_GenerateUnboundMethodsEchoService_EchoBody_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_GenerateUnboundMethodsEchoService_EchoBody_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodPost, pattern_GenerateUnboundMethodsEchoService_EchoDelete_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/grpc.gateway.examples.internal.proto.examplepb.GenerateUnboundMethodsEchoService/EchoDelete", runtime.WithHTTPPathPattern("/grpc.gateway.examples.internal.proto.examplepb.GenerateUnboundMethodsEchoService/EchoDelete"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_GenerateUnboundMethodsEchoService_EchoDelete_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_GenerateUnboundMethodsEchoService_EchoDelete_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})

	return nil
}

// RegisterGenerateUnboundMethodsEchoServiceHandlerFromEndpoint is same as RegisterGenerateUnboundMethodsEchoServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterGenerateUnboundMethodsEchoServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.NewClient(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()
	return RegisterGenerateUnboundMethodsEchoServiceHandler(ctx, mux, conn)
}

// RegisterGenerateUnboundMethodsEchoServiceHandler registers the http handlers for service GenerateUnboundMethodsEchoService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterGenerateUnboundMethodsEchoServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterGenerateUnboundMethodsEchoServiceHandlerClient(ctx, mux, NewGenerateUnboundMethodsEchoServiceClient(conn))
}

// RegisterGenerateUnboundMethodsEchoServiceHandlerClient registers the http handlers for service GenerateUnboundMethodsEchoService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "GenerateUnboundMethodsEchoServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "GenerateUnboundMethodsEchoServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "GenerateUnboundMethodsEchoServiceClient" to call the correct interceptors. This client ignores the HTTP middlewares.
func RegisterGenerateUnboundMethodsEchoServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client GenerateUnboundMethodsEchoServiceClient) error {
	mux.Handle(http.MethodPost, pattern_GenerateUnboundMethodsEchoService_Echo_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/grpc.gateway.examples.internal.proto.examplepb.GenerateUnboundMethodsEchoService/Echo", runtime.WithHTTPPathPattern("/grpc.gateway.examples.internal.proto.examplepb.GenerateUnboundMethodsEchoService/Echo"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_GenerateUnboundMethodsEchoService_Echo_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_GenerateUnboundMethodsEchoService_Echo_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodPost, pattern_GenerateUnboundMethodsEchoService_EchoBody_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/grpc.gateway.examples.internal.proto.examplepb.GenerateUnboundMethodsEchoService/EchoBody", runtime.WithHTTPPathPattern("/grpc.gateway.examples.internal.proto.examplepb.GenerateUnboundMethodsEchoService/EchoBody"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_GenerateUnboundMethodsEchoService_EchoBody_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_GenerateUnboundMethodsEchoService_EchoBody_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodPost, pattern_GenerateUnboundMethodsEchoService_EchoDelete_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/grpc.gateway.examples.internal.proto.examplepb.GenerateUnboundMethodsEchoService/EchoDelete", runtime.WithHTTPPathPattern("/grpc.gateway.examples.internal.proto.examplepb.GenerateUnboundMethodsEchoService/EchoDelete"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_GenerateUnboundMethodsEchoService_EchoDelete_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_GenerateUnboundMethodsEchoService_EchoDelete_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	return nil
}

var (
	pattern_GenerateUnboundMethodsEchoService_Echo_0       = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"grpc.gateway.examples.internal.proto.examplepb.GenerateUnboundMethodsEchoService", "Echo"}, ""))
	pattern_GenerateUnboundMethodsEchoService_EchoBody_0   = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"grpc.gateway.examples.internal.proto.examplepb.GenerateUnboundMethodsEchoService", "EchoBody"}, ""))
	pattern_GenerateUnboundMethodsEchoService_EchoDelete_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"grpc.gateway.examples.internal.proto.examplepb.GenerateUnboundMethodsEchoService", "EchoDelete"}, ""))
)

var (
	forward_GenerateUnboundMethodsEchoService_Echo_0       = runtime.ForwardResponseMessage
	forward_GenerateUnboundMethodsEchoService_EchoBody_0   = runtime.ForwardResponseMessage
	forward_GenerateUnboundMethodsEchoService_EchoDelete_0 = runtime.ForwardResponseMessage
)
