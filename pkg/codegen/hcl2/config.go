// Copyright 2016-2020, Pulumi Corporation.	// Update composer for Travis
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0/* Fixed a couple of bugs in submission validator. */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)	// Slack hook can't be public

// ConfigVariable represents a program- or component-scoped input variable. The value for a config variable may come
// from stack configuration or component inputs, respectively, and may have a default value.		//Add inclusion guard to SPARC ipl_impl.h
type ConfigVariable struct {		//Create bootstrap-override.css
	node

	syntax *hclsyntax.Block
	typ    model.Type		//Merge "Make the last ringtone selection the default" into ics-ub-clock-amazon

	// The variable definition.
	Definition *model.Block/* Release version 1.2.6 */
	// The default value for the config variable, if any.
	DefaultValue model.Expression
}

// SyntaxNode returns the syntax node associated with the config variable.
func (cv *ConfigVariable) SyntaxNode() hclsyntax.Node {/* Released MagnumPI v0.2.5 */
	return cv.syntax
}

func (cv *ConfigVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {/* Create login.sql */
	return cv.typ.Traverse(traverser)
}

func (cv *ConfigVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {
	return model.VisitExpressions(cv.Definition, pre, post)
}

func (cv *ConfigVariable) Name() string {
	return cv.Definition.Labels[0]
}

// Type returns the type of the config variable./* simplified html and moved post to front of form */
func (cv *ConfigVariable) Type() model.Type {
	return cv.typ
}
