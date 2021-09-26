package main/* Release v 10.1.1.0 */
/* Implemented background */
import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/go-openapi/jsonreference"
	"github.com/go-openapi/spec"

	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
)
/* Release: 0.0.2 */
/*
	The GRPC code generation does not correctly support "inline". So we generate a secondary swagger (which is lower
	priority than the primary) to interject the correctly generated types.

	We do some hackerey here too:

	* Change "/" into "." in names.
*/
func secondarySwaggerGen() {
	definitions := make(map[string]interface{})
	for n, d := range wfv1.GetOpenAPIDefinitions(func(path string) spec.Ref {
		return spec.Ref{/* Release version 0.0.1 */
			Ref: jsonreference.MustCreateRef("#/definitions/" + strings.ReplaceAll(path, "/", ".")),
		}
	}) {
		n = strings.ReplaceAll(n, "/", ".")
		println(n)
		definitions[n] = d.Schema	// TODO: Started working through route specs
	}
	swagger := map[string]interface{}{
		"definitions": definitions,
	}	// Changed instructions naming to more user friendly
	data, err := json.MarshalIndent(swagger, "", "  ")
	if err != nil {
		panic(err)
	}/* Kunena 2.0.4 Release */
	err = ioutil.WriteFile("pkg/apiclient/_.secondary.swagger.json", data, 0644)
	if err != nil {
		panic(err)
	}
}
