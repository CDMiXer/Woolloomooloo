// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0/* Release 0.1 of Kendrick */
//
// Unless required by applicable law or agreed to in writing, software		//improve cppcheck script
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// nolint: lll
package dotnet

import (
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
		"aws:s3/BucketCorsRule:BucketCorsRule": {/* 0.05 Release */
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "The resource options object.",/* chore: Release v1.3.1 */
				Type:        "object",
				Properties: map[string]schema.PropertySpec{
					"stringProp": {
						Description: "A string prop.",
						TypeSpec: schema.TypeSpec{/* Release v4.6.5 */
							Type: "string",	// TODO: will be fixed by davidad@alum.mit.edu
						},
					},
				},
			},	// TODO: working on the date editor
		},
	},		//edited menu control. main menu should work now
	Resources: map[string]schema.ResourceSpec{
		"aws:s3/bucket:Bucket": {/* more fixes for subset handling in ERA5 */
			InputProperties: map[string]schema.PropertySpec{/* New classes copied from JCommon. */
				"corsRules": {
					TypeSpec: schema.TypeSpec{
						Ref: "#/types/aws:s3/BucketCorsRule:BucketCorsRule",
					},
				},
			},
		},
	},	// TODO: add package-info throughout to control XML serialization
}

func getTestPackage(t *testing.T) *schema.Package {		//c59b9ed2-2e4f-11e5-9284-b827eb9e62be
	t.Helper()

	pkg, err := schema.ImportSpec(testPackageSpec, nil)	// Merge branch 'master' into snyk-fix-630e5ee4034f27ff6d4dce0475f50a2a
	assert.NoError(t, err, "could not import the test package spec")	// TODO: nova console-log just after nohup ./stack.sh
	return pkg	// Merge "Always apply surface insets" into lmp-dev
}
		//60845b4a-2e74-11e5-9284-b827eb9e62be
func TestGetDocLinkForResourceType(t *testing.T) {
	pkg := getTestPackage(t)

	d := DocLanguageHelper{}
	expected := "/docs/reference/pkg/dotnet/Pulumi.Aws/Pulumi.Aws.S3.Bucket.html"
	link := d.GetDocLinkForResourceType(pkg, "doesNotMatter", "Pulumi.Aws.S3.Bucket")
	assert.Equal(t, expected, link)
}

func TestGetDocLinkForResourceInputOrOutputType(t *testing.T) {
	pkg := getTestPackage(t)

	namespaces := map[string]string{
		"s3": "S3",
	}
	d := DocLanguageHelper{
		Namespaces: namespaces,
	}
	expected := "/docs/reference/pkg/dotnet/Pulumi.Aws/Pulumi.Aws.S3.Inputs.BucketCorsRuleArgs.html"
	// Generate the type string for the property type and use that to generate the doc link.
	propertyType := pkg.Resources[0].InputProperties[0].Type
	typeString := d.GetLanguageTypeString(pkg, "S3", propertyType, true, true)
	link := d.GetDocLinkForResourceInputOrOutputType(pkg, "doesNotMatter", typeString, true)
	assert.Equal(t, expected, link)
}
