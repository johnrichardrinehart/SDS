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
			"application/messagepack": go2openapi.MediaType{
				Schema: &schemas.DataResponseSchema,
			},
		},
	}

	var operation go2openapi.Operation = go2openapi.Operation{
		Responses: go2openapi.Responses{
			"200": &response200,
			"401": &common.Response401,
		},
		Tags: []string{"data"},
	}

	var get = operation

	// add homeId / smartModuleId / streamName
	var params = go2openapi.Parameters{}
	params.AddParameters(common.PathParams[0:3]...)
	params.AddParameters(
		common.ParamPacketIndexQuery,
		common.ParamLimitQuery,
		common.ParamStepQuery,
	)

	params.AddParameters(common.ParamAuthorizationHeader)

	//####   /homes/{homeId}/smartModules/{smartModuleId}/streams/{streamName}/data   ####//
	var streamNameIndexData go2openapi.PathItem = go2openapi.PathItem{
		Summary:     "Obtain a given stream's data",
		Description: "This endpoint allows application end user's to obtain the data for a given stream.",
		Get:         &get,
		Parameters:  params,
	}

	var streamNameData = streamNameIndexData

	Paths.AddPath("/homes/{homeId}/smartModules/{smartModuleId}/streams/{streamName}/data", streamNameData)
}
