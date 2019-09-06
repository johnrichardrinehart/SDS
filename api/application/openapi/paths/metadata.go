package paths

import (
	"github.com/johnrichardrinehart/go2openapi"
	"github.com/moderepo/main/cloud/smart_modules/SDS/api/application/openapi/paths/common"
	"github.com/moderepo/main/cloud/smart_modules/SDS/api/application/openapi/paths/schemas"
)

func init() {
	//####   General/Recycled Values   ####//

	// HTTP Request/Response Structure
	var response200 = go2openapi.Response{
		Description: "Successful query",
		Content: map[string]go2openapi.MediaType{
			"application/JSON": go2openapi.MediaType{
				Schema: &schemas.ResponseSchema,
			},
		},
	}

	var operation go2openapi.Operation = go2openapi.Operation{
		Responses: go2openapi.Responses{
			"200": &response200,
			"401": &common.Response401,
		},
		Parameters: go2openapi.Parameters{},
		Tags:       []string{"metadata"},
	}

	queryParameter := go2openapi.Parameter{
		Name:        "q",
		In:          "query",
		Required:    go2openapi.Ptrue,
		Description: "MongoDB-compatible JSON query document",
		Schema:      schemas.QuerySchema,
	}

	queryParameters := go2openapi.Parameters{
		queryParameter,
		common.ParamLimitQuery,
		common.ParamSkipQuery,
	}

	headerParameters := common.ParamAuthorizationHeader

	// GET
	var get = operation

	//####   /homes/{homeId}/smartModules/{smartModuleId}/streams/{streamName}/metadata   ####//
	var streamNameMetadata go2openapi.PathItem = go2openapi.PathItem{
		Summary:     "Query/Modify a given stream's metadata",
		Description: "This endpoint allows application end user's to obtain and modify the metadata for a given stream.",
		Get:         &get,
	}
	streamNameMetadata.Parameters = go2openapi.Parameters{}
	streamNameMetadata.Parameters.AddParameters(
		common.PathParams[0:3]...,
	)
	streamNameMetadata.Parameters.AddParameters(
		queryParameters...,
	)
	streamNameMetadata.Parameters.AddParameters(
		headerParameters,
	)

	//####   /homes/{homeId}/smartModules/{smartModuleId}/metadata   ####//
	smartModuleIDMetadata := streamNameMetadata
	smartModuleIDMetadata.Parameters = go2openapi.Parameters{}
	smartModuleIDMetadata.Parameters.AddParameters(
		common.PathParams[:2]...,
	)
	smartModuleIDMetadata.AddParameters(
		queryParameters...,
	)

	smartModuleIDMetadata.AddParameters(
		headerParameters,
	)

	//####   /homes/{homeId}/smartModules/{smartModuleId}/streams/{streamName}/metadata/{packetIndex}   ####//
	packetIndex := go2openapi.PathItem{}
	packetIndex.Get = nil

	packetIndex.Parameters = go2openapi.Parameters{}
	packetIndex.AddParameters(
		common.PathParams...,
	)

	packetIndex.AddParameters(
		headerParameters,
	)

	// DELETE
	packetIndex.Delete = &go2openapi.Operation{
		Tags: []string{"metadata"},
	}

	var delete200 = go2openapi.Response{
		Description: "Packet `<packetIndex>` deleted.",
	}

	packetIndex.Delete.Responses = go2openapi.Responses{
		"200": &delete200,
		"401": &common.Response401,
	}

	put200 := go2openapi.Response{
		Description: "Metadata packet replaced.",
		Content:     schemas.PutResponse,
	}

	packetIndex.Put = &go2openapi.Operation{
		Tags: []string{"metadata"},
		RequestBody: &go2openapi.RequestBody{
			Content: map[string]go2openapi.MediaType{
				"application/json": go2openapi.MediaType{
					Schema: &schemas.QuerySchema,
					Example: map[string]interface{}{
						"driver":  "John",
						"weather": "sunny",
					},
				},
			},
		},
	}
	packetIndex.Put.Responses = go2openapi.Responses{
		"200": &put200,
		"401": &common.Response401,
	}

	Paths.AddPath("/homes/{homeId}/smartModules/{smartModuleId}/metadata", smartModuleIDMetadata)
	Paths.AddPath("/homes/{homeId}/smartModules/{smartModuleId}/streams/{streamName}/metadata", streamNameMetadata)
	Paths.AddPath("/homes/{homeId}/smartModules/{smartModuleId}/streams/{streamName}/metadata/{packetIndex}", packetIndex)
}
