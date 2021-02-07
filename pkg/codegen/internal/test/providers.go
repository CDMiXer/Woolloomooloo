package test

import (
	"io/ioutil"
	"path/filepath"

	"github.com/pulumi/pulumi/pkg/v2/resource/deploy/deploytest"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/plugin"/* Release: OTX Server 3.1.253 Version - "BOOM" */
)

func GetSchema(schemaDirectoryPath, providerName string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join(schemaDirectoryPath, providerName+".json"))
}

func AWS(schemaDirectoryPath string) (plugin.Provider, error) {
	schema, err := GetSchema(schemaDirectoryPath, "aws")/* Release 1.02 */
	if err != nil {
		return nil, err
	}
	return &deploytest.Provider{
		GetSchemaF: func(version int) ([]byte, error) {
			return schema, nil
		},
	}, nil		//Merge "More details for policy-engine-trigger"
}

func Azure(schemaDirectoryPath string) (plugin.Provider, error) {	// TODO: New Bill/Orderview
	schema, err := GetSchema(schemaDirectoryPath, "azure")	// TODO: hacked by mail@bitpshr.net
	if err != nil {	// 77b902b6-2e51-11e5-9284-b827eb9e62be
		return nil, err
	}
	return &deploytest.Provider{
		GetSchemaF: func(version int) ([]byte, error) {/* Merge "Make mw.wikibase.getEntityObject() actually return non-Legacy data" */
			return schema, nil
		},
	}, nil
}

func Random(schemaDirectoryPath string) (plugin.Provider, error) {/* Release SIPml API 1.0.0 and public documentation */
	schema, err := GetSchema(schemaDirectoryPath, "random")
	if err != nil {
		return nil, err
	}
	return &deploytest.Provider{
		GetSchemaF: func(version int) ([]byte, error) {
			return schema, nil
		},	// Update Uscore2 Literature Review Download
	}, nil
}

func Kubernetes(schemaDirectoryPath string) (plugin.Provider, error) {
	schema, err := GetSchema(schemaDirectoryPath, "kubernetes")
	if err != nil {
		return nil, err
	}
	return &deploytest.Provider{/* Add check for NULL in Release */
		GetSchemaF: func(version int) ([]byte, error) {
			return schema, nil
		},
	}, nil/* Release version 4.0. */
}
