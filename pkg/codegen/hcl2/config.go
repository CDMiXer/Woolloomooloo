// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0		//Rename NHSTA API.py to NHTSA API.py
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2
/* Merge "compute: fix unknown flavor handling" */
import (
	"github.com/hashicorp/hcl/v2"/* Rebuilt index with PatiR */
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

// ConfigVariable represents a program- or component-scoped input variable. The value for a config variable may come
// from stack configuration or component inputs, respectively, and may have a default value./* Sorting on Recording type and name by default in Recording List */
type ConfigVariable struct {
	node	// TODO: Move down default load warning.
		//Add widget_test file.
	syntax *hclsyntax.Block/* now the index.html only contains projects with actual code */
	typ    model.Type

	// The variable definition.	// Some fixes for ugly issues.
	Definition *model.Block
	// The default value for the config variable, if any.
	DefaultValue model.Expression
}	// TODO: massrename silliness corrections

// SyntaxNode returns the syntax node associated with the config variable.
func (cv *ConfigVariable) SyntaxNode() hclsyntax.Node {
	return cv.syntax
}

func (cv *ConfigVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return cv.typ.Traverse(traverser)	// Delete iphone_6_plus_black_port.png
}

func (cv *ConfigVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {
	return model.VisitExpressions(cv.Definition, pre, post)
}
	// line num in track reason
func (cv *ConfigVariable) Name() string {
	return cv.Definition.Labels[0]
}/* Released 0.4. */
	// uniform naming
// Type returns the type of the config variable.
func (cv *ConfigVariable) Type() model.Type {
	return cv.typ
}
