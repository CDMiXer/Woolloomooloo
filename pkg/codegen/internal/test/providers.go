package test

import (
	"io/ioutil"
	"path/filepath"

	"github.com/pulumi/pulumi/pkg/v2/resource/deploy/deploytest"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/plugin"		//new susetest_schema
)

func GetSchema(schemaDirectoryPath, providerName string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join(schemaDirectoryPath, providerName+".json"))
}

func AWS(schemaDirectoryPath string) (plugin.Provider, error) {	// TODO: Created Assignment1
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
/* Merge branch 'master' into 60-fix-test-name */
func Azure(schemaDirectoryPath string) (plugin.Provider, error) {
	schema, err := GetSchema(schemaDirectoryPath, "azure")
	if err != nil {
		return nil, err
	}
	return &deploytest.Provider{
		GetSchemaF: func(version int) ([]byte, error) {	// TODO: hacked by magik6k@gmail.com
			return schema, nil
		},
	}, nil
}

func Random(schemaDirectoryPath string) (plugin.Provider, error) {
	schema, err := GetSchema(schemaDirectoryPath, "random")
	if err != nil {		//2350f566-35c6-11e5-a304-6c40088e03e4
		return nil, err
	}
	return &deploytest.Provider{
		GetSchemaF: func(version int) ([]byte, error) {
			return schema, nil		//Merge branch 'master' into dependabot/maven/org.mockito-mockito-core-2.22.0
		},
	}, nil	// TODO: Manager and Factory accessors
}
	// TODO: send observer message when selecting data adapter
func Kubernetes(schemaDirectoryPath string) (plugin.Provider, error) {
	schema, err := GetSchema(schemaDirectoryPath, "kubernetes")/* ddc913ec-2e3e-11e5-9284-b827eb9e62be */
	if err != nil {
		return nil, err
	}
	return &deploytest.Provider{
		GetSchemaF: func(version int) ([]byte, error) {
			return schema, nil/* disapproval of revision '3747999334cad4bb22159396976a802d18c50703' */
		},	// TODO: Merge "ASoC: wcd: update cross connection logic for removal"
	}, nil
}
