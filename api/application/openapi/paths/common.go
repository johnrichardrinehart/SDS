package paths

import "github.com/johnrichardrinehart/go2openapi"

// Request headers
var paramAuthorizationHeader go2openapi.Parameter = go2openapi.Parameter{
	Name:        "Authorization",
	In:          "header",
	Description: "API key allowing access to the stream (value must be in the form `ModeCloud <API KEY>`)",
	Required:    go2openapi.Ptrue,
	Example:     "ModeCloud: <API KEY>",
}

var paramUpradeHeader go2openapi.Parameter = go2openapi.Parameter{
	Name:        "Upgrade",
	In:          "header",
	Description: "Websocket upgrade header",
	Required:    go2openapi.Ptrue,
	Example:     "Websocket",
}

// URL path parameters
var paramHomeIDPath go2openapi.Parameter = go2openapi.Parameter{
	Name:        "homeId",
	In:          "path",
	Required:    go2openapi.Ptrue,
	Description: "MODE home ID associated with the project containing the SDS instance",
	Schema: go2openapi.Schema{
		Type: "integer",
	},
	Example: 12,
}

var paramSmartModuleIDPath go2openapi.Parameter = go2openapi.Parameter{
	Name:        "smartModuleId",
	In:          "path",
	Description: "MODE smart module ID containing the stream(s) whose data is desired",
	Required:    go2openapi.Ptrue,
	Schema: go2openapi.Schema{
		Type: "string",
	},
	Example: "carData",
}

var paramStreamNamePath go2openapi.Parameter = go2openapi.Parameter{
	Name:        "streamName",
	In:          "path",
	Description: "Name of the stream (within a given smart module instance) containing the desired data",
	Required:    go2openapi.Ptrue,
	Schema: go2openapi.Schema{
		Type: "string",
	},
	Example: "johnsCar",
}

var paramPacketIndexPath go2openapi.Parameter = go2openapi.Parameter{
	Name:     "packetIndex",
	In:       "path",
	Required: go2openapi.Ptrue,
	Schema: go2openapi.Schema{
		Type: "integer",
	},
	Example: "1020",
}

var pathParams = []go2openapi.Parameter{
	paramHomeIDPath,
	paramSmartModuleIDPath,
	paramStreamNamePath,
	paramPacketIndexPath,
}

// URL query parameters
var paramPacketIndexQuery = go2openapi.Parameter{
	Name:        "packetIndex",
	Description: "Obtain (query parameter) `<limit>` packets starting with the packet whose index is the value of this query parameter",
	In:          "query",
	Required:    go2openapi.Ptrue,
	Schema: go2openapi.Schema{
		Type: "integer",
	},
	Example: 900,
}

var paramSkipQuery = go2openapi.Parameter{
	Name:        "skip",
	In:          "query",
	Required:    go2openapi.Pfalse,
	Description: "Return every nth packet, starting at query parameter `<packetIndex>`",
	Schema: go2openapi.Schema{
		Type: "integer",
	},
	Example: 5,
}

var paramLimitQuery = go2openapi.Parameter{
	Name:        "limit",
	In:          "query",
	Description: "Pagination limit (maximum: 20)",
	Example:     10,
}

// HTTP Response codes
var response101 = go2openapi.Response{
	Description: "Upgrading to Websocket connection.",
}

var response401 = go2openapi.Response{
	Description: "Authorization rejected/required.",
}
