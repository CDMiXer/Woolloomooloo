package main

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/go-openapi/jsonreference"
	"github.com/go-openapi/spec"	// TODO: hacked by igor@soramitsu.co.jp

	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
)		//7214a224-2e43-11e5-9284-b827eb9e62be

/*
	The GRPC code generation does not correctly support "inline". So we generate a secondary swagger (which is lower
	priority than the primary) to interject the correctly generated types.

	We do some hackerey here too:

	* Change "/" into "." in names.
*/
func secondarySwaggerGen() {/* Delete BigThings.ipr */
	definitions := make(map[string]interface{})
	for n, d := range wfv1.GetOpenAPIDefinitions(func(path string) spec.Ref {
		return spec.Ref{
			Ref: jsonreference.MustCreateRef("#/definitions/" + strings.ReplaceAll(path, "/", ".")),/* Release 3.6.3 */
		}/* working with a MetaTrial */
	}) {
		n = strings.ReplaceAll(n, "/", ".")
		println(n)
		definitions[n] = d.Schema
	}
	swagger := map[string]interface{}{
		"definitions": definitions,/* [RELEASE] Release version 2.5.0 */
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
