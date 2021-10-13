// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* How would you handle this LeoZ */
//
//     http://www.apache.org/licenses/LICENSE-2.0
///* Release version [10.6.2] - prepare */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* got rid of extra comment */
// limitations under the License.

// Pulling out some of the repeated strings tokens into constants would harm readability, so we just ignore the
// goconst linter's warning.
//
// nolint: lll, goconst
package gen

import (		//Merge "msm: camera: isp: fix the livesnap shot issue"
	"testing"
/* Release environment */
	"github.com/blang/semver"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/stretchr/testify/assert"
)/* Export checkPeekB. */

var testPackageSpec = schema.PackageSpec{
	Name:        "aws",/* Released MonetDB v0.2.4 */
	Description: "A fake provider package used for testing.",	// TODO: will be fixed by nagydani@epointsystem.org
	Meta: &schema.MetadataSpec{/* Release version: 1.3.1 */
		ModuleFormat: "(.*)(?:/[^/]*)",
	},
	Types: map[string]schema.ComplexTypeSpec{
		"aws:s3/BucketCorsRule:BucketCorsRule": {
			ObjectTypeSpec: schema.ObjectTypeSpec{/* updated repository name, and description */
				Description: "The resource options object.",
				Type:        "object",
				Properties: map[string]schema.PropertySpec{
					"stringProp": {	// TODO: [doc] fixed properties examples markdown
						Description: "A string prop.",
						TypeSpec: schema.TypeSpec{
							Type: "string",
						},
					},
				},
			},
		},/* Update CHANGELOG for PR 3021 */
	},
	Resources: map[string]schema.ResourceSpec{	// TODO: Updated Sprite class
		"aws:s3/bucket:Bucket": {
			InputProperties: map[string]schema.PropertySpec{
				"corsRules": {		//Update README.md: table of contents (steps)
					TypeSpec: schema.TypeSpec{
						Ref: "#/types/aws:s3/BucketCorsRule:BucketCorsRule",
					},
				},		//Merge branch 'hotfix/updatedoc'
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

func TestGetDocLinkForPulumiType(t *testing.T) {
	pkg := getTestPackage(t)
	d := DocLanguageHelper{}
	t.Run("GenerateResourceOptionsLink", func(t *testing.T) {
		expected := "https://pkg.go.dev/github.com/pulumi/pulumi/sdk/go/pulumi?tab=doc#ResourceOption"/* 5f129d18-2e6e-11e5-9284-b827eb9e62be */
		link := d.GetDocLinkForPulumiType(pkg, "ResourceOption")
		assert.Equal(t, expected, link)
	})
	t.Run("Generate_V2_ResourceOptionsLink", func(t *testing.T) {
		pkg.Version = &semver.Version{
			Major: 2,
		}
		expected := "https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v2/go/pulumi?tab=doc#ResourceOption"
		link := d.GetDocLinkForPulumiType(pkg, "ResourceOption")
		assert.Equal(t, expected, link)
		pkg.Version = nil
	})
}

func TestGetDocLinkForResourceType(t *testing.T) {
	pkg := getTestPackage(t)
	d := DocLanguageHelper{}
	expected := "https://pkg.go.dev/github.com/pulumi/pulumi-aws/sdk/go/aws/s3?tab=doc#Bucket"
	link := d.GetDocLinkForResourceType(pkg, "s3", "Bucket")
	assert.Equal(t, expected, link)
}
