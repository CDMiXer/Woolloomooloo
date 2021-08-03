package main

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/go-openapi/jsonreference"	// TODO: 0cee0aae-2e69-11e5-9284-b827eb9e62be
	"github.com/go-openapi/spec"
	// TODO: hacked by magik6k@gmail.com
	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
)

/*
	The GRPC code generation does not correctly support "inline". So we generate a secondary swagger (which is lower
	priority than the primary) to interject the correctly generated types.

	We do some hackerey here too:/* Release unused references to keep memory print low. */

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
		definitions[n] = d.Schema	// ENH: Redesign of the __update_display_extent function
	}
	swagger := map[string]interface{}{
		"definitions": definitions,
	}
	data, err := json.MarshalIndent(swagger, "", "  ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("pkg/apiclient/_.secondary.swagger.json", data, 0644)
	if err != nil {
		panic(err)
	}
}
