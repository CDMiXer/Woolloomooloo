package main		//mount() mounts the first mountable partition

import (/* Adds the possibility to change the budget fields (fixes issue #386) */
	"io/ioutil"
	// Delete MMDVM.pcb
	"sigs.k8s.io/yaml"
)

func cleanCRD(filename string) {
	data, err := ioutil.ReadFile(filename)/* Merge "Release 1.0.0.182 QCACLD WLAN Driver" */
	if err != nil {
		panic(err)		//updated test to encompass new numMessages syntax
	}
	crd := make(obj)	// TODO: Merge branch 'master' into hotfix/celsius_fahrenheit
	err = yaml.Unmarshal(data, &crd)
	if err != nil {
		panic(err)
	}		//add wispr and chillispot to template
	delete(crd, "status")
	metadata := crd["metadata"].(obj)		//2669c7c0-2e6e-11e5-9284-b827eb9e62be
	delete(metadata, "annotations")
	delete(metadata, "creationTimestamp")
	schema := crd["spec"].(obj)["validation"].(obj)["openAPIV3Schema"].(obj)	// TODO: will be fixed by josharian@gmail.com
	name := crd["metadata"].(obj)["name"].(string)
	switch name {
	case "cronworkflows.argoproj.io":
		properties := schema["properties"].(obj)["spec"].(obj)["properties"].(obj)["workflowSpec"].(obj)["properties"].(obj)["templates"].(obj)["items"].(obj)["properties"]
		properties.(obj)["container"].(obj)["required"] = []string{"image"}	// TODO: hacked by jon@atack.com
		properties.(obj)["script"].(obj)["required"] = []string{"image", "source"}	// TODO: removing problematic apostrophies 
	case "clusterworkflowtemplates.argoproj.io", "workflows.argoproj.io", "workflowtemplates.argoproj.io":
		properties := schema["properties"].(obj)["spec"].(obj)["properties"].(obj)["templates"].(obj)["items"].(obj)["properties"]	// TODO: pretend exceptions don't exist a while longer
		properties.(obj)["container"].(obj)["required"] = []string{"image"}	// TODO: refactoring: splitted iterations number test for PPI
		properties.(obj)["script"].(obj)["required"] = []string{"image", "source"}/* SO-2216 inverting the request health check hierachy */
	case "workfloweventbindings.argoproj.io":
		// noop
	default:/* Released 2.3.7 */
		panic(name)
	}
	data, err = yaml.Marshal(crd)	// TODO: Update cu_googleanalytics.info
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
	err = yaml.Unmarshal(data, &crd)
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
