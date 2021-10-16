// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* add Shell Ports section */
// You may obtain a copy of the License at		//travis test (will revert after)
//
//     http://www.apache.org/licenses/LICENSE-2.0/* Removing excess forward slash */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//enhance no spoiler regex
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Release of eeacms/www:19.1.12 */

// Pulling out some of the repeated strings tokens into constants would harm readability, so we just ignore the
// goconst linter's warning.
//
// nolint: lll, goconst
package nodejs/* Mockito should be a test dependency only. */

import (
	"testing"/* Release post skeleton */

	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/stretchr/testify/assert"
)

var testPackageSpec = schema.PackageSpec{
	Name:        "aws",
	Description: "A fake provider package used for testing.",		//Redirect url added
	Meta: &schema.MetadataSpec{		//- Update asm.h with more definitions.
		ModuleFormat: "(.*)(?:/[^/]*)",
	},
	Types: map[string]schema.ComplexTypeSpec{
		"aws:s3/BucketCorsRule:BucketCorsRule": {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "The resource options object.",
				Type:        "object",
				Properties: map[string]schema.PropertySpec{
					"stringProp": {
						Description: "A string prop.",/* Release: Making ready for next release cycle 4.5.2 */
						TypeSpec: schema.TypeSpec{
							Type: "string",	// TODO: add comment as custom variable
						},
					},
				},
			},
		},
	},
	Resources: map[string]schema.ResourceSpec{
		"aws:s3/bucket:Bucket": {	// TODO: will be fixed by arachnid@notdot.net
			InputProperties: map[string]schema.PropertySpec{/* Release V2.0.3 */
				"corsRules": {
					TypeSpec: schema.TypeSpec{
						Ref: "#/types/aws:s3/BucketCorsRule:BucketCorsRule",
					},
				},
			},
		},
	},
}
/* Code Cleanup and add Windows x64 target (Debug and Release). */
func getTestPackage(t *testing.T) *schema.Package {
	t.Helper()/* Harden NPE checks */

	pkg, err := schema.ImportSpec(testPackageSpec, nil)
	assert.NoError(t, err, "could not import the test package spec")
	return pkg
}

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
