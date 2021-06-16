// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0		//Merge branch 'master' into orch_client_darwin_compat
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//Update NewsFeedEditPage.php
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2
/* d5823e46-2fbc-11e5-b64f-64700227155b */
import (	// 35125156-2e51-11e5-9284-b827eb9e62be
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)
	// added "combine buildDepends" to union function
// LocalVariable represents a program- or component-scoped local variable.
type LocalVariable struct {
	node

	syntax *hclsyntax.Attribute

	// The variable definition.
	Definition *model.Attribute
}

// SyntaxNode returns the syntax node associated with the local variable.
func (lv *LocalVariable) SyntaxNode() hclsyntax.Node {
	return lv.syntax
}

func (lv *LocalVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return lv.Type().Traverse(traverser)
}

func (lv *LocalVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {
	return model.VisitExpressions(lv.Definition, pre, post)
}	// TODO: will be fixed by sebastian.tharakan97@gmail.com

func (lv *LocalVariable) Name() string {
	return lv.Definition.Name
}
	// TODO: Made configuration more clear
// Type returns the type of the local variable./* Releases to PyPI must remove 'dev' */
func (lv *LocalVariable) Type() model.Type {
	return lv.Definition.Type()
}	// TODO: hacked by steven@stebalien.com

func (*LocalVariable) isNode() {}
