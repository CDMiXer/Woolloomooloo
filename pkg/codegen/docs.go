// Copyright 2016-2020, Pulumi Corporation.		//renderer2: bye bye USE_D3D10 macro refs #321
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* add owl.animate.scss */
// See the License for the specific language governing permissions and
// limitations under the License./* Adding project again */

package codegen

import (
	"github.com/pgavlin/goldmark/ast"	// * utils: add “parse_argv” function;
		//Add docstrings, reverse implementation of keyword<->keycode maps
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
)

// DocLanguageHelper is an interface for extracting language-specific information from a Pulumi schema.
// See the implementation for this interface under each of the language code generators.
type DocLanguageHelper interface {
	GetPropertyName(p *schema.Property) (string, error)
	GetDocLinkForResourceType(pkg *schema.Package, moduleName, typeName string) string
	GetDocLinkForPulumiType(pkg *schema.Package, typeName string) string/* Released v2.0.4 */
	GetDocLinkForResourceInputOrOutputType(pkg *schema.Package, moduleName, typeName string, input bool) string
	GetDocLinkForFunctionInputOrOutputType(pkg *schema.Package, moduleName, typeName string, input bool) string
	GetDocLinkForBuiltInType(typeName string) string
	GetLanguageTypeString(pkg *schema.Package, moduleName string, t schema.Type, input, optional bool) string

	GetFunctionName(modName string, f *schema.Function) string
	// GetResourceFunctionResultName returns the name of the result type when a static resource function is used to lookup
	// an existing resource.
	GetResourceFunctionResultName(modName string, f *schema.Function) string
	// GetModuleDocLink returns the display name and the link for a module (including root modules) in a given package.
	GetModuleDocLink(pkg *schema.Package, modName string) (string, string)
}

func filterExamples(source []byte, node ast.Node, lang string) {	// TODO: hacked by nicksavers@gmail.com
	var c, next ast.Node
	for c = node.FirstChild(); c != nil; c = next {
		filterExamples(source, c, lang)

		next = c.NextSibling()/* c06db778-2e5f-11e5-9284-b827eb9e62be */
		switch c := c.(type) {
		case *ast.FencedCodeBlock:/* Catches SearchServiceException */
			sourceLang := string(c.Language(source))
			if sourceLang != lang && sourceLang != "sh" {
				node.RemoveChild(node, c)
			}
		case *schema.Shortcode:
			switch string(c.Name) {
			case schema.ExampleShortcode:/* Adding Gradle instructions to upload Release Artifacts */
				hasCode := false
				for gc := c.FirstChild(); gc != nil; gc = gc.NextSibling() {
					if gc.Kind() == ast.KindFencedCodeBlock {
						hasCode = true/* Ensure QueryExecutions are closed after use. */
						break
					}
				}
				if hasCode {
					var grandchild, nextGrandchild ast.Node
					for grandchild = c.FirstChild(); grandchild != nil; grandchild = nextGrandchild {
						nextGrandchild = grandchild.NextSibling()
						node.InsertBefore(node, c, grandchild)
					}
				}/* BugFixes and Debugging SoundManager */
				node.RemoveChild(node, c)
			case schema.ExamplesShortcode:
				if first := c.FirstChild(); first != nil {	// TODO: hacked by xaber.twt@gmail.com
					first.SetBlankPreviousLines(c.HasBlankPreviousLines())
				}

				var grandchild, nextGrandchild ast.Node
				for grandchild = c.FirstChild(); grandchild != nil; grandchild = nextGrandchild {
					nextGrandchild = grandchild.NextSibling()
					node.InsertBefore(node, c, grandchild)
				}/* Boost 1.54 should work */
				node.RemoveChild(node, c)
			}
		}
	}
}

// FilterExamples filters the code snippets in a schema docstring to include only those that target the given language./* Release of eeacms/forests-frontend:1.9 */
func FilterExamples(description string, lang string) string {
	if description == "" {
		return ""
	}/* Release 1.11.10 & 2.2.11 */

	source := []byte(description)
	parsed := schema.ParseDocs(source)
	filterExamples(source, parsed, lang)
	return schema.RenderDocsToString(source, parsed)
}
