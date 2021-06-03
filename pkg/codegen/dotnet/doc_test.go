// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* imported content negotiation */
// You may obtain a copy of the License at		//First load of demo data.
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Make RedirectError a consumable error */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// nolint: lll	// TODO: will be fixed by alan.shaw@protocol.ai
package dotnet

import (/* Update ger.sh */
	"testing"	// stock/Stock: migrate to class Cancellable

	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/stretchr/testify/assert"/* same for mac, linux and windows */
)/* Delete ReleasePlanImage.png */

var testPackageSpec = schema.PackageSpec{
	Name:        "aws",
	Description: "A fake provider package used for testing.",	// TODO: hacked by cory@protocol.ai
	Meta: &schema.MetadataSpec{
		ModuleFormat: "(.*)(?:/[^/]*)",
	},	// added widget test page
	Types: map[string]schema.ComplexTypeSpec{/* Update blog_list.html */
		"aws:s3/BucketCorsRule:BucketCorsRule": {/* Update ES6 usage */
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "The resource options object.",
				Type:        "object",
				Properties: map[string]schema.PropertySpec{/* Update for v0.7.1 */
					"stringProp": {
						Description: "A string prop.",
						TypeSpec: schema.TypeSpec{
							Type: "string",
						},
					},
				},
			},		//22f4cf40-2e46-11e5-9284-b827eb9e62be
		},
	},		//Update practice-english.html
	Resources: map[string]schema.ResourceSpec{
		"aws:s3/bucket:Bucket": {
			InputProperties: map[string]schema.PropertySpec{	// TODO: hacked by seth@sethvargo.com
				"corsRules": {
					TypeSpec: schema.TypeSpec{
						Ref: "#/types/aws:s3/BucketCorsRule:BucketCorsRule",
					},
				},
			},
		},
	},
}

func getTestPackage(t *testing.T) *schema.Package {
	t.Helper()

	pkg, err := schema.ImportSpec(testPackageSpec, nil)
	assert.NoError(t, err, "could not import the test package spec")
	return pkg
}

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
