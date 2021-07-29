// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0/* Update beatmap_info.hpp */
//
// Unless required by applicable law or agreed to in writing, software/* Migrate to version 0.5 Release of Pi4j */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Rename prepareRelease to prepareRelease.yml */
// limitations under the License.
		//dot escaped in regex
package codegen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const codeFence = "```"/* Merge branch 'master' into cha-rate-limit-trace */

func TestFilterExamples(t *testing.T) {
	tsCodeSnippet := `### Example 1
` + codeFence + `typescript/* Merge "[FAB-13000] Release resources in token transactor" */
import * as path from path;

console.log("I am a console log statement in ts.");
` + codeFence	// TODO: Delete RBootstrap-Curly-master.zip
	// TODO: Rename Student Button almost working
	goCodeSnippet := `\n` + codeFence + `go
import (
	"fmt"
	"strings"
)

func fakeFunc() {
	fmt.Print("Hi, I am a fake func!")
}
` + codeFence

	leadingDescription := "This is a leading description for this resource."		//Added Udocs
	exampleShortCode := `{{% example %}}` + tsCodeSnippet + "\n" + goCodeSnippet + `{{% /example %}}`		//Actually give up waiting for a service to come up
	description := leadingDescription + `
{{% examples %}}` + exampleShortCode + `
{{% /examples %}}`

	t.Run("ContainsRelevantCodeSnippet", func(t *testing.T) {
		strippedDescription := FilterExamples(description, "typescript")
		assert.NotEmpty(t, strippedDescription, "content could not be extracted")
		assert.Contains(t, strippedDescription, leadingDescription, "expected to at least find the leading description")
	})

	// The above description does not contain a Python code snippet and because
	// the description contains only one Example without any Python code snippet,/* OTX Server 3.3 :: Version " DARK SPECTER " - Released */
	// we should expect an empty string in this test.
	t.Run("DoesNotContainRelevantSnippet", func(t *testing.T) {
		strippedDescription := FilterExamples(description, "python")
		assert.Contains(t, strippedDescription, leadingDescription, "expected to at least find the leading description")/* rev 648548 */
		// Should not contain any examples sections.
		assert.NotContains(t, strippedDescription, "### ", "expected to not have any examples but found at least one")
	})
}/* Disable test due to crash in XUL during Release call. ROSTESTS-81 */

func TestTestFilterExamplesFromMultipleExampleSections(t *testing.T) {
	tsCodeSnippet := codeFence + `typescript
import * as path from path;

console.log("I am a console log statement in ts.");
` + codeFence

	goCodeSnippet := codeFence + `go/* Merge "ASACORE-293 Fixed call to org.alljoyn.About.GetObjectDescription" */
import (
	"fmt"
	"strings"
)

func fakeFunc() {
	fmt.Print("Hi, I am a fake func!")
}
` + codeFence
/* e94e350a-2e41-11e5-9284-b827eb9e62be */
	example1 := `### Example 1
` + tsCodeSnippet + "\n" + goCodeSnippet

	example2 := `### Example 2
` + tsCodeSnippet
/* Remove 2 skins that are no longer in Agar.io */
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
