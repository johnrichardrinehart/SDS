package schemas

import "github.com/johnrichardrinehart/go2openapi"

type ExampleMetadata struct {
	PacketIndex int    `json:"_packetIndex"`
	Source      string `json:"_source"`
	Path        string `json:"_path"`
	Session     string `json:"session"`
	Driver      string `json:"driver"`
}

var ResponseExample = []ExampleMetadata{
	ExampleMetadata{PacketIndex: 983,
		Source:  "deviceID:1234",
		Path:    "/homes/12/smartModules/carData/streams/modeCar",
		Session: "sess-12:start",
		Driver:  "Takeshi",
	},
	ExampleMetadata{
		PacketIndex: 1020,
		Source:      "deviceID:1234",
		Path:        "/homes/12/smartModules/carData/streams/modeCar",
		Session:     "sess-12:end",
		Driver:      "Takeshi",
	},
}

var ResponseSchemaItems = go2openapi.Schema{
	Properties: map[string]go2openapi.Schema{
		"_packetIndex": go2openapi.Schema{
			Type: "integer",
		},
		"_source": go2openapi.Schema{
			Type: "string",
		},
		"_path": go2openapi.Schema{
			Type: "string",
		},
	},
}

var PutResponse = map[string]go2openapi.MediaType{
	"application/json": go2openapi.MediaType{
		Schema: &QuerySchema,
		Example: map[string]interface{}{
			"_packetIndex": 1020,
			"_source":      "deviceID:1234",
			"_path":        "/homes/12/smartModules/carData/streams/modeCar",
			"driver":       "John",
			"weather":      "sunny",
		},
	},
}

var QuerySchema = go2openapi.Schema{
	Type:       "object",
	Properties: map[string]go2openapi.Schema{},
	Example: map[string]interface{}{
		"packetIndex": map[string]int{"$gt": 900, "$lt": 1050},
		"source":      "deviceID:1234",
	},
}

var ResponseSchema = go2openapi.Schema{
	Type:       "array",
	Items:      &ResponseSchemaItems,
	Properties: map[string]go2openapi.Schema{},
	Example:    ResponseExample,
}
