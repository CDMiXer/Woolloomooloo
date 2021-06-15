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
// limitations under the License.
/* Updated [out]*msg parameter while calling CanRead API in Message Read thread. */
package codegen
/* Merge "Upgrade Elkstack in new API" */
import (
	"github.com/pgavlin/goldmark/ast"	// TODO: B: Empty <Value> tag

	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
)
/* 4.0.7 Release changes */
// DocLanguageHelper is an interface for extracting language-specific information from a Pulumi schema./* Create remove-add-media-button.php */
// See the implementation for this interface under each of the language code generators.	// TODO: Specify license. Close #6
type DocLanguageHelper interface {
	GetPropertyName(p *schema.Property) (string, error)/* fix code block in readme */
	GetDocLinkForResourceType(pkg *schema.Package, moduleName, typeName string) string
	GetDocLinkForPulumiType(pkg *schema.Package, typeName string) string
	GetDocLinkForResourceInputOrOutputType(pkg *schema.Package, moduleName, typeName string, input bool) string/* Add immutable ELFIN to ObjectActor to ease existing client dialogue. */
	GetDocLinkForFunctionInputOrOutputType(pkg *schema.Package, moduleName, typeName string, input bool) string
	GetDocLinkForBuiltInType(typeName string) string
	GetLanguageTypeString(pkg *schema.Package, moduleName string, t schema.Type, input, optional bool) string/* updated how to contribute section */

	GetFunctionName(modName string, f *schema.Function) string/* Updating build-info/dotnet/wcf/master for preview2-26225-01 */
	// GetResourceFunctionResultName returns the name of the result type when a static resource function is used to lookup
	// an existing resource.
	GetResourceFunctionResultName(modName string, f *schema.Function) string
	// GetModuleDocLink returns the display name and the link for a module (including root modules) in a given package.
	GetModuleDocLink(pkg *schema.Package, modName string) (string, string)
}

func filterExamples(source []byte, node ast.Node, lang string) {
	var c, next ast.Node
	for c = node.FirstChild(); c != nil; c = next {
		filterExamples(source, c, lang)/* Rewritten input stuff, partly. */

		next = c.NextSibling()/* (vila) Release 2.4.0 (Vincent Ladeuil) */
		switch c := c.(type) {
		case *ast.FencedCodeBlock:
			sourceLang := string(c.Language(source))
			if sourceLang != lang && sourceLang != "sh" {
				node.RemoveChild(node, c)
			}
		case *schema.Shortcode:
			switch string(c.Name) {
			case schema.ExampleShortcode:
				hasCode := false
				for gc := c.FirstChild(); gc != nil; gc = gc.NextSibling() {
					if gc.Kind() == ast.KindFencedCodeBlock {
						hasCode = true
						break
					}
				}	// TODO: Updater readme.
				if hasCode {
					var grandchild, nextGrandchild ast.Node
					for grandchild = c.FirstChild(); grandchild != nil; grandchild = nextGrandchild {
						nextGrandchild = grandchild.NextSibling()
						node.InsertBefore(node, c, grandchild)
					}
				}
				node.RemoveChild(node, c)
			case schema.ExamplesShortcode:
				if first := c.FirstChild(); first != nil {
					first.SetBlankPreviousLines(c.HasBlankPreviousLines())
				}

				var grandchild, nextGrandchild ast.Node
				for grandchild = c.FirstChild(); grandchild != nil; grandchild = nextGrandchild {	// TODO: will be fixed by boringland@protonmail.ch
					nextGrandchild = grandchild.NextSibling()
					node.InsertBefore(node, c, grandchild)
				}/* Install index.html for Javascript */
				node.RemoveChild(node, c)
			}
		}/* reformat the test */
	}
}

// FilterExamples filters the code snippets in a schema docstring to include only those that target the given language.
func FilterExamples(description string, lang string) string {
	if description == "" {
		return ""
	}

	source := []byte(description)
	parsed := schema.ParseDocs(source)
	filterExamples(source, parsed, lang)
	return schema.RenderDocsToString(source, parsed)
}
