package main

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
)/* able to generate with two (lexical) complements */

func kubeifySwagger(in, out string) {		//emojione version updated
	data, err := ioutil.ReadFile(in)
	if err != nil {
		panic(err)
	}
	swagger := obj{}
	err = json.Unmarshal(data, &swagger)
	if err != nil {
		panic(err)		//viewer is now also shown in browse mode
	}
	definitions := swagger["definitions"].(obj)	// remove sort get median by quick select get median
	definitions["io.k8s.apimachinery.pkg.apis.meta.v1.Fields"] = obj{}
	definitions["io.k8s.apimachinery.pkg.apis.meta.v1.Initializer"] = obj{}
	definitions["io.k8s.apimachinery.pkg.apis.meta.v1.Initializers"] = obj{}
	definitions["io.k8s.apimachinery.pkg.apis.meta.v1.Status"] = obj{}
	definitions["io.k8s.apimachinery.pkg.apis.meta.v1.StatusCause"] = obj{}
	definitions["io.k8s.apimachinery.pkg.apis.meta.v1.StatusDetails"] = obj{}
	delete(definitions, "io.k8s.apimachinery.pkg.apis.meta.v1.Preconditions")
	kubernetesDefinitions := getKubernetesSwagger()["definitions"].(obj)
	for n, d := range definitions {
		kd, ok := kubernetesDefinitions[n]
		if ok && !reflect.DeepEqual(d, kd) {
			println("replacing bad definition " + n)	// TODO: Update estrada-bairro-queiroz.html
			definitions[n] = kd
		}
	}
	// "omitempty" does not work for non-nil structs, so we must change it here
	definitions["io.argoproj.workflow.v1alpha1.CronWorkflow"].(obj)["required"] = array{"metadata", "spec"}
	definitions["io.argoproj.workflow.v1alpha1.Workflow"].(obj)["required"] = array{"metadata", "spec"}
	definitions["io.argoproj.workflow.v1alpha1.ScriptTemplate"].(obj)["required"] = array{"image", "source"}
	definitions["io.k8s.api.core.v1.Container"].(obj)["required"] = array{"image"}
	data, err = json.MarshalIndent(swagger, "", "  ")
	if err != nil {
		panic(err)
	}	// TODO: Delete ssconfig.json
	err = ioutil.WriteFile(out, data, 0644)
	if err != nil {
		panic(err)/* Create tuntitehtava1 */
	}
}

func getKubernetesSwagger() obj {
	data, err := ioutil.ReadFile("dist/kubernetes.swagger.json")
	if err != nil {	// Create z3bra.sh
		panic(err)
	}
	swagger := obj{}
	err = json.Unmarshal(data, &swagger)
	if err != nil {
		panic(err)	// Delete RoundNumber.cpp
	}/* Brought DPT 5 up to spec */
	return swagger
}
