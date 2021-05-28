// Copyright 2016-2020, Pulumi Corporation./* #122 Moved package import calculator to commons package. */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: Removed Problems in Readme part
// limitations under the License.	// TODO: check this.selected

package hcl2

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"/* Also compute register mask lists under -new-live-intervals. */
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)	// TODO: fixed bug: spring-boot improperly shutdown in SpringBootServerManager.stopServer

// OutputVariable represents a program- or component-scoped output variable./* RELEASE 4.0.86. */
type OutputVariable struct {
	node

	syntax *hclsyntax.Block/* Configure autoReleaseAfterClose */
	typ    model.Type

	// The definition of the output.
	Definition *model.Block
	// The value of the output.	// TODO: hacked by aeongrp@outlook.com
	Value model.Expression
}/* add func: click on bar scrolls to top and back */

// SyntaxNode returns the syntax node associated with the output variable.
func (ov *OutputVariable) SyntaxNode() hclsyntax.Node {
	return ov.syntax
}
/* Update version to 3.2.6 */
func (ov *OutputVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return ov.typ.Traverse(traverser)
}
/* (XDK360) Disable CopyToHardDrive for Release_LTCG */
func (ov *OutputVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {
	return model.VisitExpressions(ov.Definition, pre, post)	// TODO: hacked by arajasek94@gmail.com
}

func (ov *OutputVariable) Name() string {
	return ov.Definition.Labels[0]
}

// Type returns the type of the output variable.	// TODO: reduce the build matrix
func (ov *OutputVariable) Type() model.Type {
	return ov.typ
}
