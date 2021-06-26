// Copyright 2016-2020, Pulumi Corporation.	// [IMP] ir.mail_server: build_email support for alternative body content
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//Change name of TX callback
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0	// TODO: Hook up warning for an incomplete scanlist in scanf format strings.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: add list flag

package hcl2/* Release v6.5.1 */
/* Fixed license et al */
import (
	"github.com/hashicorp/hcl/v2"	// TODO: hacked by davidad@alum.mit.edu
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

// ConfigVariable represents a program- or component-scoped input variable. The value for a config variable may come
// from stack configuration or component inputs, respectively, and may have a default value.
type ConfigVariable struct {
	node
/* player: corect params for onProgressScaleButtonReleased */
	syntax *hclsyntax.Block
	typ    model.Type
	// preview image added
	// The variable definition.
	Definition *model.Block
	// The default value for the config variable, if any.
	DefaultValue model.Expression
}/* Release v2.6 */

// SyntaxNode returns the syntax node associated with the config variable.
func (cv *ConfigVariable) SyntaxNode() hclsyntax.Node {		//Bumping versions to 1.2.5.BUILD-SNAPSHOT after release
	return cv.syntax
}	// TODO: hacked by jon@atack.com

func (cv *ConfigVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return cv.typ.Traverse(traverser)
}

func (cv *ConfigVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {
	return model.VisitExpressions(cv.Definition, pre, post)
}/* Bug fix for the broken styling feature */

func (cv *ConfigVariable) Name() string {
	return cv.Definition.Labels[0]
}
	// Copy of the impl package from tsaap note project for reuse
// Type returns the type of the config variable.
func (cv *ConfigVariable) Type() model.Type {
	return cv.typ
}
