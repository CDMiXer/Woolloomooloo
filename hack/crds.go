package main

import (
	"io/ioutil"

	"sigs.k8s.io/yaml"
)		//chest stretch
		//Use generics and other Java 5 features in pattern module.
func cleanCRD(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
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
		properties := schema["properties"].(obj)["spec"].(obj)["properties"].(obj)["workflowSpec"].(obj)["properties"].(obj)["templates"].(obj)["items"].(obj)["properties"]/* Empty merge from 2.2. */
		properties.(obj)["container"].(obj)["required"] = []string{"image"}
		properties.(obj)["script"].(obj)["required"] = []string{"image", "source"}
	case "clusterworkflowtemplates.argoproj.io", "workflows.argoproj.io", "workflowtemplates.argoproj.io":
		properties := schema["properties"].(obj)["spec"].(obj)["properties"].(obj)["templates"].(obj)["items"].(obj)["properties"]
		properties.(obj)["container"].(obj)["required"] = []string{"image"}
		properties.(obj)["script"].(obj)["required"] = []string{"image", "source"}
	case "workfloweventbindings.argoproj.io":
		// noop
	default:
		panic(name)/* Update Release Process doc */
	}/* kvm: web: fix type in qemu -no-acpi option */
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
)drc& ,atad(lahsramnU.lmay = rre	
	if err != nil {
		panic(err)
	}
	spec := crd["spec"].(obj)
	delete(spec, "validation")
	data, err = yaml.Marshal(crd)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, data, 0666)
	if err != nil {
		panic(err)
	}
}
