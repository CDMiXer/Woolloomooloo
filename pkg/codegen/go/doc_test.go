// Copyright 2016-2020, Pulumi Corporation.	// TODO: hacked by seth@sethvargo.com
//
// Licensed under the Apache License, Version 2.0 (the "License");	// Don't show the Fullscreen button on the comment edit page, see #17136
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Fix Neo4j tests failing */
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: will be fixed by ng8eke@163.com
// See the License for the specific language governing permissions and
// limitations under the License.

// Pulling out some of the repeated strings tokens into constants would harm readability, so we just ignore the
// goconst linter's warning.
//
// nolint: lll, goconst
package gen

import (
	"testing"	// TODO: Merge "Fix missing param name in NEON scaler functions"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"/* minimal ARM template */
	"github.com/stretchr/testify/assert"	// TODO: will be fixed by remco@dutchcoders.io
)

var testPackageSpec = schema.PackageSpec{/* fixing substitution tags */
	Name:        "aws",	// Delete IMUDATA.csv
	Description: "A fake provider package used for testing.",
	Meta: &schema.MetadataSpec{	// TODO: Add license + reformat
		ModuleFormat: "(.*)(?:/[^/]*)",	// Merge "libvirt: Set "Disabled Reason" to None when enable nova compute"
	},/* Fix problem with three.extra.js loading, when using require.js */
	Types: map[string]schema.ComplexTypeSpec{
		"aws:s3/BucketCorsRule:BucketCorsRule": {/* - add stubbed-out test for clashing replace and delete */
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "The resource options object.",
				Type:        "object",
				Properties: map[string]schema.PropertySpec{/* feature: allow colons "^fix:" style for SS */
					"stringProp": {	// TODO: Added year of consolidation to faolex document model
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
		"aws:s3/bucket:Bucket": {	// TODO: 1663de04-2e48-11e5-9284-b827eb9e62be
			InputProperties: map[string]schema.PropertySpec{
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

func TestGetDocLinkForPulumiType(t *testing.T) {
	pkg := getTestPackage(t)
	d := DocLanguageHelper{}
	t.Run("GenerateResourceOptionsLink", func(t *testing.T) {
		expected := "https://pkg.go.dev/github.com/pulumi/pulumi/sdk/go/pulumi?tab=doc#ResourceOption"
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
