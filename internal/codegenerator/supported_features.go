package codegenerator

import (
	"reflect"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

func supportedCodeGeneratorFeatures() uint64 {
	// Enable support for optional keyword in proto3 and support for editions.
	return uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL) | uint64(pluginpb.CodeGeneratorResponse_FEATURE_SUPPORTS_EDITIONS)
}

// SetSupportedFeaturesOnPluginGen sets supported proto3 features and attempts to set edition range.
func SetSupportedFeaturesOnPluginGen(gen *protogen.Plugin) {
	currentFeatures := supportedCodeGeneratorFeatures()
	gen.SupportedFeatures = currentFeatures

	if (currentFeatures & uint64(pluginpb.CodeGeneratorResponse_FEATURE_SUPPORTS_EDITIONS)) != 0 {
		// Attempt to set Min/Max editions via reflection.
		// This is to work around potential mismatches in protogen library versions
		// used by 'go install' for the plugin vs. the main module's declared version.
		// The actual fields on protogen.Plugin are descriptorpb.Edition, not *int32.
		val := reflect.ValueOf(gen).Elem() // gen is a pointer

		// Using MinSupportedEdition from where it would be defined, assuming internal/editions
		// For this direct modification, hardcoding for now or assuming available.
		// Let's use the direct value for now to simplify, actual value comes from editions package.
		minEditionVal := descriptorpb.Edition_EDITION_2023
		maxEditionVal := descriptorpb.Edition_EDITION_2023

		minEditionField := val.FieldByName("MinimumEdition")
		if minEditionField.IsValid() && minEditionField.CanSet() && minEditionField.Type() == reflect.TypeOf(minEditionVal) {
			minEditionField.Set(reflect.ValueOf(minEditionVal))
		}

		maxEditionField := val.FieldByName("MaximumEdition")
		if maxEditionField.IsValid() && maxEditionField.CanSet() && maxEditionField.Type() == reflect.TypeOf(maxEditionVal) {
			maxEditionField.Set(reflect.ValueOf(maxEditionVal))
		}
	}
}

// SetSupportedFeaturesOnCodeGeneratorResponse sets supported proto3 features
// and edition support range on pluginpb.CodeGeneratorResponse.
func SetSupportedFeaturesOnCodeGeneratorResponse(resp *pluginpb.CodeGeneratorResponse) {
	sf := supportedCodeGeneratorFeatures()
	minEd := int32(descriptorpb.Edition_EDITION_2023)
	maxEd := int32(descriptorpb.Edition_EDITION_2023)

	resp.SupportedFeatures = &sf
	if (sf & uint64(pluginpb.CodeGeneratorResponse_FEATURE_SUPPORTS_EDITIONS)) != 0 {
		resp.MinimumEdition = &minEd
		resp.MaximumEdition = &maxEd
	}
}
