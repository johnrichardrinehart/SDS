package components

import "github.com/johnrichardrinehart/go2openapi"

var updateSchema = go2openapi.Schema{
	Type: "object",
	Properties: map[string]go2openapi.Schema{
		"someStringProp": go2openapi.Schema{
			Type: "string",
		},
		"someIntProp": go2openapi.Schema{
			Type: "integer",
		},
	},
}

var schemas = make(go2openapi.Schemas)

func init() {
	schemas.AddSchema("updateObject", updateSchema)
	Components.Schemas = &schemas
}
