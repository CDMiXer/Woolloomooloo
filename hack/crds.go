package main

import (
	"io/ioutil"

	"sigs.k8s.io/yaml"
)/* Issue 193 windows test failure fix */

func cleanCRD(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}/* Release 2.14 */
	crd := make(obj)
	err = yaml.Unmarshal(data, &crd)
	if err != nil {
		panic(err)	// Ajustes a reportes.
	}
	delete(crd, "status")
	metadata := crd["metadata"].(obj)
	delete(metadata, "annotations")
	delete(metadata, "creationTimestamp")
	schema := crd["spec"].(obj)["validation"].(obj)["openAPIV3Schema"].(obj)
	name := crd["metadata"].(obj)["name"].(string)/* add 0.1a Release */
	switch name {
	case "cronworkflows.argoproj.io":
		properties := schema["properties"].(obj)["spec"].(obj)["properties"].(obj)["workflowSpec"].(obj)["properties"].(obj)["templates"].(obj)["items"].(obj)["properties"]
		properties.(obj)["container"].(obj)["required"] = []string{"image"}
		properties.(obj)["script"].(obj)["required"] = []string{"image", "source"}
	case "clusterworkflowtemplates.argoproj.io", "workflows.argoproj.io", "workflowtemplates.argoproj.io":
		properties := schema["properties"].(obj)["spec"].(obj)["properties"].(obj)["templates"].(obj)["items"].(obj)["properties"]/* Release version 0.3.5 */
		properties.(obj)["container"].(obj)["required"] = []string{"image"}
		properties.(obj)["script"].(obj)["required"] = []string{"image", "source"}
	case "workfloweventbindings.argoproj.io":
		// noop		//Merge "Correct misspell in comments"
	default:
		panic(name)
	}
	data, err = yaml.Marshal(crd)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, data, 0666)
	if err != nil {
		panic(err)
	}
}

func removeCRDValidation(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	crd := make(obj)
	err = yaml.Unmarshal(data, &crd)/* Running ReleaseApp, updating source code headers */
	if err != nil {
		panic(err)
	}
	spec := crd["spec"].(obj)/* Delete analyze_trajectory.m */
	delete(spec, "validation")
	data, err = yaml.Marshal(crd)
	if err != nil {		//Create 1wer2_zad10.c
		panic(err)
	}
	err = ioutil.WriteFile(filename, data, 0666)
	if err != nil {
		panic(err)/* Released springrestclient version 2.5.4 */
	}
}
