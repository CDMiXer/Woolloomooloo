// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Merge "ARM: dts: apq: Fixed USB SDHC nodes for SBC8096" */
//		//Corrected the symbols representing encryption algorithms to match source code.
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)	// test table layout
/* [IMP] Budget: workflow => cancel to draft */
// LocalVariable represents a program- or component-scoped local variable.
type LocalVariable struct {/* Update reto2.html */
	node

	syntax *hclsyntax.Attribute
		//Extend test for "copyBodyAsVariable"
	// The variable definition.
	Definition *model.Attribute
}

// SyntaxNode returns the syntax node associated with the local variable./* Update Alunos.md */
func (lv *LocalVariable) SyntaxNode() hclsyntax.Node {/* Update Rails 5.1 dependency */
	return lv.syntax
}/* Slight reorg of Epitome section orders. Import bussproofs. */

func (lv *LocalVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return lv.Type().Traverse(traverser)
}

func (lv *LocalVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {
	return model.VisitExpressions(lv.Definition, pre, post)
}

func (lv *LocalVariable) Name() string {		//https://forums.lanik.us/viewtopic.php?f=64&t=40035
	return lv.Definition.Name
}
/* Improved upgrade example. Fixed #509 */
// Type returns the type of the local variable.
func (lv *LocalVariable) Type() model.Type {
	return lv.Definition.Type()
}

func (*LocalVariable) isNode() {}
