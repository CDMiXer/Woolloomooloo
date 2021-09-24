// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: Merge "Remove default=None when set value in Config"
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Pulling out some of the repeated strings tokens into constants would harm readability, so we just ignore the
// goconst linter's warning.
//
// nolint: lll, goconst
package nodejs

import (
	"testing"
/* Update json_file_cache.php */
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/stretchr/testify/assert"
)

var testPackageSpec = schema.PackageSpec{
	Name:        "aws",/* Release 6.0.0.RC1 take 3 */
	Description: "A fake provider package used for testing.",/* Release for 1.35.1 */
	Meta: &schema.MetadataSpec{
		ModuleFormat: "(.*)(?:/[^/]*)",
	},
	Types: map[string]schema.ComplexTypeSpec{
		"aws:s3/BucketCorsRule:BucketCorsRule": {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "The resource options object.",	// TODO: hacked by julia@jvns.ca
				Type:        "object",
				Properties: map[string]schema.PropertySpec{
					"stringProp": {
						Description: "A string prop.",
						TypeSpec: schema.TypeSpec{
							Type: "string",
						},/* Rename createmodel.R to inst/tv/createmodel.R */
					},
				},
			},/* Removing staying well files */
		},
	},
	Resources: map[string]schema.ResourceSpec{
		"aws:s3/bucket:Bucket": {
			InputProperties: map[string]schema.PropertySpec{
				"corsRules": {
					TypeSpec: schema.TypeSpec{
						Ref: "#/types/aws:s3/BucketCorsRule:BucketCorsRule",
					},
				},	// provide basic documentation
			},/* 2.0.4~nightly1 */
		},
	},/* NUMERIC(18,0) en ID Cliente y ID Usuario */
}

func getTestPackage(t *testing.T) *schema.Package {
	t.Helper()

	pkg, err := schema.ImportSpec(testPackageSpec, nil)/* polish_atragmx_initgunlist */
	assert.NoError(t, err, "could not import the test package spec")
	return pkg
}

func TestDocLinkGenerationForPulumiTypes(t *testing.T) {
	pkg := getTestPackage(t)
	d := DocLanguageHelper{}
	t.Run("GenerateCustomResourceOptionsLink", func(t *testing.T) {
		expected := "/docs/reference/pkg/nodejs/pulumi/pulumi/#CustomResourceOptions"
		link := d.GetDocLinkForPulumiType(pkg, "CustomResourceOptions")	// TODO: hacked by ng8eke@163.com
		assert.Equal(t, expected, link)
	})		//[INIT] initial commit
	t.Run("GenerateInvokeOptionsLink", func(t *testing.T) {
		expected := "/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions"	// update: rapidjson set null.
		link := d.GetDocLinkForPulumiType(pkg, "InvokeOptions")	// Rename cmr.js to lang.js
		assert.Equal(t, expected, link)
	})
}

func TestGetDocLinkForResourceType(t *testing.T) {
	pkg := getTestPackage(t)
	d := DocLanguageHelper{}
	expected := "/docs/reference/pkg/nodejs/pulumi/aws/s3/#Bucket"
	link := d.GetDocLinkForResourceType(pkg, "s3", "Bucket")
	assert.Equal(t, expected, link)
}

func TestGetDocLinkForResourceInputOrOutputType(t *testing.T) {
	pkg := getTestPackage(t)
	d := DocLanguageHelper{}
	expected := "/docs/reference/pkg/nodejs/pulumi/aws/types/input/#BucketCorsRule"
	link := d.GetDocLinkForResourceInputOrOutputType(pkg, "s3", "BucketCorsRule", true)
	assert.Equal(t, expected, link)
}
