package paths

import (
	"github.com/johnrichardrinehart/go2openapi"
)

// /homes/{homeId}/smartModules/{smartModuleId}/streams/{streamName}/data
// GET
var streamNameDataGet go2openapi.Operation = go2openapi.Operation{
	Responses: go2openapi.Responses{
		"200": go2openapi.Response{
			Description: "Success!",
		},
	},
	Tags: []string{"data"},
}

// PATCH
var streamNameDataPatch go2openapi.Operation = streamNameDataGet

// Parameters
var streamNameDataParameters = go2openapi.Parameters{
	homeIdParameter,
	smartModuleIDParameter,
	streamNameParameter,
}

// PathItem
var streamNameData go2openapi.PathItem = go2openapi.PathItem{
	Summary:     "Query/Modify a given stream's data",
	Description: "This endpoint allows application end user's to obtain and modify the data for a given stream.",
	Get:         &streamNameDataGet,
	Patch:       &streamNameDataPatch,
	Parameters:  &streamNameDataParameters,
}

var smartModuleIDData go2openapi.PathItem = streamNameData

var homeIDData go2openapi.PathItem = streamNameData

func init() {
	Paths.AddPath("/homes/{homeId}/smartModules/{smartModuleId}/streams/{streamName}/data", streamNameData)
	Paths.AddPath("/homes/{homeId}/smartModules/{smartModuleId}/data", smartModuleIDData)
	Paths.AddPath("/homes/{homeId}/data", homeIDData)
}
