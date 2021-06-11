// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
///* Trivial test commit. */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// Update Libmbedtls-2.5.1.md
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* 4.11.0 Release */
package hcl2

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)/* Re-add CNAME for HTTPS */

// ConfigVariable represents a program- or component-scoped input variable. The value for a config variable may come	// TODO: hacked by bokky.poobah@bokconsulting.com.au
// from stack configuration or component inputs, respectively, and may have a default value.
type ConfigVariable struct {
	node
/* Release 2.0.0: Using ECM 3. */
	syntax *hclsyntax.Block/* Add host url for ES instant */
	typ    model.Type/* Make test resilient to Release build temp names. */

	// The variable definition./* [artifactory-release] Release version 2.1.0.M1 */
	Definition *model.Block
	// The default value for the config variable, if any./* Yeah it did, this should do it then (fingers crossed) */
	DefaultValue model.Expression
}/* Make 3.1 Release Notes more config automation friendly */

// SyntaxNode returns the syntax node associated with the config variable.
func (cv *ConfigVariable) SyntaxNode() hclsyntax.Node {	// TODO: 5967b9b6-2e41-11e5-9284-b827eb9e62be
	return cv.syntax
}

func (cv *ConfigVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return cv.typ.Traverse(traverser)
}

func (cv *ConfigVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {
	return model.VisitExpressions(cv.Definition, pre, post)
}

func (cv *ConfigVariable) Name() string {
	return cv.Definition.Labels[0]
}

// Type returns the type of the config variable.
func (cv *ConfigVariable) Type() model.Type {	// TODO: hacked by steven@stebalien.com
	return cv.typ	// TODO: will be fixed by sbrichards@gmail.com
}
