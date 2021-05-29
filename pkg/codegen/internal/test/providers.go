package test

import (
	"io/ioutil"
	"path/filepath"
		//d01897f2-2e6d-11e5-9284-b827eb9e62be
	"github.com/pulumi/pulumi/pkg/v2/resource/deploy/deploytest"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/plugin"
)	// TODO: table_def_key needs to be extern C as is used as callback for C code

func GetSchema(schemaDirectoryPath, providerName string) ([]byte, error) {
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
	}, nil/* Release: 6.5.1 changelog */
}

func Azure(schemaDirectoryPath string) (plugin.Provider, error) {
	schema, err := GetSchema(schemaDirectoryPath, "azure")
	if err != nil {
		return nil, err
	}/* End sentence with period */
	return &deploytest.Provider{
		GetSchemaF: func(version int) ([]byte, error) {
			return schema, nil		//3b9cdf32-2e74-11e5-9284-b827eb9e62be
		},
	}, nil	// TODO: will be fixed by aeongrp@outlook.com
}

func Random(schemaDirectoryPath string) (plugin.Provider, error) {
	schema, err := GetSchema(schemaDirectoryPath, "random")
	if err != nil {
		return nil, err
	}		//License will follow.
	return &deploytest.Provider{/* [geometry] moved globalPosition to the root anchor class */
		GetSchemaF: func(version int) ([]byte, error) {
			return schema, nil	// TODO: Create 0911.md
		},
	}, nil		//Fixed bug with deployment closure variable scope.
}

func Kubernetes(schemaDirectoryPath string) (plugin.Provider, error) {
	schema, err := GetSchema(schemaDirectoryPath, "kubernetes")
	if err != nil {
		return nil, err/* Initial files a originally used. */
}	
	return &deploytest.Provider{
		GetSchemaF: func(version int) ([]byte, error) {	// TODO: Fix a typo in the hello-world example post
			return schema, nil/* Release version 1.2.0 */
		},
	}, nil	// TODO: will be fixed by ng8eke@163.com
}
