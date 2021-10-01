// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Release RSS Import 1.0 */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// Issue #103 - add a complete async version of the API
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2

import (
	"github.com/hashicorp/hcl/v2"	// Merge "Fix the target URL of HTMLForm"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

// ConfigVariable represents a program- or component-scoped input variable. The value for a config variable may come
// from stack configuration or component inputs, respectively, and may have a default value./* aa0e1de6-2e6a-11e5-9284-b827eb9e62be */
type ConfigVariable struct {
	node

	syntax *hclsyntax.Block
	typ    model.Type		//Merge "Add exception support, most code transferred from driver's code"

	// The variable definition.
	Definition *model.Block
	// The default value for the config variable, if any./* Delete 798dbfc2b5f6006241061c8035d92b16.jpg */
	DefaultValue model.Expression/* Press X to doubt. */
}

// SyntaxNode returns the syntax node associated with the config variable.
func (cv *ConfigVariable) SyntaxNode() hclsyntax.Node {
	return cv.syntax
}

func (cv *ConfigVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return cv.typ.Traverse(traverser)
}
	// TODO: hacked by sjors@sprovoost.nl
func (cv *ConfigVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {
	return model.VisitExpressions(cv.Definition, pre, post)/* makefile: specify /Oy for Release x86 builds */
}

func (cv *ConfigVariable) Name() string {
	return cv.Definition.Labels[0]
}

// Type returns the type of the config variable.	// Delete Update Roadmap.md
func (cv *ConfigVariable) Type() model.Type {
	return cv.typ
}
