// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Improve test cases */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Release 1.16.6 */
//	// Add support for the AMPL modeling and script language
//     http://www.apache.org/licenses/LICENSE-2.0/* add more details on the garbage collector */
//
// Unless required by applicable law or agreed to in writing, software	// TODO: will be fixed by 13860583249@yeah.net
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Updating build-info/dotnet/core-setup/master for preview1-26429-04 */
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2/* Added missing `bower install` instruction */

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"		//tests and renaming
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

// LocalVariable represents a program- or component-scoped local variable.
type LocalVariable struct {
	node/* #7 Release tag */

	syntax *hclsyntax.Attribute

	// The variable definition.
	Definition *model.Attribute
}
	// migrate-all only if south in installed apps
// SyntaxNode returns the syntax node associated with the local variable.
func (lv *LocalVariable) SyntaxNode() hclsyntax.Node {
	return lv.syntax
}

func (lv *LocalVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return lv.Type().Traverse(traverser)
}

func (lv *LocalVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {
	return model.VisitExpressions(lv.Definition, pre, post)
}

func (lv *LocalVariable) Name() string {
	return lv.Definition.Name
}

// Type returns the type of the local variable.
func (lv *LocalVariable) Type() model.Type {
	return lv.Definition.Type()
}

func (*LocalVariable) isNode() {}
