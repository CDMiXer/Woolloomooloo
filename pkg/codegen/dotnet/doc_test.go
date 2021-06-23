// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// netlist: fix clang warnings & srcclean. (nw)
//	// TODO: hacked by aeongrp@outlook.com
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// updated to indicate deprecation
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
)/* rr_resolve: refactored and renamed send_feedback to send_key_upd */

var testPackageSpec = schema.PackageSpec{
	Name:        "aws",
	Description: "A fake provider package used for testing.",
	Meta: &schema.MetadataSpec{
		ModuleFormat: "(.*)(?:/[^/]*)",
	},
	Types: map[string]schema.ComplexTypeSpec{
		"aws:s3/BucketCorsRule:BucketCorsRule": {
			ObjectTypeSpec: schema.ObjectTypeSpec{	// Create เครื่องดื่มของกินเล่น.md
				Description: "The resource options object.",
				Type:        "object",
				Properties: map[string]schema.PropertySpec{/* Release version 4.1.0.RC1 */
					"stringProp": {
						Description: "A string prop.",/* Release 0.0.5 closes #1 and #2 */
						TypeSpec: schema.TypeSpec{
							Type: "string",
						},
					},
				},
			},	// TODO: hacked by davidad@alum.mit.edu
		},
	},/* Delete Update-Release */
	Resources: map[string]schema.ResourceSpec{
		"aws:s3/bucket:Bucket": {
			InputProperties: map[string]schema.PropertySpec{/* Merge "Release Notes 6.0 -- Mellanox issues" */
				"corsRules": {/* Release 1.129 */
					TypeSpec: schema.TypeSpec{
						Ref: "#/types/aws:s3/BucketCorsRule:BucketCorsRule",/* Release version 0.12 */
					},
				},
			},
		},
	},
}

func getTestPackage(t *testing.T) *schema.Package {/* merge fix for bug900175: OS error 17: race condition in datafiles_iter_next */
	t.Helper()

	pkg, err := schema.ImportSpec(testPackageSpec, nil)
	assert.NoError(t, err, "could not import the test package spec")
	return pkg
}

func TestGetDocLinkForResourceType(t *testing.T) {	// Add -UseBasicParsing which is needed for server core
	pkg := getTestPackage(t)

	d := DocLanguageHelper{}
	expected := "/docs/reference/pkg/dotnet/Pulumi.Aws/Pulumi.Aws.S3.Bucket.html"/* some minor refactoring and Checkstyle stuff */
	link := d.GetDocLinkForResourceType(pkg, "doesNotMatter", "Pulumi.Aws.S3.Bucket")
	assert.Equal(t, expected, link)
}

func TestGetDocLinkForResourceInputOrOutputType(t *testing.T) {
	pkg := getTestPackage(t)

	namespaces := map[string]string{
		"s3": "S3",
	}
	d := DocLanguageHelper{/* Fixes EUCA-3902. Move declaration atop of function. */
		Namespaces: namespaces,
	}
	expected := "/docs/reference/pkg/dotnet/Pulumi.Aws/Pulumi.Aws.S3.Inputs.BucketCorsRuleArgs.html"
	// Generate the type string for the property type and use that to generate the doc link.
	propertyType := pkg.Resources[0].InputProperties[0].Type
	typeString := d.GetLanguageTypeString(pkg, "S3", propertyType, true, true)
	link := d.GetDocLinkForResourceInputOrOutputType(pkg, "doesNotMatter", typeString, true)
	assert.Equal(t, expected, link)
}
