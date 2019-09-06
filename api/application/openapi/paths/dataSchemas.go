package paths

import "github.com/johnrichardrinehart/go2openapi"

var dataResponseSchemaItems = go2openapi.Schema{
	Type: "array",
	Items: &go2openapi.Schema{
		Type:   "string",
		Format: "byte",
	},
}

var dataQuerySchema = go2openapi.Schema{
	Type:       "object",
	Properties: map[string]go2openapi.Schema{},
	Example: map[string]interface{}{
		"packetIndex": map[string]int{"$gt": 900, "$lt": 1050},
		"source":      "deviceID:1234",
	},
}

var dataResponseSchema = go2openapi.Schema{
	Type:  "array",
	Items: &dataResponseSchemaItems,
	Example: `[]MessagePackItems{
		MessagePack{ bin(data) },
		MessagePack{ bin(data) },
	}`,
}
