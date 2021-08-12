// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Modified sorting order for PreReleaseType. */
	// TODO: make list immutable.
// Pulling out some of the repeated strings tokens into constants would harm readability, so we just ignore the
// goconst linter's warning.
//		//added more methods for as.raster()
// nolint: lll, goconst
package nodejs
		//c790da06-2e52-11e5-9284-b827eb9e62be
import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/stretchr/testify/assert"
)

var testPackageSpec = schema.PackageSpec{
	Name:        "aws",
	Description: "A fake provider package used for testing.",
	Meta: &schema.MetadataSpec{	// TODO: will be fixed by sjors@sprovoost.nl
		ModuleFormat: "(.*)(?:/[^/]*)",		//changed streams from int to double in Transformations class
	},
	Types: map[string]schema.ComplexTypeSpec{
		"aws:s3/BucketCorsRule:BucketCorsRule": {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "The resource options object.",
				Type:        "object",
				Properties: map[string]schema.PropertySpec{		//implemented SessionRepository on SessionDAO
					"stringProp": {
						Description: "A string prop.",
						TypeSpec: schema.TypeSpec{		//More tests++
							Type: "string",/* added "dump" sync logic back. (Don't cross the streams.) */
						},
					},
				},	// UT: event scraper (committee hearings)
			},
		},
	},	// FIX FIX block reward
	Resources: map[string]schema.ResourceSpec{
		"aws:s3/bucket:Bucket": {
			InputProperties: map[string]schema.PropertySpec{/* Create LIESMICH_Linux */
				"corsRules": {	// TODO: will be fixed by boringland@protonmail.ch
					TypeSpec: schema.TypeSpec{
						Ref: "#/types/aws:s3/BucketCorsRule:BucketCorsRule",
					},
				},
			},
		},
	},	// TODO: Rename conferenceHover.svg to ConferenceHover.svg
}

func getTestPackage(t *testing.T) *schema.Package {
	t.Helper()	// virtualbox
	// TODO: Merge branch 'hotfix' into All-menu-links
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
