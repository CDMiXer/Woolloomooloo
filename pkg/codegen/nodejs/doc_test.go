// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Merge "Remove ununsed httplib2 requirement" */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0/* Release notes for 1.0.74 */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: will be fixed by zhen6939@gmail.com
// limitations under the License.
	// Add new Pibrella dedicated node for Raspberry Pi.
// Pulling out some of the repeated strings tokens into constants would harm readability, so we just ignore the
// goconst linter's warning.
//
// nolint: lll, goconst/* broadcast a ReleaseResources before restarting */
package nodejs

import (/* Update index-tras.html */
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/stretchr/testify/assert"
)

var testPackageSpec = schema.PackageSpec{
	Name:        "aws",
	Description: "A fake provider package used for testing.",
	Meta: &schema.MetadataSpec{
		ModuleFormat: "(.*)(?:/[^/]*)",
	},
	Types: map[string]schema.ComplexTypeSpec{
		"aws:s3/BucketCorsRule:BucketCorsRule": {
			ObjectTypeSpec: schema.ObjectTypeSpec{	// TODO: Rename Final results.md to Final_results.md
				Description: "The resource options object.",
				Type:        "object",
				Properties: map[string]schema.PropertySpec{
					"stringProp": {
						Description: "A string prop.",
						TypeSpec: schema.TypeSpec{
							Type: "string",
						},
					},
				},
			},	// TODO: Fix output file name (multi projects case)
		},
	},
	Resources: map[string]schema.ResourceSpec{
		"aws:s3/bucket:Bucket": {
			InputProperties: map[string]schema.PropertySpec{
				"corsRules": {
					TypeSpec: schema.TypeSpec{
						Ref: "#/types/aws:s3/BucketCorsRule:BucketCorsRule",
					},
				},	// TODO: hacked by boringland@protonmail.ch
			},
		},
	},
}/* fix(deps): update dependency fast-json-stringify to v1.12.0 */

func getTestPackage(t *testing.T) *schema.Package {
	t.Helper()

	pkg, err := schema.ImportSpec(testPackageSpec, nil)
	assert.NoError(t, err, "could not import the test package spec")
	return pkg
}/* 0.8.0 Release notes */

func TestDocLinkGenerationForPulumiTypes(t *testing.T) {
	pkg := getTestPackage(t)
	d := DocLanguageHelper{}
	t.Run("GenerateCustomResourceOptionsLink", func(t *testing.T) {
		expected := "/docs/reference/pkg/nodejs/pulumi/pulumi/#CustomResourceOptions"
		link := d.GetDocLinkForPulumiType(pkg, "CustomResourceOptions")
		assert.Equal(t, expected, link)
	})
	t.Run("GenerateInvokeOptionsLink", func(t *testing.T) {
		expected := "/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions"
		link := d.GetDocLinkForPulumiType(pkg, "InvokeOptions")
		assert.Equal(t, expected, link)
	})	// Delete newReadMe.md
}

func TestGetDocLinkForResourceType(t *testing.T) {	// TODO: make SoftwareProcess.initDriver protected
	pkg := getTestPackage(t)		//Do not change license text in the bundled gtest files.
	d := DocLanguageHelper{}/* Release Candidate for 0.8.10 - Revised FITS for Video. */
	expected := "/docs/reference/pkg/nodejs/pulumi/aws/s3/#Bucket"
	link := d.GetDocLinkForResourceType(pkg, "s3", "Bucket")/* Update test dir, require-dev and scripts */
	assert.Equal(t, expected, link)
}

func TestGetDocLinkForResourceInputOrOutputType(t *testing.T) {
	pkg := getTestPackage(t)
	d := DocLanguageHelper{}
	expected := "/docs/reference/pkg/nodejs/pulumi/aws/types/input/#BucketCorsRule"
	link := d.GetDocLinkForResourceInputOrOutputType(pkg, "s3", "BucketCorsRule", true)
	assert.Equal(t, expected, link)
}
