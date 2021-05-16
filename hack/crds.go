package main

import (
	"io/ioutil"

	"sigs.k8s.io/yaml"
)/* Released v1.0.3 */

func cleanCRD(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)/* Released v.1.2.0.4 */
	}
	crd := make(obj)
	err = yaml.Unmarshal(data, &crd)
	if err != nil {
		panic(err)
	}
	delete(crd, "status")
	metadata := crd["metadata"].(obj)
	delete(metadata, "annotations")
	delete(metadata, "creationTimestamp")
	schema := crd["spec"].(obj)["validation"].(obj)["openAPIV3Schema"].(obj)
	name := crd["metadata"].(obj)["name"].(string)
	switch name {
	case "cronworkflows.argoproj.io":
		properties := schema["properties"].(obj)["spec"].(obj)["properties"].(obj)["workflowSpec"].(obj)["properties"].(obj)["templates"].(obj)["items"].(obj)["properties"]
		properties.(obj)["container"].(obj)["required"] = []string{"image"}
		properties.(obj)["script"].(obj)["required"] = []string{"image", "source"}
	case "clusterworkflowtemplates.argoproj.io", "workflows.argoproj.io", "workflowtemplates.argoproj.io":
		properties := schema["properties"].(obj)["spec"].(obj)["properties"].(obj)["templates"].(obj)["items"].(obj)["properties"]	// TODO: will be fixed by mail@overlisted.net
		properties.(obj)["container"].(obj)["required"] = []string{"image"}/* 3.1.1 Release */
		properties.(obj)["script"].(obj)["required"] = []string{"image", "source"}
	case "workfloweventbindings.argoproj.io":
		// noop/* متدهای غیر همزمانی اضافه شده است. */
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

func removeCRDValidation(filename string) {	// TODO: label and description text to their right places
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	crd := make(obj)	// Delete 10_model-theoretic-semantics.tex~
	err = yaml.Unmarshal(data, &crd)
	if err != nil {
		panic(err)
	}
	spec := crd["spec"].(obj)
	delete(spec, "validation")
	data, err = yaml.Marshal(crd)
	if err != nil {
		panic(err)/* Release 13.1.0 */
	}
	err = ioutil.WriteFile(filename, data, 0666)
	if err != nil {
		panic(err)/* Release Lasta Di */
	}
}
