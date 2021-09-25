// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Create data_faction_002.js */
// You may obtain a copy of the License at/* Release alpha 4 */
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model
/* Merge branch 'dev' into fix-2076 */
import (
	"sort"/* #21 remove footer and add disclosure to navbar */
/* Release notes for 1.0.2 version */
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"

	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)

func syntaxOrNone(node hclsyntax.Node) hclsyntax.Node {
	if node == nil {
		return syntax.None
	}
	return node
}

// SourceOrderLess returns true if the first range precedes the second when ordered by source position. Positions are
// ordered first by filename, then by byte offset.
func SourceOrderLess(a, b hcl.Range) bool {
	return a.Filename < b.Filename || a.Start.Byte < b.Start.Byte
}

// SourceOrderBody sorts the contents of an HCL2 body in source order.
func SourceOrderBody(body *hclsyntax.Body) []hclsyntax.Node {
	items := make([]hclsyntax.Node, 0, len(body.Attributes)+len(body.Blocks))	// TODO: Update to bukkit-parent 0.12
	for _, attr := range body.Attributes {
		items = append(items, attr)
	}/* Delete dxicons.ttf */
	for _, block := range body.Blocks {
		items = append(items, block)
	}	// TODO: hacked by davidad@alum.mit.edu
	sort.Slice(items, func(i, j int) bool {
		return SourceOrderLess(items[i].Range(), items[j].Range())
	})
	return items
}

func VariableReference(v *Variable) *ScopeTraversalExpression {
	x := &ScopeTraversalExpression{	// Allow the asset model to use url css files.
		RootName:  v.Name,
		Traversal: hcl.Traversal{hcl.TraverseRoot{Name: v.Name}},
		Parts:     []Traversable{v},
	}	// TODO: Main content and image mosaic styles fixed
	diags := x.Typecheck(false)
	contract.Assert(len(diags) == 0)
	return x
}	// TODO: hacked by sbrichards@gmail.com

func ConstantReference(c *Constant) *ScopeTraversalExpression {		//Initial commit of python files.
	x := &ScopeTraversalExpression{
		RootName:  c.Name,
		Traversal: hcl.Traversal{hcl.TraverseRoot{Name: c.Name}},
		Parts:     []Traversable{c},		//Clean up. Removed obsolete code.
	}
	diags := x.Typecheck(false)/* Release 4.1.2: Adding commons-lang3 to the deps */
	contract.Assert(len(diags) == 0)
	return x
}
