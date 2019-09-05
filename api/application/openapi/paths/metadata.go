package paths

import (
	"github.com/johnrichardrinehart/go2openapi"
)

func init() {
	//####   General/Recycled Values   ####//

	// HTTP Request/Response Structure
	var response200 = go2openapi.Response{
		Description: "Successful query",
		Content: map[string]go2openapi.MediaType{
			"application/JSON": go2openapi.MediaType{
				Schema: &responseSchema,
			},
		},
	}

	var response401 = go2openapi.Response{
		Description: "Authorization rejected/required.",
	}

	// var response406 = go2openapi.Response{
	// 	Description: "Invalid metadata requested to update (e.g., read-only properties).",
	// }

	var operation go2openapi.Operation = go2openapi.Operation{
		Responses: go2openapi.Responses{
			"200": &response200,
			"401": &response401,
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
	var pathParameters = go2openapi.Parameters{
		paramAuthorizationHeader,
		paramHomeIDPath,
		paramSmartModuleIDPath,
		paramStreamNamePath,
	}

	// GET
	var get = operation
	get.Parameters.AddParameters(
		paramLimitQuery,
	)

	//####   /homes/{homeId}/smartModules/{smartModuleId}/streams/{streamName}/metadata   ####//
	var streamNameMetadata go2openapi.PathItem = go2openapi.PathItem{
		Summary:     "Query/Modify a given stream's metadata",
		Description: "This endpoint allows application end user's to obtain and modify the metadata for a given stream.",
		Get:         &get,
		Parameters:  pathParameters,
	}

	//####   /homes/{homeId}/smartModules/{smartModuleId}/streams/{streamName}/metadata/{packetIndex}   ####//
	pathPacketIndex := streamNameMetadata
	pathPacketIndex.Get = nil
	pathPacketIndex.AddParameters(go2openapi.Parameter{
		Name:     "packetIndex",
		In:       "query",
		Required: go2openapi.Ptrue,
		Example:  1020,
	})
	pathPacketIndex.Delete = &go2openapi.Operation{
		Tags: []string{"metadata"},
	}
	pathPacketIndex.Delete.RequestBody = &go2openapi.RequestBody{
		Content: putRequest,
	}

	var delete200 = &go2openapi.Response{
		Description: "Packet <`packetIndex`> deleted.",
		Content:     putResponse,
	}
	pathPacketIndex.Delete.Responses = map[string]*go2openapi.Response{
		"200": delete200,
		"401": &response401}

	pathPacketIndex.Put = pathPacketIndex.Delete
	put200 := &go2openapi.Response{
		Description: "Metadata packet replaced.",
		Content:     putResponse,
	}
	pathPacketIndex.Put.Responses["200"] = put200

	//####   /homes/{homeId}/smartModules/{smartModuleId}/metadata   ####//
	smartModuleIDMetadata := streamNameMetadata
	smartModuleIDMetadata.Parameters = pathParameters[:3]

	//####   /homes/{homeId}/metadata   ####//
	homeIDMetadata := streamNameMetadata
	homeIDMetadata.Parameters = pathParameters[0:2]

	// Paths.AddPath("/homes/{homeId}/metadata", homeIDMetadata)
	Paths.AddPath("/homes/{homeId}/smartModules/{smartModuleId}/metadata", smartModuleIDMetadata)
	Paths.AddPath("/homes/{homeId}/smartModules/{smartModuleId}/streams/{streamName}/metadata", streamNameMetadata)
	Paths.AddPath("/homes/{homeId}/smartModules/{smartModuleId}/streams/{streamName}/metadata/{packetIndex}", pathPacketIndex)
}
