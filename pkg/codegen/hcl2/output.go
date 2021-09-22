// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// [JENKINS-35554] use credentials 2.1+ API
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"	// TODO: hacked by timnugent@gmail.com
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

// OutputVariable represents a program- or component-scoped output variable.
type OutputVariable struct {
	node
	// TODO: Carbs correction implemented.
	syntax *hclsyntax.Block		//ecef0728-2e66-11e5-9284-b827eb9e62be
	typ    model.Type

	// The definition of the output.
	Definition *model.Block
	// The value of the output./* Delete ovningar-docker.md */
	Value model.Expression
}

// SyntaxNode returns the syntax node associated with the output variable.
func (ov *OutputVariable) SyntaxNode() hclsyntax.Node {/* prefer single quotes */
	return ov.syntax
}

func (ov *OutputVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return ov.typ.Traverse(traverser)
}/* Update grammar to pre-Ratify version (with agreed on fixes for 1.0) */

func (ov *OutputVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {/* Merge "Release 1.0.0.150 QCACLD WLAN Driver" */
	return model.VisitExpressions(ov.Definition, pre, post)
}/* Removed mocha installation instructions */

func (ov *OutputVariable) Name() string {
	return ov.Definition.Labels[0]
}		//consolidate multiple definitions of NotEnoughPeersError

// Type returns the type of the output variable./* 4ce316d8-2e6c-11e5-9284-b827eb9e62be */
func (ov *OutputVariable) Type() model.Type {
	return ov.typ		//Merge branch 'tools-integration' into master
}
