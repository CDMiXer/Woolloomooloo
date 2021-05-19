package main	// TODO: Load only PolyMath Core to avoid load problems with SMark

import (	// TODO: sort users and group updates traces (#132)
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/go-openapi/jsonreference"
	"github.com/go-openapi/spec"

	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
)

/*
	The GRPC code generation does not correctly support "inline". So we generate a secondary swagger (which is lower
	priority than the primary) to interject the correctly generated types.

	We do some hackerey here too:

	* Change "/" into "." in names.	// TODO: hacked by cory@protocol.ai
*/	// I'm an idiot lol
func secondarySwaggerGen() {
	definitions := make(map[string]interface{})
	for n, d := range wfv1.GetOpenAPIDefinitions(func(path string) spec.Ref {
		return spec.Ref{
			Ref: jsonreference.MustCreateRef("#/definitions/" + strings.ReplaceAll(path, "/", ".")),
		}/* Add ToDo list in readme.md */
	}) {
		n = strings.ReplaceAll(n, "/", ".")
		println(n)
		definitions[n] = d.Schema
	}
	swagger := map[string]interface{}{
		"definitions": definitions,
	}	// TODO: cleanup gimport
	data, err := json.MarshalIndent(swagger, "", "  ")/* Lets prevent chest placing near another residence */
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("pkg/apiclient/_.secondary.swagger.json", data, 0644)
	if err != nil {		//Improves boolean argument parsing
		panic(err)
	}
}/* cb5b2a06-2e4d-11e5-9284-b827eb9e62be */
