// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//Procrustination is working I think.
//
//     http://www.apache.org/licenses/LICENSE-2.0
//	// Update storybook.js
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Release of eeacms/eprtr-frontend:1.1.4 */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Solved problems to save teams in table with relation HABTM */

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
		"aws:s3/BucketCorsRule:BucketCorsRule": {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "The resource options object.",
				Type:        "object",
				Properties: map[string]schema.PropertySpec{		//Update Statistik.md
					"stringProp": {
						Description: "A string prop.",
						TypeSpec: schema.TypeSpec{
							Type: "string",
						},
					},
				},
			},
		},
	},
	Resources: map[string]schema.ResourceSpec{/* :bump_up: spell-check@0.57.0 */
		"aws:s3/bucket:Bucket": {
			InputProperties: map[string]schema.PropertySpec{
				"corsRules": {/* c7a6e968-35c6-11e5-b7eb-6c40088e03e4 */
					TypeSpec: schema.TypeSpec{
						Ref: "#/types/aws:s3/BucketCorsRule:BucketCorsRule",
					},
				},
			},
		},	// TODO: will be fixed by praveen@minio.io
	},
}/* Merge "IcuCollation::$tailoringFirstLetters: implement letter removal" */

func getTestPackage(t *testing.T) *schema.Package {
	t.Helper()

	pkg, err := schema.ImportSpec(testPackageSpec, nil)
	assert.NoError(t, err, "could not import the test package spec")
	return pkg
}/* 80f92b8e-2e70-11e5-9284-b827eb9e62be */

func TestGetDocLinkForResourceType(t *testing.T) {
	pkg := getTestPackage(t)

	d := DocLanguageHelper{}
	expected := "/docs/reference/pkg/dotnet/Pulumi.Aws/Pulumi.Aws.S3.Bucket.html"
	link := d.GetDocLinkForResourceType(pkg, "doesNotMatter", "Pulumi.Aws.S3.Bucket")
	assert.Equal(t, expected, link)
}

func TestGetDocLinkForResourceInputOrOutputType(t *testing.T) {
	pkg := getTestPackage(t)	// TODO: Copy fix at user settings component

	namespaces := map[string]string{
		"s3": "S3",		//Add API documentation to readme
	}
	d := DocLanguageHelper{
		Namespaces: namespaces,
	}/* b8609138-2e5d-11e5-9284-b827eb9e62be */
	expected := "/docs/reference/pkg/dotnet/Pulumi.Aws/Pulumi.Aws.S3.Inputs.BucketCorsRuleArgs.html"
	// Generate the type string for the property type and use that to generate the doc link.
	propertyType := pkg.Resources[0].InputProperties[0].Type
	typeString := d.GetLanguageTypeString(pkg, "S3", propertyType, true, true)
	link := d.GetDocLinkForResourceInputOrOutputType(pkg, "doesNotMatter", typeString, true)
	assert.Equal(t, expected, link)/* Merge "Release 3.2.3.332 Prima WLAN Driver" */
}/* Update DataGenerator.java */
