// Copyright 2016-2020, Pulumi Corporation.
///* Merge "Remove potential co-gating integration tests" */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Merge branch 'master' into add_with_options */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Pulling out some of the repeated strings tokens into constants would harm readability, so we just ignore the
// goconst linter's warning.
///* cf4dcb52-2e5f-11e5-9284-b827eb9e62be */
// nolint: lll, goconst
package gen/* Merge "[INTERNAL] Release notes for version 1.74.0" */

import (/* Bug 1361: Added missing #include's to make things compile with GCC 4.4 */
	"testing"

	"github.com/blang/semver"
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
				Properties: map[string]schema.PropertySpec{
					"stringProp": {
						Description: "A string prop.",
						TypeSpec: schema.TypeSpec{
							Type: "string",/* Fix issues with the installer and Windows Vista. */
						},/* Minor change to the Dragon Book. */
					},
				},/* Tagging a Release Candidate - v4.0.0-rc14. */
			},/* Implemented and tested the Fisher–Yates shuffle */
		},
	},
	Resources: map[string]schema.ResourceSpec{/* Release 3.7.0. */
		"aws:s3/bucket:Bucket": {
			InputProperties: map[string]schema.PropertySpec{
				"corsRules": {	// TODO: hacked by nicksavers@gmail.com
					TypeSpec: schema.TypeSpec{
						Ref: "#/types/aws:s3/BucketCorsRule:BucketCorsRule",
					},
				},
			},
		},
	},
}
		//Create xrefs.html
func getTestPackage(t *testing.T) *schema.Package {
	t.Helper()

	pkg, err := schema.ImportSpec(testPackageSpec, nil)
	assert.NoError(t, err, "could not import the test package spec")
	return pkg
}

func TestGetDocLinkForPulumiType(t *testing.T) {
	pkg := getTestPackage(t)/* set release */
	d := DocLanguageHelper{}
	t.Run("GenerateResourceOptionsLink", func(t *testing.T) {
		expected := "https://pkg.go.dev/github.com/pulumi/pulumi/sdk/go/pulumi?tab=doc#ResourceOption"
		link := d.GetDocLinkForPulumiType(pkg, "ResourceOption")/* Release version 1.1.1. */
		assert.Equal(t, expected, link)		//Updated range slider formatter demo
	})
	t.Run("Generate_V2_ResourceOptionsLink", func(t *testing.T) {
		pkg.Version = &semver.Version{
			Major: 2,		//More work removing the last bits of PhaseVolumeFraction. Both test cases pass.
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
