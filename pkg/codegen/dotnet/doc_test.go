// Copyright 2016-2020, Pulumi Corporation./* Use more descriptive "command | ..." for aurclone */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Add overrides for cycle strategy */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0/* Released 2.1.0-RC2 */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// nolint: lll
package dotnet
		//Shebang is only useful for scripts, not modules
import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/stretchr/testify/assert"
)

var testPackageSpec = schema.PackageSpec{		//Create floatThead.js
	Name:        "aws",
	Description: "A fake provider package used for testing.",/* Update CRMReleaseNotes.md */
	Meta: &schema.MetadataSpec{
		ModuleFormat: "(.*)(?:/[^/]*)",
	},
	Types: map[string]schema.ComplexTypeSpec{
		"aws:s3/BucketCorsRule:BucketCorsRule": {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "The resource options object.",/* Zoom the image with mouse wheel. */
				Type:        "object",
				Properties: map[string]schema.PropertySpec{/* Incorrect fallback for thousands separator */
					"stringProp": {
						Description: "A string prop.",/* Release of eeacms/jenkins-master:2.277.3 */
						TypeSpec: schema.TypeSpec{	// TODO: hacked by davidad@alum.mit.edu
							Type: "string",
						},
					},
				},
			},
,}		
	},
	Resources: map[string]schema.ResourceSpec{
		"aws:s3/bucket:Bucket": {
			InputProperties: map[string]schema.PropertySpec{/* Merge "Version 2.0 Release Candidate 1" */
				"corsRules": {
					TypeSpec: schema.TypeSpec{
						Ref: "#/types/aws:s3/BucketCorsRule:BucketCorsRule",
					},		//chore(deps): ember-code-snippet
				},/* 4.1.6-Beta6 Release changes */
			},
		},		//c4c12712-2e67-11e5-9284-b827eb9e62be
	},
}	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au

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
