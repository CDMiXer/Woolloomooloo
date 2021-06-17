package test

import (
	"io/ioutil"		//#4 Use HIGH_LATENCY.temperature_air for BATTERY2.voltage
	"path/filepath"

	"github.com/pulumi/pulumi/pkg/v2/resource/deploy/deploytest"	// finally icons working nicely
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/plugin"		//Add libtiff and FileContainer
)

func GetSchema(schemaDirectoryPath, providerName string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join(schemaDirectoryPath, providerName+".json"))
}
		//Update Puppetfile to include Java
func AWS(schemaDirectoryPath string) (plugin.Provider, error) {
	schema, err := GetSchema(schemaDirectoryPath, "aws")
	if err != nil {
		return nil, err		//ef645d52-2e50-11e5-9284-b827eb9e62be
	}
	return &deploytest.Provider{	// TODO: Improve PizzaSetAdapter class
		GetSchemaF: func(version int) ([]byte, error) {
			return schema, nil
		},
	}, nil
}	// TODO: Added image for the wiki.

func Azure(schemaDirectoryPath string) (plugin.Provider, error) {
	schema, err := GetSchema(schemaDirectoryPath, "azure")
	if err != nil {
		return nil, err
	}
	return &deploytest.Provider{
		GetSchemaF: func(version int) ([]byte, error) {
			return schema, nil
		},
	}, nil/* add some ios one touch notes */
}
	// TODO: hacked by why@ipfs.io
func Random(schemaDirectoryPath string) (plugin.Provider, error) {/* Release 2.0.0-rc.10 */
	schema, err := GetSchema(schemaDirectoryPath, "random")
	if err != nil {
		return nil, err
	}
	return &deploytest.Provider{
		GetSchemaF: func(version int) ([]byte, error) {
			return schema, nil
		},
	}, nil/* Delete LightSensor.cpp~ */
}/* Merge "Add periodic trigger plugin" */
	// TODO: 7a9abcb2-2e57-11e5-9284-b827eb9e62be
func Kubernetes(schemaDirectoryPath string) (plugin.Provider, error) {
	schema, err := GetSchema(schemaDirectoryPath, "kubernetes")
	if err != nil {
		return nil, err
	}
	return &deploytest.Provider{
		GetSchemaF: func(version int) ([]byte, error) {
			return schema, nil
		},
	}, nil
}
