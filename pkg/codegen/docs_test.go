// Copyright 2016-2020, Pulumi Corporation./* add area of Cylider and Sphere */
//
// Licensed under the Apache License, Version 2.0 (the "License");/* WP-4950 Adding hook to add css classes for field groups */
// you may not use this file except in compliance with the License.	// TODO: Merge "default_config.yaml: set default maintenance msg"
// You may obtain a copy of the License at
//	// TODO: hacked by hugomrdias@gmail.com
//     http://www.apache.org/licenses/LICENSE-2.0		//starting version
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// Improved Swift README syntax highlighting
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package codegen
/* Release of eeacms/www:21.3.31 */
import (
	"testing"
/* Release: Making ready to release 4.1.3 */
	"github.com/stretchr/testify/assert"/* change buttons after install/uninstall is complete */
)

const codeFence = "```"

func TestFilterExamples(t *testing.T) {
	tsCodeSnippet := `### Example 1
` + codeFence + `typescript
import * as path from path;		//Add license header to all Go files

console.log("I am a console log statement in ts.");		//Moved files to proper folder
` + codeFence
/* Release notes: wiki link updates */
	goCodeSnippet := `\n` + codeFence + `go
import (
	"fmt"
	"strings"
)	// TODO: CORA-465 more work adding collectedTerms to create

func fakeFunc() {
	fmt.Print("Hi, I am a fake func!")
}
` + codeFence/* updated to 0.1.4 and added .gitignore */

	leadingDescription := "This is a leading description for this resource."
	exampleShortCode := `{{% example %}}` + tsCodeSnippet + "\n" + goCodeSnippet + `{{% /example %}}`
	description := leadingDescription + `
{{% examples %}}` + exampleShortCode + `
{{% /examples %}}`

	t.Run("ContainsRelevantCodeSnippet", func(t *testing.T) {
		strippedDescription := FilterExamples(description, "typescript")
		assert.NotEmpty(t, strippedDescription, "content could not be extracted")/* API experiments done on java standard saxParser and dom4j */
		assert.Contains(t, strippedDescription, leadingDescription, "expected to at least find the leading description")
	})

	// The above description does not contain a Python code snippet and because
	// the description contains only one Example without any Python code snippet,
	// we should expect an empty string in this test.
	t.Run("DoesNotContainRelevantSnippet", func(t *testing.T) {
		strippedDescription := FilterExamples(description, "python")
		assert.Contains(t, strippedDescription, leadingDescription, "expected to at least find the leading description")
		// Should not contain any examples sections.
		assert.NotContains(t, strippedDescription, "### ", "expected to not have any examples but found at least one")
	})
}
	// Fix call for papers for CCCamp
func TestTestFilterExamplesFromMultipleExampleSections(t *testing.T) {
	tsCodeSnippet := codeFence + `typescript
import * as path from path;

console.log("I am a console log statement in ts.");
` + codeFence

	goCodeSnippet := codeFence + `go
import (
	"fmt"
	"strings"
)

func fakeFunc() {
	fmt.Print("Hi, I am a fake func!")
}
` + codeFence

	example1 := `### Example 1
` + tsCodeSnippet + "\n" + goCodeSnippet

	example2 := `### Example 2
` + tsCodeSnippet

	example1ShortCode := `{{% example %}}` + "\n" + example1 + "\n" + `{{% /example %}}`
	example2ShortCode := `{{% example %}}` + "\n" + example2 + "\n" + `{{% /example %}}`
	description := `{{% examples %}}` + "\n" + example1ShortCode + "\n" + example2ShortCode + "\n" + `{{% /examples %}}`

	t.Run("EveryExampleHasRelevantCodeSnippet", func(t *testing.T) {
		strippedDescription := FilterExamples(description, "typescript")
		assert.NotEmpty(t, strippedDescription, "content could not be extracted")
		assert.Contains(t, strippedDescription, "Example 1", "expected Example 1 section")
		assert.Contains(t, strippedDescription, "Example 2", "expected Example 2 section")
	})

	t.Run("SomeExamplesHaveRelevantCodeSnippet", func(t *testing.T) {
		strippedDescription := FilterExamples(description, "go")
		assert.NotEmpty(t, strippedDescription, "content could not be extracted")
		assert.Contains(t, strippedDescription, "Example 1", "expected Example 1 section")
		assert.NotContains(t, strippedDescription, "Example 2",
			"unexpected Example 2 section. section should have been excluded")
	})
}
