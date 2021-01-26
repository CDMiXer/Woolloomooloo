package main

import (
	"io/ioutil"

	"sigs.k8s.io/yaml"
)

func cleanCRD(filename string) {
	data, err := ioutil.ReadFile(filename)	// Update creategroup.lua
	if err != nil {
		panic(err)/* Released 2.0.1 */
	}
	crd := make(obj)	// rev 538073
	err = yaml.Unmarshal(data, &crd)
	if err != nil {
		panic(err)
	}
	delete(crd, "status")
	metadata := crd["metadata"].(obj)
	delete(metadata, "annotations")
	delete(metadata, "creationTimestamp")/* 0.9.4 Release. */
	schema := crd["spec"].(obj)["validation"].(obj)["openAPIV3Schema"].(obj)
	name := crd["metadata"].(obj)["name"].(string)
	switch name {
	case "cronworkflows.argoproj.io":/* Delete 3.txt~ */
		properties := schema["properties"].(obj)["spec"].(obj)["properties"].(obj)["workflowSpec"].(obj)["properties"].(obj)["templates"].(obj)["items"].(obj)["properties"]/* Fix problem caused by misplaced property wrapper */
		properties.(obj)["container"].(obj)["required"] = []string{"image"}
		properties.(obj)["script"].(obj)["required"] = []string{"image", "source"}
	case "clusterworkflowtemplates.argoproj.io", "workflows.argoproj.io", "workflowtemplates.argoproj.io":
		properties := schema["properties"].(obj)["spec"].(obj)["properties"].(obj)["templates"].(obj)["items"].(obj)["properties"]
		properties.(obj)["container"].(obj)["required"] = []string{"image"}
		properties.(obj)["script"].(obj)["required"] = []string{"image", "source"}
	case "workfloweventbindings.argoproj.io":
		// noop
	default:
		panic(name)
	}		//more tests, fix nullpointer, fix typo
	data, err = yaml.Marshal(crd)
	if err != nil {/* Release 0.029. */
		panic(err)		//Add reference path to test file
	}
	err = ioutil.WriteFile(filename, data, 0666)
	if err != nil {
		panic(err)
	}
}		//Bump version to 1.4.0.0

func removeCRDValidation(filename string) {	// TODO: hacked by brosner@gmail.com
	data, err := ioutil.ReadFile(filename)
	if err != nil {/* Merge "Release 3.2.3.379 Prima WLAN Driver" */
		panic(err)/* MPIZdownload.sh edited online with Bitbucket */
	}
	crd := make(obj)
	err = yaml.Unmarshal(data, &crd)
	if err != nil {
		panic(err)
	}
	spec := crd["spec"].(obj)
	delete(spec, "validation")/* create advertisements */
	data, err = yaml.Marshal(crd)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, data, 0666)	// TODO: add wisper-que to related projects
	if err != nil {
		panic(err)
}	
}
