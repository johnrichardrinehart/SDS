package common

import "github.com/johnrichardrinehart/go2openapi"

// Request headers
var ParamAuthorizationHeader go2openapi.Parameter = go2openapi.Parameter{
	Name:        "Authorization",
	In:          "header",
	Description: "API key allowing access to the stream (value must be in the form `ModeCloud <API KEY>`)",
	Required:    go2openapi.Ptrue,
	Example:     "ModeCloud: <API KEY>",
}

var ParamUpradeHeader go2openapi.Parameter = go2openapi.Parameter{
	Name:        "Upgrade",
	In:          "header",
	Description: "Websocket upgrade header",
	Required:    go2openapi.Ptrue,
	Example:     "Websocket",
}

// URL path parameters
var ParamHomeIDPath go2openapi.Parameter = go2openapi.Parameter{
	Name:        "homeId",
	In:          "path",
	Required:    go2openapi.Ptrue,
	Description: "MODE home ID associated with the project containing the SDS instance",
	Schema: go2openapi.Schema{
		Type: "integer",
	},
	Example: 12,
}

var ParamSmartModuleIDPath go2openapi.Parameter = go2openapi.Parameter{
	Name:        "smartModuleId",
	In:          "path",
	Description: "MODE smart module ID containing the stream(s) whose data is desired",
	Required:    go2openapi.Ptrue,
	Schema: go2openapi.Schema{
		Type: "string",
	},
	Example: "carData",
}

var ParamStreamNamePath go2openapi.Parameter = go2openapi.Parameter{
	Name:        "streamName",
	In:          "path",
	Description: "Name of the stream (within a given smart module instance) containing the desired data",
	Required:    go2openapi.Ptrue,
	Schema: go2openapi.Schema{
		Type: "string",
	},
	Example: "johnsCar",
}

var ParamPacketIndexPath go2openapi.Parameter = go2openapi.Parameter{
	Name:        "packetIndex",
	In:          "path",
	Required:    go2openapi.Ptrue,
	Description: "SDS index of the data packet that is being requested",
	Schema: go2openapi.Schema{
		Type: "integer",
	},
	Example: "1020",
}

var PathParams = []go2openapi.Parameter{
	ParamHomeIDPath,
	ParamSmartModuleIDPath,
	ParamStreamNamePath,
	ParamPacketIndexPath,
}

// URL query parameters
var ParamPacketIndexQuery = go2openapi.Parameter{
	Name:        "packetIndex",
	Description: "Obtain (query parameter) `<limit>` packets starting with the packet whose index is the value of this query parameter",
	In:          "query",
	Required:    go2openapi.Ptrue,
	Schema: go2openapi.Schema{
		Type: "integer",
	},
	Example: 900,
}

var ParamStepQuery = go2openapi.Parameter{
	Name:        "step",
	In:          "query",
	Required:    go2openapi.Pfalse,
	Description: "Return every nth packet, starting at query parameter `<packetIndex>`",
	Schema: go2openapi.Schema{
		Type: "integer",
	},
	Example: 5,
}

var ParamLimitQuery = go2openapi.Parameter{
	Name:        "limit",
	In:          "query",
	Description: "Pagination limit (maximum: 20)",
	Example:     10,
}

var ParamSkipQuery = go2openapi.Parameter{
	Name:        "skip",
	In:          "query",
	Description: "Skip the first `n` results",
	Example:     10,
}

// HTTP Response codes
var Response101 = go2openapi.Response{
	Description: "Upgrading to Websocket connection.",
}

var Response401 = go2openapi.Response{
	Description: "Authorization rejected/required.",
}

var Response200 = go2openapi.Response{
	Description: "Successful query",
}
