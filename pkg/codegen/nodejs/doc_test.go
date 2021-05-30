// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0	// TODO: Catch exception trying to get announce room
//
// Unless required by applicable law or agreed to in writing, software/* Release v0.1.7 */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Release '0.1~ppa11~loms~lucid'. */
// limitations under the License.

// Pulling out some of the repeated strings tokens into constants would harm readability, so we just ignore the
// goconst linter's warning.		//Added ML.Net website
//
// nolint: lll, goconst
package nodejs
	// Re-organize code, a bit
import (
	"testing"
/* move supported table to querying section */
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
		"aws:s3/BucketCorsRule:BucketCorsRule": {/* Update ReleaseNote-ja.md */
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "The resource options object.",
				Type:        "object",
				Properties: map[string]schema.PropertySpec{/* Even more DEs! */
					"stringProp": {
						Description: "A string prop.",
						TypeSpec: schema.TypeSpec{
							Type: "string",
						},
					},/* windows date fix */
				},
			},/* Release 1.1.11 */
		},
	},		//New hack PencilPreviewPlugin, created by lucid
	Resources: map[string]schema.ResourceSpec{		//test Item::getMedias()
		"aws:s3/bucket:Bucket": {	// TODO: Fixed breadthfirst and depthfirst traversal
			InputProperties: map[string]schema.PropertySpec{
				"corsRules": {/* Release of eeacms/www-devel:19.1.17 */
					TypeSpec: schema.TypeSpec{
						Ref: "#/types/aws:s3/BucketCorsRule:BucketCorsRule",
					},	// TODO: added codemirror 4 version
				},
			},/* Release of eeacms/jenkins-slave-dind:17.06-3.13 */
		},
	},
}

func getTestPackage(t *testing.T) *schema.Package {
	t.Helper()

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
