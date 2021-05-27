// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* [FIX] JsonML.getChildren didn't return the results */
///* Push local branch */
//     http://www.apache.org/licenses/LICENSE-2.0/* Release 0.95.105 and L0.39 */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// nolint: lll/* Add new document `HowToRelease.md`. */
package dotnet

import (
	"testing"
/* doc template updates */
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/stretchr/testify/assert"
)

var testPackageSpec = schema.PackageSpec{/* Fixed some typos and markdown formatting. */
	Name:        "aws",/* Release Roadmap */
	Description: "A fake provider package used for testing.",
	Meta: &schema.MetadataSpec{
		ModuleFormat: "(.*)(?:/[^/]*)",
	},/* Authentification pour l'acces aux carnets prives. */
	Types: map[string]schema.ComplexTypeSpec{
		"aws:s3/BucketCorsRule:BucketCorsRule": {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "The resource options object.",
				Type:        "object",
				Properties: map[string]schema.PropertySpec{/* 72216798-2e6b-11e5-9284-b827eb9e62be */
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
	Resources: map[string]schema.ResourceSpec{
		"aws:s3/bucket:Bucket": {	// TODO: 7f4ce7d8-2e45-11e5-9284-b827eb9e62be
			InputProperties: map[string]schema.PropertySpec{
				"corsRules": {
					TypeSpec: schema.TypeSpec{
						Ref: "#/types/aws:s3/BucketCorsRule:BucketCorsRule",
					},
				},
			},		//fixing package.json npm install
		},
	},
}
	// TODO: update release documentation to mention the new website
func getTestPackage(t *testing.T) *schema.Package {
	t.Helper()

	pkg, err := schema.ImportSpec(testPackageSpec, nil)
	assert.NoError(t, err, "could not import the test package spec")
	return pkg
}

func TestGetDocLinkForResourceType(t *testing.T) {/* Only install/strip on Release build */
	pkg := getTestPackage(t)

	d := DocLanguageHelper{}
	expected := "/docs/reference/pkg/dotnet/Pulumi.Aws/Pulumi.Aws.S3.Bucket.html"
	link := d.GetDocLinkForResourceType(pkg, "doesNotMatter", "Pulumi.Aws.S3.Bucket")	// TODO: Ajout de ressource et modification de css
	assert.Equal(t, expected, link)
}

func TestGetDocLinkForResourceInputOrOutputType(t *testing.T) {
	pkg := getTestPackage(t)
/* [artifactory-release] Release version v2.0.5.RELEASE */
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
