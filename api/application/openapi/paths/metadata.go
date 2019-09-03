package paths

import (
	"github.com/johnrichardrinehart/go2openapi"
)

func init() {
	// minimum information needed for all subsequent endpoints and methods
	var operation go2openapi.Operation = go2openapi.Operation{
		Responses: go2openapi.Responses{
			"200": go2openapi.Response{
				Description: "Successful query",
				Content: map[string]go2openapi.MediaType{
					"application/JSON": go2openapi.MediaType{
						Schema: &responseSchema,
					},
				},
			},
		},
		Parameters: go2openapi.Parameters{
			go2openapi.Parameter{
				Name:        "q",
				In:          "query",
				Description: "MongoDB-compatible JSON query document",
				Required:    go2openapi.Ptrue,
				Schema:      querySchema,
			},
		},
		Tags: []string{"metadata"},
	}

	// Parameters
	var streamNameMetadataParameters = go2openapi.Parameters{
		homeIdParameter,
		smartModuleIDParameter,
		streamNameParameter,
	}

	// /homes/{homeId}/smartModules/{smartModuleId}/streams/{streamName}/metadata
	// GET
	var streamNameMetadataGet = operation
	// PATCH
	var streamNameMetadataPatch = operation
	// GET

	// PathItem
	var streamNameMetadata go2openapi.PathItem = go2openapi.PathItem{
		Summary:     "Query/Modify a given stream's metadata",
		Description: "This endpoint allows application end user's to obtain and modify the metadata for a given stream.",
		Get:         &streamNameMetadataGet,
		Patch:       &streamNameMetadataPatch,
		Parameters:  streamNameMetadataParameters,
	}

	var smartModuleIDMetadata go2openapi.PathItem = streamNameMetadata

	var homeIDMetadata go2openapi.PathItem = streamNameMetadata

	smartModuleIDMetadata.Patch.Parameters = nil
	homeIDMetadata.Patch.Parameters = nil
	Paths.AddPath("/homes/{homeId}/smartModules/{smartModuleId}/streams/{streamName}/metadata", streamNameMetadata)
	Paths.AddPath("/homes/{homeId}/smartModules/{smartModuleId}/metadata", smartModuleIDMetadata)
	Paths.AddPath("/homes/{homeId}/metadata", homeIDMetadata)
}
