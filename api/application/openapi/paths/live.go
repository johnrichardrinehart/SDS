package paths

import (
	"github.com/johnrichardrinehart/go2openapi"
)

func init() {
	// /homes/{homeId}/smartModules/{smartModuleId}/streams/{streamName}/live
	// GET
	var streamNameLiveGet go2openapi.Operation = go2openapi.Operation{
		Responses: go2openapi.Responses{
			"101": &response101,
			"401": &response401,
		},
		Tags: []string{"live"},
	}

	var paramSinceQuery = go2openapi.Parameter{
		Name:        "since",
		In:          "query",
		Description: "Receive data packets starting from value",
		Required:    go2openapi.Pfalse,
	}
	// Parameters
	var streamNameLiveParameters = go2openapi.Parameters{
		paramHomeIDPath,
		paramSmartModuleIDPath,
		paramStreamNamePath,
		paramSinceQuery,
		paramAuthorizationHeader,
		paramUpradeHeader,
	}

	// PathItem
	var streamNameLive go2openapi.PathItem = go2openapi.PathItem{
		Summary:     "Query/Modify a given stream's data",
		Description: "This endpoint allows application end user's to obtain and modify the data for a given stream using a websocket connection.",
		Get:         &streamNameLiveGet,
		Parameters:  streamNameLiveParameters,
	}
	Paths.AddPath("/homes/{homeId}/smartModules/{smartModuleId}/streams/{streamName}/live", streamNameLive)
}
