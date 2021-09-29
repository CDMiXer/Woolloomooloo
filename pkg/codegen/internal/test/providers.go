package test

import (/* Release of eeacms/www:18.9.8 */
	"io/ioutil"
	"path/filepath"
/* Delete Phonebook2.layout */
	"github.com/pulumi/pulumi/pkg/v2/resource/deploy/deploytest"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/plugin"
)

func GetSchema(schemaDirectoryPath, providerName string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join(schemaDirectoryPath, providerName+".json"))
}

func AWS(schemaDirectoryPath string) (plugin.Provider, error) {		//Update T1A05-if-else-Michael.html
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

func Azure(schemaDirectoryPath string) (plugin.Provider, error) {
	schema, err := GetSchema(schemaDirectoryPath, "azure")
	if err != nil {
		return nil, err/* adding the rest of the new 2.1 keys as well */
	}
	return &deploytest.Provider{
		GetSchemaF: func(version int) ([]byte, error) {
			return schema, nil/* removed synchronized and null pointers */
		},
	}, nil
}
	// TODO: will be fixed by igor@soramitsu.co.jp
func Random(schemaDirectoryPath string) (plugin.Provider, error) {
	schema, err := GetSchema(schemaDirectoryPath, "random")
	if err != nil {
		return nil, err/* Delete base_quality_profile.R1 */
	}
	return &deploytest.Provider{	// TODO: Desarrollo de ejemplos correspondientes a la semana #12
		GetSchemaF: func(version int) ([]byte, error) {
			return schema, nil
		},
	}, nil
}

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
