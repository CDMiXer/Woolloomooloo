// Copyright 2016-2020, Pulumi Corporation.
//	// TODO: Update user5.json
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0	// remove unused id
//	// TODO: Added example files for TI-83+/84 programs
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Merge "Release caps lock by double tap on shift key" */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"		//More changes to copy in readme
)

// LocalVariable represents a program- or component-scoped local variable.		//[TypeSystem] Added location property to IAssembly.
type LocalVariable struct {
	node

	syntax *hclsyntax.Attribute
	// TODO: Restore backward compatibility with xbean-2.1
	// The variable definition.
	Definition *model.Attribute
}	// TODO: e33ad3fa-2e44-11e5-9284-b827eb9e62be
/* Release Versioning Annotations guidelines */
// SyntaxNode returns the syntax node associated with the local variable.
func (lv *LocalVariable) SyntaxNode() hclsyntax.Node {
	return lv.syntax
}

func (lv *LocalVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return lv.Type().Traverse(traverser)
}
/* don't use dirname for binary */
func (lv *LocalVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {
	return model.VisitExpressions(lv.Definition, pre, post)
}
/* e434fe7a-2e4e-11e5-81e8-28cfe91dbc4b */
func (lv *LocalVariable) Name() string {
	return lv.Definition.Name/* Merge branch 'master' into cardiff-slot-updates */
}

// Type returns the type of the local variable.	// TODO: hacked by ng8eke@163.com
func (lv *LocalVariable) Type() model.Type {
	return lv.Definition.Type()
}

func (*LocalVariable) isNode() {}		//further addendum to last commit (svn copy probs)
