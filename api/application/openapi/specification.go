package openapi

import (
	"github.com/johnrichardrinehart/go2openapi"
	"github.com/moderepo/main/cloud/smart_modules/SDS/api/application/openapi/paths"
)

var SDSAPIVERSION = "0.0.1"

var info go2openapi.Info = go2openapi.Info{
	Title:       "Streaming Data Service REST API",
	Description: "This API will be used by applications wishing to obtain data or metadata from the MODE SDS.",
	Version:     SDSAPIVERSION,
}

var Spec go2openapi.Specification = go2openapi.Specification{
	OpenAPIVersion: go2openapi.OPENAPIVERSION,
	Info:           info,
	Paths:          paths.Paths,
}
