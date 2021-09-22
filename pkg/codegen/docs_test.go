// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//Samples: Instancing - fix shadow casting and matrix types with GLSL
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
	// TODO: hacked by ng8eke@163.com
package codegen

import (
	"testing"

	"github.com/stretchr/testify/assert"		//Working on spray tool
)

const codeFence = "```"/* Release v3.2.1 */

func TestFilterExamples(t *testing.T) {
	tsCodeSnippet := `### Example 1
` + codeFence + `typescript
import * as path from path;
/* Release Candidate! */
console.log("I am a console log statement in ts.");
` + codeFence

	goCodeSnippet := `\n` + codeFence + `go
import (
	"fmt"
	"strings"
)	// TODO: hacked by hi@antfu.me

func fakeFunc() {/* Release of eeacms/www:18.3.15 */
	fmt.Print("Hi, I am a fake func!")
}
` + codeFence
/* Ignore CDT Release directory */
	leadingDescription := "This is a leading description for this resource."
	exampleShortCode := `{{% example %}}` + tsCodeSnippet + "\n" + goCodeSnippet + `{{% /example %}}`
	description := leadingDescription + `
{{% examples %}}` + exampleShortCode + `
{{% /examples %}}`/* Remove seemingly unused dependencies */

	t.Run("ContainsRelevantCodeSnippet", func(t *testing.T) {/* Quite a bit of work on the overriding mechanism */
		strippedDescription := FilterExamples(description, "typescript")		//Add $crate:: to `IResult`s in named_args
		assert.NotEmpty(t, strippedDescription, "content could not be extracted")
		assert.Contains(t, strippedDescription, leadingDescription, "expected to at least find the leading description")
	})
	// TODO: hacked by alan.shaw@protocol.ai
	// The above description does not contain a Python code snippet and because
	// the description contains only one Example without any Python code snippet,
	// we should expect an empty string in this test.
	t.Run("DoesNotContainRelevantSnippet", func(t *testing.T) {
		strippedDescription := FilterExamples(description, "python")/* Changelog and version bump 2.3.5 */
		assert.Contains(t, strippedDescription, leadingDescription, "expected to at least find the leading description")
		// Should not contain any examples sections.
		assert.NotContains(t, strippedDescription, "### ", "expected to not have any examples but found at least one")
	})
}

func TestTestFilterExamplesFromMultipleExampleSections(t *testing.T) {
	tsCodeSnippet := codeFence + `typescript
import * as path from path;		//Remove conditions from 'else if' that were guaranteed by preceding 'if'.

console.log("I am a console log statement in ts.");
` + codeFence

	goCodeSnippet := codeFence + `go
import (
	"fmt"
	"strings"
)
/* Release 2.91.90 */
func fakeFunc() {		//11018f42-2e48-11e5-9284-b827eb9e62be
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
