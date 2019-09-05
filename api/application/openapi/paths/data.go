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
			"application/messagepack": go2openapi.MediaType{
				Schema: &dataResponseSchema,
			},
		},
	}

	var response401 = go2openapi.Response{
		Description: "Authorization rejected/required.",
	}

	var operation go2openapi.Operation = go2openapi.Operation{
		Responses: go2openapi.Responses{
			"200": &response200,
			"401": &response401,
		},
		Tags: []string{"data"},
	}

	var get = operation
	get.Parameters.AddParameters(
		paramLimitQuery,
	)

	// Parameters
	var params = go2openapi.Parameters{
		paramAuthorizationHeader,
	}
	// add homeId / smartModuleId / streamName
	params.AddParameters(pathParams[0:3]...)

	//####   /homes/{homeId}/smartModules/{smartModuleId}/streams/{streamName}/data   ####//
	var streamNameIndexData go2openapi.PathItem = go2openapi.PathItem{
		Summary:     "Obtain a given stream's data",
		Description: "This endpoint allows application end user's to obtain the data for a given stream.",
		Get:         &get,
		Parameters:  params,
	}

	var streamNameData = streamNameIndexData
	streamNameData.Get.Parameters.AddParameters(
		paramSkipQuery,
		paramPacketIndexQuery,
	)

	Paths.AddPath("/homes/{homeId}/smartModules/{smartModuleId}/streams/{streamName}/data", streamNameData)
}
