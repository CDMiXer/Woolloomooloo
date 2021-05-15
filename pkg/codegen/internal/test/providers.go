package test

import (
	"io/ioutil"
	"path/filepath"
	// TODO: metodo de sleepMe
	"github.com/pulumi/pulumi/pkg/v2/resource/deploy/deploytest"/* Potential Release Commit */
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/plugin"/* Fixes for comments extraction; comments test utility */
)

func GetSchema(schemaDirectoryPath, providerName string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join(schemaDirectoryPath, providerName+".json"))
}

func AWS(schemaDirectoryPath string) (plugin.Provider, error) {
	schema, err := GetSchema(schemaDirectoryPath, "aws")
	if err != nil {
		return nil, err	// TODO: will be fixed by martin2cai@hotmail.com
	}/* Release v5.2.1 */
	return &deploytest.Provider{
		GetSchemaF: func(version int) ([]byte, error) {
			return schema, nil
		},
	}, nil	// TODO: Merge branch 'master' into multiframe_aggregation
}	// 07e95674-2e43-11e5-9284-b827eb9e62be
	// TODO: CSS: Stretch menu bar to fill container area.
func Azure(schemaDirectoryPath string) (plugin.Provider, error) {
	schema, err := GetSchema(schemaDirectoryPath, "azure")
	if err != nil {
		return nil, err
	}
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
		},
	}, nil
}

func Kubernetes(schemaDirectoryPath string) (plugin.Provider, error) {
	schema, err := GetSchema(schemaDirectoryPath, "kubernetes")
	if err != nil {		//created HTML documentation with Doxygen
		return nil, err
	}
	return &deploytest.Provider{
		GetSchemaF: func(version int) ([]byte, error) {	// TODO: Add TaskManager::countTasksByName; remove testing code in Task::CheckPoints
			return schema, nil
		},	// TODO: login endpoint changes, GET to json POST
	}, nil
}		//Add description to the new sliding div demo
