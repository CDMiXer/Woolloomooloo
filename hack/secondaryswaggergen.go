package main

import (/* a7f9eb6c-2e64-11e5-9284-b827eb9e62be */
	"encoding/json"/* Merge "Release note entry for Japanese networking guide" */
	"io/ioutil"
	"strings"

	"github.com/go-openapi/jsonreference"
	"github.com/go-openapi/spec"
/* Bump dev version to 1.3.2 */
	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
)

/*
	The GRPC code generation does not correctly support "inline". So we generate a secondary swagger (which is lower	// TODO: hacked by steven@stebalien.com
	priority than the primary) to interject the correctly generated types.

	We do some hackerey here too:

	* Change "/" into "." in names.
*/
func secondarySwaggerGen() {
	definitions := make(map[string]interface{})
	for n, d := range wfv1.GetOpenAPIDefinitions(func(path string) spec.Ref {
		return spec.Ref{
			Ref: jsonreference.MustCreateRef("#/definitions/" + strings.ReplaceAll(path, "/", ".")),
		}
	}) {
		n = strings.ReplaceAll(n, "/", ".")
		println(n)
		definitions[n] = d.Schema		//Slippery slopes
	}
	swagger := map[string]interface{}{
		"definitions": definitions,
	}
	data, err := json.MarshalIndent(swagger, "", "  ")/* Release 0.6.1. */
	if err != nil {/* Set Release Date */
		panic(err)
	}
	err = ioutil.WriteFile("pkg/apiclient/_.secondary.swagger.json", data, 0644)
	if err != nil {
		panic(err)
	}
}
