package paths

import (
	"github.com/johnrichardrinehart/go2openapi"
	"github.com/moderepo/main/cloud/smart_modules/SDS/api/application/openapi/paths/common"
	"github.com/moderepo/main/cloud/smart_modules/SDS/api/application/openapi/paths/schemas"
)

func init() {

	// /homes/{homeId}/smartModules/{smartModuleId}/streams/{streamName}\
	streamGet200 := common.Response200
	streamGet200.Content = schemas.StreamSchema

	streamsGet200 := common.Response200
	streamsGet200.Content = schemas.StreamsSchema

	// GET
	streamStatusGet := go2openapi.Operation{
		Responses: go2openapi.Responses{
			"200": &streamGet200,
			"401": &common.Response401,
		},
		Tags: []string{"status"},
	}

	parameters := go2openapi.Parameters{}

	parameters.AddParameters(
		common.PathParams[0:3]...,
	)

	parameters.AddParameters(
		common.ParamAuthorizationHeader,
	)

	pathItem := go2openapi.PathItem{
		Get:        &streamStatusGet,
		Parameters: parameters,
	}

	// one stream
	streamStatus := pathItem
	streamStatus.Description = "This endpoint allows application end user's to obtain current information regarding the status of a stream (which source currently owns a stream, the current data PacketIndex, etc.)."
	streamStatus.Summary = "Obtain the status of a given stream"

	// all the streams
	streamsStatus := pathItem
	streamsStatus.Parameters = go2openapi.Parameters{}
	streamsStatus.Parameters.AddParameters(
		common.PathParams[0:2]...,
	)
	streamsStatus.Parameters.AddParameters(
		common.ParamAuthorizationHeader,
	)
	streamsStatus.Description = "This endpoint allows application end user's to obtain current information regarding the all of the streams within a given smart module instance (which source currently owns a stream, the current data PacketIndex, etc.)."
	streamsStatus.Summary = "Obtain the statuses of all streams"

	streamsStatus.Get = &go2openapi.Operation{
		Responses: go2openapi.Responses{
			"200": &streamsGet200,
			"401": &common.Response401,
		},
		Tags: []string{"status"},
	}

	Paths.AddPath("/homes/{homeId}/smartModules/{smartModuleId}/streams/{streamName}", streamStatus)
	Paths.AddPath("/homes/{homeId}/smartModules/{smartModuleId}/streams", streamsStatus)
}
