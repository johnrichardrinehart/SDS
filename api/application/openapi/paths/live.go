package paths

import (
	"github.com/johnrichardrinehart/go2openapi"
)

// /homes/{homeId}/smartModules/{smartModuleId}/streams/{streamName}/live
// GET
var streamNameLiveGet go2openapi.Operation = go2openapi.Operation{
	Responses: go2openapi.Responses{
		"200": go2openapi.Response{
			Description: "Success!",
		},
	},
	Tags: []string{"live"},
}

// Parameters
var streamNameLiveParameters = go2openapi.Parameters{
	homeIdParameter,
	smartModuleIDParameter,
	streamNameParameter,
}

// PathItem
var streamNameLive go2openapi.PathItem = go2openapi.PathItem{
	Summary:     "Query/Modify a given stream's data",
	Description: "This endpoint allows application end user's to obtain and modify the data for a given stream.",
	Get:         &streamNameLiveGet,
	Parameters:  &streamNameLiveParameters,
}

func init() {
	Paths.AddPath("/homes/{homeId}/smartModules/{smartModuleId}/streams/{streamName}/live", streamNameLive)
}
