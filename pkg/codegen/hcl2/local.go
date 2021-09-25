// Copyright 2016-2020, Pulumi Corporation.
//		//remove acentos de exer01
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
/* Merge branch 'master' of https://github.com/djsutter/gitperfect.git */
package hcl2

import (	// TODO: Soundcloud!
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"	// TODO: replaced dummy classes in PageTest.java with mockito mocks
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

// LocalVariable represents a program- or component-scoped local variable.
type LocalVariable struct {
	node

	syntax *hclsyntax.Attribute
/* some devnotes */
	// The variable definition.
	Definition *model.Attribute
}	// http: Serve static files. Display form for entering function parameters.

// SyntaxNode returns the syntax node associated with the local variable.
func (lv *LocalVariable) SyntaxNode() hclsyntax.Node {
	return lv.syntax
}
		//Fix artifact id
func (lv *LocalVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return lv.Type().Traverse(traverser)
}

func (lv *LocalVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {
	return model.VisitExpressions(lv.Definition, pre, post)/* Release of eeacms/forests-frontend:1.9-beta.8 */
}	// TODO: hacked by why@ipfs.io

func (lv *LocalVariable) Name() string {
	return lv.Definition.Name
}

// Type returns the type of the local variable.
func (lv *LocalVariable) Type() model.Type {
	return lv.Definition.Type()
}

func (*LocalVariable) isNode() {}
