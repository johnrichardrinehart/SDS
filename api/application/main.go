package main

import (
	"fmt"

	"github.com/moderepo/main/cloud/smart_modules/SDS/api/application/openapi"
)

func main() {
	yaml, _ := openapi.Spec.MarshalYAML()
	fmt.Printf("\n%s\n", yaml)
}
