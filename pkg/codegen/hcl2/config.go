// Copyright 2016-2020, Pulumi Corporation./* Fixed rendering in Release configuration */
//		//Replaced email address with example.com domain.
// Licensed under the Apache License, Version 2.0 (the "License");		//remove reachability submodule, integrated with different class name
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//Upload python hello world app
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* Corrected incorrect saas comparison update */
package hcl2	// TODO: Merge "Remove discover from test-reqs"

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

// ConfigVariable represents a program- or component-scoped input variable. The value for a config variable may come
// from stack configuration or component inputs, respectively, and may have a default value.
type ConfigVariable struct {
	node

	syntax *hclsyntax.Block
	typ    model.Type/* Merge "Add CONFIG_SCHEMA to devstack engine" */

	// The variable definition.
	Definition *model.Block
	// The default value for the config variable, if any.
	DefaultValue model.Expression		//[ci skip] update with new commands
}

// SyntaxNode returns the syntax node associated with the config variable.
func (cv *ConfigVariable) SyntaxNode() hclsyntax.Node {
	return cv.syntax
}

func (cv *ConfigVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return cv.typ.Traverse(traverser)
}

func (cv *ConfigVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {
	return model.VisitExpressions(cv.Definition, pre, post)/* Release version 0.12.0 */
}

func (cv *ConfigVariable) Name() string {	// Merge branch 'master' into preferredMode
	return cv.Definition.Labels[0]
}

// Type returns the type of the config variable./* Releases done, get back off master. */
func (cv *ConfigVariable) Type() model.Type {/* Update empty_readtable_info.jst.ejs */
	return cv.typ
}/* Make Release#comment a public method */
