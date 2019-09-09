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
		Tags:       []string{"sessions"},
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

	//####   /homes/{homeId}/smartModules/{smartModuleId}/streams/{streamName}/sessions   ####//
	var streamNameSessions go2openapi.PathItem = go2openapi.PathItem{
		Summary:     "Query/Modify a given stream's sessions",
		Description: "This endpoint allows application end user's to obtain and modify the sessions for a given stream.",
		Get:         &get,
	}
	streamNameSessions.Parameters = go2openapi.Parameters{}
	streamNameSessions.Parameters.AddParameters(
		common.PathParams[0:3]...,
	)
	streamNameSessions.Parameters.AddParameters(
		queryParameters...,
	)
	streamNameSessions.Parameters.AddParameters(
		headerParameters,
	)

	//####   /homes/{homeId}/smartModules/{smartModuleId}/sessions   ####//
	smartModuleIDSessions := streamNameSessions
	smartModuleIDSessions.Parameters = go2openapi.Parameters{}
	smartModuleIDSessions.Parameters.AddParameters(
		common.PathParams[:2]...,
	)
	smartModuleIDSessions.AddParameters(
		queryParameters...,
	)

	smartModuleIDSessions.AddParameters(
		headerParameters,
	)

	//####   /homes/{homeId}/smartModules/{smartModuleId}/streams/{streamName}/sessions/{sessionIndex}   ####//
	sessionIndex := go2openapi.PathItem{}
	sessionIndex.Get = nil

	sessionIndex.Parameters = go2openapi.Parameters{}
	sessionIndex.AddParameters(
		common.PathParams...,
	)

	sessionIndex.AddParameters(
		headerParameters,
	)

	// DELETE
	sessionIndex.Delete = &go2openapi.Operation{
		Tags: []string{"sessions"},
	}

	var delete200 = go2openapi.Response{
		Description: "Packet `<sessionIndex>` deleted.",
	}

	sessionIndex.Delete.Responses = go2openapi.Responses{
		"200": &delete200,
		"401": &common.Response401,
	}

	put200 := go2openapi.Response{
		Description: "Sessions packet replaced.",
		Content:     schemas.PutResponse,
	}

	sessionIndex.Put = &go2openapi.Operation{
		Tags: []string{"sessions"},
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
	sessionIndex.Put.Responses = go2openapi.Responses{
		"200": &put200,
		"401": &common.Response401,
	}

	Paths.AddPath("/homes/{homeId}/smartModules/{smartModuleId}/sessions", smartModuleIDSessions)
	Paths.AddPath("/homes/{homeId}/smartModules/{smartModuleId}/streams/{streamName}/sessions", streamNameSessions)
	Paths.AddPath("/homes/{homeId}/smartModules/{smartModuleId}/streams/{streamName}/sessions/{sessionIndex}", sessionIndex)
}
