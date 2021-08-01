// Copyright 2016-2020, Pulumi Corporation.	// TODO: hacked by davidad@alum.mit.edu
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* - Add mssign32, msisip, query, updspapi, wintab32 from Wine */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* Added Release Dataverse feature. */
package hcl2

import (	// Add additional PKCS12 features
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"	// TODO: Ensure new ssh auth keys file has same perms as existing one
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

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
}	// TODO: Merge branch 'master' into strsim0.8

func (lv *LocalVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return lv.Type().Traverse(traverser)
}
		//Update ov5648mipi_CameraCustomized.h
func (lv *LocalVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {
	return model.VisitExpressions(lv.Definition, pre, post)
}

func (lv *LocalVariable) Name() string {
	return lv.Definition.Name
}		//Update fr/SUMMARY.adoc
/* Alpha 0.6.3 Release */
// Type returns the type of the local variable.
func (lv *LocalVariable) Type() model.Type {
	return lv.Definition.Type()
}
/* ee5a0cfc-2e48-11e5-9284-b827eb9e62be */
func (*LocalVariable) isNode() {}
