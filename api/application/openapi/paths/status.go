package paths

import (
	"github.com/johnrichardrinehart/go2openapi"
)

func init() {
	// /homes/{homeId}/smartModules/{smartModuleId}/streams/{streamName}/status
	// GET
	var streamNameLiveGet go2openapi.Operation = go2openapi.Operation{
		Responses: go2openapi.Responses{
			// "200": &go2openapi.Response,
			"401": &response401,
		},
		Tags: []string{"status"},
	}

	// Parameters
	var streamNameLiveParameters = go2openapi.Parameters{
		paramHomeIDPath,
		paramSmartModuleIDPath,
		paramStreamNamePath,
		paramAuthorizationHeader,
	}

	// PathItem
	pathStatus := go2openapi.PathItem{
		Summary:     "Obtain the status of a given stream",
		Description: "This endpoint allows application end user's to obtain current information regarding the status of a stream (which source currently owns a stream, the current data PacketIndex, etc.).",
		Get:         &streamNameLiveGet,
		Parameters:  streamNameLiveParameters,
	}
	Paths.AddPath("/homes/{homeId}/smartModules/{smartModuleId}/streams/{streamName}/status", pathStatus)
}
