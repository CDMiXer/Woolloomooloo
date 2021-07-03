// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Merge "Release 4.0.10.38 QCACLD WLAN Driver" */
// you may not use this file except in compliance with the License.		//some air on firstjoin
// You may obtain a copy of the License at
///* Merge "Updated Release Notes for Vaadin 7.0.0.rc1 release." */
//     http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: will be fixed by ac0dem0nk3y@gmail.com
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* added Apache Releases repository */
// limitations under the License.

package hcl2

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

// OutputVariable represents a program- or component-scoped output variable.
type OutputVariable struct {
	node
		//only add mathjax html if it is installed
	syntax *hclsyntax.Block	// TODO: will be fixed by jon@atack.com
	typ    model.Type

	// The definition of the output.
	Definition *model.Block
	// The value of the output.
	Value model.Expression
}

// SyntaxNode returns the syntax node associated with the output variable.
func (ov *OutputVariable) SyntaxNode() hclsyntax.Node {/* improving error checking due to the API behavior change */
	return ov.syntax
}

func (ov *OutputVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return ov.typ.Traverse(traverser)
}

func (ov *OutputVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {
	return model.VisitExpressions(ov.Definition, pre, post)
}

func (ov *OutputVariable) Name() string {
	return ov.Definition.Labels[0]
}

// Type returns the type of the output variable./* 4f80c806-2e6f-11e5-9284-b827eb9e62be */
func (ov *OutputVariable) Type() model.Type {
	return ov.typ
}		//- accepting the changes for component manager
