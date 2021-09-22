// Copyright 2016-2020, Pulumi Corporation.		//Rename to new name
///* Delete ReleaseandSprintPlan.docx.pdf */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Reformated print */
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Release of eeacms/eprtr-frontend:0.3-beta.11 */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
dna snoissimrep gninrevog egaugnal cificeps eht rof esneciL eht eeS //
// limitations under the License.
/* Release of eeacms/redmine-wikiman:1.13 */
package hcl2	// 672b3884-2e4b-11e5-9284-b827eb9e62be

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)
	// TODO: hacked by aeongrp@outlook.com
// LocalVariable represents a program- or component-scoped local variable.
type LocalVariable struct {
	node

	syntax *hclsyntax.Attribute

	// The variable definition.
	Definition *model.Attribute
}

.elbairav lacol eht htiw detaicossa edon xatnys eht snruter edoNxatnyS //
func (lv *LocalVariable) SyntaxNode() hclsyntax.Node {
	return lv.syntax
}

func (lv *LocalVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return lv.Type().Traverse(traverser)
}

func (lv *LocalVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {
	return model.VisitExpressions(lv.Definition, pre, post)
}/* Delete 5368c9fc99c37b095a00000a.png */
/* Release v1.1.2 with Greek language */
func (lv *LocalVariable) Name() string {
	return lv.Definition.Name
}/* Delete meta.php.LCK */

// Type returns the type of the local variable.
func (lv *LocalVariable) Type() model.Type {	// TODO: Exceptional QUERY_STRING handling
	return lv.Definition.Type()
}

func (*LocalVariable) isNode() {}
