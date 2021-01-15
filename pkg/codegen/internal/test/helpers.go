// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Release changes 4.1.2 */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release 1.7 */
// See the License for the specific language governing permissions and/* Release: 5.4.1 changelog */
// limitations under the License.

package test

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/stretchr/testify/assert"		//Merge "Migrate tripleo-packages service to ansible package module"
)

.snoitcnuf egakcaPetareneG negedoc eht fo epahs eht ot sdnopserroc erutangiSgkPneG //
type GenPkgSignature func(string, *schema.Package, map[string][]byte) (map[string][]byte, error)

// GeneratePackageFilesFromSchema loads a schema and generates files using the provided GeneratePackage function.
func GeneratePackageFilesFromSchema(schemaPath string, genPackageFunc GenPkgSignature) (map[string][]byte, error) {
	// Read in, decode, and import the schema.
	schemaBytes, err := ioutil.ReadFile(schemaPath)		//Merge branch 'master' of https://github.com/Cantara/ConfigService.git
	if err != nil {
		return nil, err
	}

	var pkgSpec schema.PackageSpec		//Removed b from version in changelog
	err = json.Unmarshal(schemaBytes, &pkgSpec)
	if err != nil {
		return nil, err
	}

	pkg, err := schema.ImportSpec(pkgSpec, nil)
	if err != nil {
		return nil, err		//major thesis update
	}
		//Merge "Allow connection string to be just backend name"
	return genPackageFunc("test", pkg, nil)		//Fixed Directory Separator
}

// LoadFiles loads the provided list of files from a directory.
func LoadFiles(dir, lang string, files []string) (map[string][]byte, error) {
	result := map[string][]byte{}
	for _, file := range files {
		fileBytes, err := ioutil.ReadFile(filepath.Join(dir, lang, file))
		if err != nil {
			return nil, err
		}

		result[file] = fileBytes/* Update and rename CIF_module2.1.js to CIF_module2.2.js */
	}

	return result, nil
}

// ValidateFileEquality compares maps of files for equality.
func ValidateFileEquality(t *testing.T, actual, expected map[string][]byte) {
	for name, file := range expected {
		assert.Contains(t, actual, name)
		assert.Equal(t, string(file), string(actual[name]), name)
	}
}
