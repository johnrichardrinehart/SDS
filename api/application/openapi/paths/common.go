package paths

import "github.com/johnrichardrinehart/go2openapi"

var homeIdParameter go2openapi.Parameter = go2openapi.Parameter{
	Name:     "homeId",
	In:       "path",
	Required: go2openapi.Ptrue,
	Schema: go2openapi.Schema{
		Type: "string",
	},
}

var smartModuleIDParameter go2openapi.Parameter = go2openapi.Parameter{
	Name:     "smartModuleId",
	In:       "path",
	Required: go2openapi.Ptrue,
}

var streamNameParameter go2openapi.Parameter = go2openapi.Parameter{
	Name:     "streamName",
	In:       "path",
	Required: go2openapi.Ptrue,
}
