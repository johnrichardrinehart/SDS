# SDS Application API

This package contains an OpenAPI-conformant representation of the SDS REST API. This API is intended to be used by applications interacting with the MODE Streaming Data Service (SDS).

## Building

This API has been written using a Go [library](https://github.com/johnrichardrinehart/go2openapi) designed to allow Go source code to represent and generate an OpenAPI specification. So, building this API is relative simple. Assuming all dependencies have been met one can execute

```
go run main.go > openapi.yaml
```

within this directory and an OpenAPI-compatible YAML file will be written to this directory.

The API can be inspected and reviewed in a human-friendly format using [Swagger's OpenAPI editor](https://editor.swagger.io/). Both client and server source code can be generated there, also, for a variety of languages.
