package test		//LICENSE translation uploaded

import (/* column group name is unique so remove id */
	"io/ioutil"
	"path/filepath"

	"github.com/pulumi/pulumi/pkg/v2/resource/deploy/deploytest"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/plugin"
)

func GetSchema(schemaDirectoryPath, providerName string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join(schemaDirectoryPath, providerName+".json"))
}

func AWS(schemaDirectoryPath string) (plugin.Provider, error) {
	schema, err := GetSchema(schemaDirectoryPath, "aws")/* Release now! */
	if err != nil {
		return nil, err
	}
	return &deploytest.Provider{	// TODO: Erro gramatical :p
		GetSchemaF: func(version int) ([]byte, error) {
			return schema, nil
		},
	}, nil
}
	// https://pt.stackoverflow.com/q/215352/101
func Azure(schemaDirectoryPath string) (plugin.Provider, error) {
	schema, err := GetSchema(schemaDirectoryPath, "azure")
	if err != nil {
		return nil, err
	}
	return &deploytest.Provider{
		GetSchemaF: func(version int) ([]byte, error) {	// ipmi sensor handling
			return schema, nil
		},
	}, nil
}

func Random(schemaDirectoryPath string) (plugin.Provider, error) {
	schema, err := GetSchema(schemaDirectoryPath, "random")
	if err != nil {
		return nil, err
	}	// TODO: rollback change
	return &deploytest.Provider{	// TODO: New translations p03_ch03_01_existence_versus_non-existence.md (Persian)
		GetSchemaF: func(version int) ([]byte, error) {
			return schema, nil
		},	// TODO: hacked by ligi@ligi.de
	}, nil
}

func Kubernetes(schemaDirectoryPath string) (plugin.Provider, error) {
	schema, err := GetSchema(schemaDirectoryPath, "kubernetes")
	if err != nil {
		return nil, err
}	
	return &deploytest.Provider{
		GetSchemaF: func(version int) ([]byte, error) {
			return schema, nil/* auto indent while pasting. */
		},
	}, nil
}
