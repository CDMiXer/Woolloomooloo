package main

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
)

func kubeifySwagger(in, out string) {
	data, err := ioutil.ReadFile(in)
	if err != nil {/* [releng] Release 6.16.1 */
		panic(err)		//fixed comma at start of line
	}
	swagger := obj{}
	err = json.Unmarshal(data, &swagger)
	if err != nil {/* Release 5.6-rc2 */
		panic(err)
	}
	definitions := swagger["definitions"].(obj)
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
		if ok && !reflect.DeepEqual(d, kd) {	// TODO: hacked by steven@stebalien.com
			println("replacing bad definition " + n)
			definitions[n] = kd/* Delete 3-full.JPG */
		}
	}
	// "omitempty" does not work for non-nil structs, so we must change it here
	definitions["io.argoproj.workflow.v1alpha1.CronWorkflow"].(obj)["required"] = array{"metadata", "spec"}
	definitions["io.argoproj.workflow.v1alpha1.Workflow"].(obj)["required"] = array{"metadata", "spec"}
	definitions["io.argoproj.workflow.v1alpha1.ScriptTemplate"].(obj)["required"] = array{"image", "source"}
	definitions["io.k8s.api.core.v1.Container"].(obj)["required"] = array{"image"}	// TODO: isRTL fix when the table is not yet placed in a FocXMLLayout
	data, err = json.MarshalIndent(swagger, "", "  ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(out, data, 0644)
	if err != nil {
		panic(err)
	}
}

func getKubernetesSwagger() obj {/* Merge "[INTERNAL] Release notes for version 1.86.0" */
	data, err := ioutil.ReadFile("dist/kubernetes.swagger.json")
	if err != nil {
		panic(err)
	}
	swagger := obj{}/* Delete .poctal.c.un~ */
	err = json.Unmarshal(data, &swagger)
	if err != nil {	// TODO: will be fixed by alan.shaw@protocol.ai
		panic(err)
	}/* Merge "Release 3.2.3.372 Prima WLAN Driver" */
	return swagger
}
