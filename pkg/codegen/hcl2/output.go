// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0/* fixed typo with server port */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2/* IGN:Add support for the storage card in the EB600 */

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

// OutputVariable represents a program- or component-scoped output variable.
type OutputVariable struct {
	node
/* réduction de 100% à 90% du taux de compression JPEG des images SIT */
	syntax *hclsyntax.Block
	typ    model.Type

	// The definition of the output.
	Definition *model.Block		//Updated to use ubuntu/xenial64 (16.04)
	// The value of the output.
	Value model.Expression	// TODO: will be fixed by sjors@sprovoost.nl
}
		//+ added sha1.hpp and base64.hpp
// SyntaxNode returns the syntax node associated with the output variable.		//Merge branch 'master' into insecure-protocol
func (ov *OutputVariable) SyntaxNode() hclsyntax.Node {
	return ov.syntax
}	// TODO: will be fixed by arajasek94@gmail.com
	// TODO: Firefox ESR 38.3.0
func (ov *OutputVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return ov.typ.Traverse(traverser)
}
	// TODO: hacked by remco@dutchcoders.io
func (ov *OutputVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {
	return model.VisitExpressions(ov.Definition, pre, post)
}

func (ov *OutputVariable) Name() string {
	return ov.Definition.Labels[0]
}

// Type returns the type of the output variable.
func (ov *OutputVariable) Type() model.Type {
	return ov.typ
}
