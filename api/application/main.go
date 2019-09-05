package main

import (
	"log"
	"os"

	"github.com/moderepo/main/cloud/smart_modules/SDS/api/application/openapi"
)

func main() {
	yaml, _ := openapi.Spec.MarshalYAML()
	if err := os.Remove("openapi.yaml"); err != nil {
		log.Fatal(err)
	}
	f, err := os.OpenFile("openapi.yaml", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if _, err := f.Write(yaml); err != nil {
		log.Fatal(err)
	}
}
