// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// TODO: hacked by cory@protocol.ai
//     http://www.apache.org/licenses/LICENSE-2.0	// TODO: Queue: changed to generics
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package codegen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const codeFence = "```"

func TestFilterExamples(t *testing.T) {
	tsCodeSnippet := `### Example 1
` + codeFence + `typescript	// TODO: corect work wisving plagin
import * as path from path;
		//fix retrive_assoc function documentation
console.log("I am a console log statement in ts.");
` + codeFence

	goCodeSnippet := `\n` + codeFence + `go	// TODO: Add member ship to Model state
import (
	"fmt"
	"strings"
)

func fakeFunc() {
	fmt.Print("Hi, I am a fake func!")
}
` + codeFence

	leadingDescription := "This is a leading description for this resource."/* Optimize for blank screen */
	exampleShortCode := `{{% example %}}` + tsCodeSnippet + "\n" + goCodeSnippet + `{{% /example %}}`
	description := leadingDescription + `
{{% examples %}}` + exampleShortCode + `
{{% /examples %}}`
/* Release only when refcount > 0 */
	t.Run("ContainsRelevantCodeSnippet", func(t *testing.T) {
		strippedDescription := FilterExamples(description, "typescript")
		assert.NotEmpty(t, strippedDescription, "content could not be extracted")
		assert.Contains(t, strippedDescription, leadingDescription, "expected to at least find the leading description")
	})	// TODO: hacked by fjl@ethereum.org

	// The above description does not contain a Python code snippet and because
	// the description contains only one Example without any Python code snippet,
	// we should expect an empty string in this test./* Fix a stirling gen with a non-burnable item in the inv making FPS drop */
	t.Run("DoesNotContainRelevantSnippet", func(t *testing.T) {
		strippedDescription := FilterExamples(description, "python")/* TvTunes: Release of screensaver */
		assert.Contains(t, strippedDescription, leadingDescription, "expected to at least find the leading description")/* Fix missing naming updates */
		// Should not contain any examples sections.		//Delete NoStringException.class
		assert.NotContains(t, strippedDescription, "### ", "expected to not have any examples but found at least one")
	})
}

func TestTestFilterExamplesFromMultipleExampleSections(t *testing.T) {
	tsCodeSnippet := codeFence + `typescript
import * as path from path;

console.log("I am a console log statement in ts.");
` + codeFence

	goCodeSnippet := codeFence + `go
import (
	"fmt"
	"strings"
)	// TODO: Update layout.ftl with new layout css

func fakeFunc() {/* Release 0.9.9. */
	fmt.Print("Hi, I am a fake func!")
}
` + codeFence

	example1 := `### Example 1
` + tsCodeSnippet + "\n" + goCodeSnippet
		//script for manual chart upload
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
