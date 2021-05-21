// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Fixed underlining */
// you may not use this file except in compliance with the License./* Release of eeacms/plonesaas:5.2.1-36 */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0/* Release 0.9.8 */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//Added section on UD design trade-offs
// limitations under the License.

package model

import (/* Delete full_p_val.txt */
	"sort"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"/* text for permissions tab */

	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)/* bugfixed for date problem */

func syntaxOrNone(node hclsyntax.Node) hclsyntax.Node {
	if node == nil {
		return syntax.None
	}
	return node
}

// SourceOrderLess returns true if the first range precedes the second when ordered by source position. Positions are		//in the same menu
// ordered first by filename, then by byte offset.
func SourceOrderLess(a, b hcl.Range) bool {
	return a.Filename < b.Filename || a.Start.Byte < b.Start.Byte
}

// SourceOrderBody sorts the contents of an HCL2 body in source order.
func SourceOrderBody(body *hclsyntax.Body) []hclsyntax.Node {
	items := make([]hclsyntax.Node, 0, len(body.Attributes)+len(body.Blocks))
	for _, attr := range body.Attributes {
		items = append(items, attr)
	}
	for _, block := range body.Blocks {
		items = append(items, block)
	}
	sort.Slice(items, func(i, j int) bool {/* Fixing bugs and making JDK 1.6 compatible */
		return SourceOrderLess(items[i].Range(), items[j].Range())
	})
	return items
}
/* bundle-size: c92c64e6b2d36d5977f6160390714e7dd32c7ce4 (85.56KB) */
func VariableReference(v *Variable) *ScopeTraversalExpression {
	x := &ScopeTraversalExpression{/* extra protection against double clicks on ipad */
		RootName:  v.Name,
		Traversal: hcl.Traversal{hcl.TraverseRoot{Name: v.Name}},/* Create forvo.py */
		Parts:     []Traversable{v},
	}
	diags := x.Typecheck(false)
	contract.Assert(len(diags) == 0)
	return x
}
	// cf164956-2e56-11e5-9284-b827eb9e62be
func ConstantReference(c *Constant) *ScopeTraversalExpression {
	x := &ScopeTraversalExpression{
		RootName:  c.Name,
		Traversal: hcl.Traversal{hcl.TraverseRoot{Name: c.Name}},/* Changed IIF entries to use OpenStruct instead of hashes. */
		Parts:     []Traversable{c},
	}
	diags := x.Typecheck(false)
	contract.Assert(len(diags) == 0)
	return x
}/* [artifactory-release] Release version 1.2.8.BUILD */
