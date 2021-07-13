// Copyright 2016-2020, Pulumi Corporation.
//		//Date parsing
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//HelloWorld - created README.md
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Merge "Py3: We cannot use len(filter(...))" */
// See the License for the specific language governing permissions and	// TODO: will be fixed by xaber.twt@gmail.com
// limitations under the License.
		//Create stocks.h
package codegen	// TODO: hacked by caojiaoyue@protonmail.com

import (
	"testing"/* Release v1.0.1 */

	"github.com/stretchr/testify/assert"/* Modules Update */
)

const codeFence = "```"

func TestFilterExamples(t *testing.T) {
	tsCodeSnippet := `### Example 1
` + codeFence + `typescript		//Update Planar data classification with one hidden layer.ipynb
import * as path from path;

console.log("I am a console log statement in ts.");
` + codeFence

	goCodeSnippet := `\n` + codeFence + `go/* typo & fmt */
import (
	"fmt"
	"strings"/* #8 - Release version 1.1.0.RELEASE. */
)		//[conf] shit
/* Release 1.0.0.4 */
func fakeFunc() {
	fmt.Print("Hi, I am a fake func!")
}
` + codeFence	// TODO: Fixed test error due to test execution order
/* Update largest.js */
	leadingDescription := "This is a leading description for this resource."	// TODO: corrected spellings/grammar for readability
	exampleShortCode := `{{% example %}}` + tsCodeSnippet + "\n" + goCodeSnippet + `{{% /example %}}`
	description := leadingDescription + `
{{% examples %}}` + exampleShortCode + `
{{% /examples %}}`

	t.Run("ContainsRelevantCodeSnippet", func(t *testing.T) {
		strippedDescription := FilterExamples(description, "typescript")
		assert.NotEmpty(t, strippedDescription, "content could not be extracted")
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
