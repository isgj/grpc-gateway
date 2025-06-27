package editions

import "google.golang.org/protobuf/types/descriptorpb"

// MinSupportedEdition is the minimum protobuf edition supported by the gateway.
const MinSupportedEdition = descriptorpb.Edition_EDITION_2023

// MaxSupportedEdition is the maximum protobuf edition supported by the gateway.
const MaxSupportedEdition = descriptorpb.Edition_EDITION_2023

// IsEditionSupported checks if the given edition is within the supported range.
func IsEditionSupported(edition descriptorpb.Edition) bool {
	if edition == descriptorpb.Edition_EDITION_UNKNOWN {
		// Assuming EDITION_UNKNOWN might be treated as implicitly supported
		// or should be handled by callers based on syntax (proto2/proto3).
		// For explicit edition checks, UNKNOWN is likely not "supported" in terms of new features.
		return false
	}
	return edition >= MinSupportedEdition && edition <= MaxSupportedEdition
}
