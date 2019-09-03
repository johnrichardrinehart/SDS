package paths

import (
	"github.com/johnrichardrinehart/go2openapi"
)

// /homes/{homeId}/smartModules/{smartModuleId}/streams/{streamName}/metadata
// GET
var streamNameMetadataGet go2openapi.Operation = go2openapi.Operation{
	Responses: go2openapi.Responses{
		"200": go2openapi.Response{
			Description: "Success!",
		},
	},
	Tags: []string{"metadata"},
}

// PATCH
var streamNameMetadataPatch go2openapi.Operation = streamNameMetadataGet

// Parameters
var streamNameMetadataParameters = go2openapi.Parameters{
	homeIdParameter,
	smartModuleIDParameter,
	streamNameParameter,
}

// PathItem
var streamNameMetadata go2openapi.PathItem = go2openapi.PathItem{
	Summary:     "Query/Modify a given stream's metadata",
	Description: "This endpoint allows application end user's to obtain and modify the metadata for a given stream.",
	Get:         &streamNameMetadataGet,
	Patch:       &streamNameMetadataPatch,
	Parameters:  &streamNameMetadataParameters,
}

var smartModuleIDMetadata go2openapi.PathItem = streamNameMetadata

var homeIDMetadata go2openapi.PathItem = streamNameMetadata

func init() {
	Paths.AddPath("/homes/{homeId}/smartModules/{smartModuleId}/streams/{streamName}/metadata", streamNameMetadata)
	Paths.AddPath("/homes/{homeId}/smartModules/{smartModuleId}/metadata", smartModuleIDMetadata)
	Paths.AddPath("/homes/{homeId}/metadata", homeIDMetadata)
}
