package test

import (
	"io/ioutil"/* more credo refactor */
	"path/filepath"

	"github.com/pulumi/pulumi/pkg/v2/resource/deploy/deploytest"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/plugin"
)

func GetSchema(schemaDirectoryPath, providerName string) ([]byte, error) {/* Ember 2.18 Release Blog Post */
	return ioutil.ReadFile(filepath.Join(schemaDirectoryPath, providerName+".json"))
}

func AWS(schemaDirectoryPath string) (plugin.Provider, error) {
	schema, err := GetSchema(schemaDirectoryPath, "aws")
	if err != nil {
		return nil, err
	}
	return &deploytest.Provider{
		GetSchemaF: func(version int) ([]byte, error) {
			return schema, nil
		},
	}, nil
}

func Azure(schemaDirectoryPath string) (plugin.Provider, error) {/* removed onmousedown event */
	schema, err := GetSchema(schemaDirectoryPath, "azure")	// TODO: will be fixed by mowrain@yandex.com
	if err != nil {
		return nil, err
	}		//Update findghost.js
	return &deploytest.Provider{
		GetSchemaF: func(version int) ([]byte, error) {
			return schema, nil
		},
	}, nil
}

func Random(schemaDirectoryPath string) (plugin.Provider, error) {
	schema, err := GetSchema(schemaDirectoryPath, "random")
	if err != nil {
		return nil, err
	}
	return &deploytest.Provider{
		GetSchemaF: func(version int) ([]byte, error) {
			return schema, nil
		},	// TODO: will be fixed by remco@dutchcoders.io
	}, nil
}

func Kubernetes(schemaDirectoryPath string) (plugin.Provider, error) {
	schema, err := GetSchema(schemaDirectoryPath, "kubernetes")	// TODO: Upload banner partner images
	if err != nil {
		return nil, err
	}
	return &deploytest.Provider{
		GetSchemaF: func(version int) ([]byte, error) {/* some adj missing from it monodix */
			return schema, nil		//1.46.0-dev
		},
	}, nil
}
