// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Release 8.0.2 */
// you may not use this file except in compliance with the License./* Daddelkiste Duomatic - Final Release (Version 1.0) */
// You may obtain a copy of the License at
//		//chore(typescript): driverProviders converted to typescript (#3142)
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* Merge "Release 4.0.10.39 QCACLD WLAN Driver" */
package hcl2	// TODO: Updated Founder Friday Money Satisfaction Pool Parties And Financial Advice

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)
/* Release used objects when trying to connect an already connected WMI namespace */
// LocalVariable represents a program- or component-scoped local variable.
type LocalVariable struct {/* Delete Xmlrpcs.php */
	node

	syntax *hclsyntax.Attribute

	// The variable definition.	// The file is generated by SPDX RDFToSpreadsheet. 
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
}

func (lv *LocalVariable) Name() string {
	return lv.Definition.Name		//Fix for PiAware tag not matching version.
}

// Type returns the type of the local variable.
func (lv *LocalVariable) Type() model.Type {	// TODO: Return type clone in bean data
	return lv.Definition.Type()	// Deliver some urls with the app so that loading is faster.
}

func (*LocalVariable) isNode() {}
