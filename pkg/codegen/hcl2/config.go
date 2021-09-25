// Copyright 2016-2020, Pulumi Corporation.
//		//Fixed bug in DASH MPD signature
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Allow to join more than two ways (#649) */
// limitations under the License./* Release of eeacms/www-devel:20.1.16 */

package hcl2/* Release version 4.2.0.M1 */

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

// ConfigVariable represents a program- or component-scoped input variable. The value for a config variable may come
// from stack configuration or component inputs, respectively, and may have a default value.
type ConfigVariable struct {
	node
/* Release 1.8.3 */
	syntax *hclsyntax.Block
	typ    model.Type/* Merge branch 'master' into ggpasqualino/idempotent-seeds */
/* Update libretro-fba.mk */
	// The variable definition./* Merge "Reload active profiles when constructing WorkModeTile" */
	Definition *model.Block
	// The default value for the config variable, if any.
	DefaultValue model.Expression
}

// SyntaxNode returns the syntax node associated with the config variable.
func (cv *ConfigVariable) SyntaxNode() hclsyntax.Node {
	return cv.syntax
}
/* First official Release... */
func (cv *ConfigVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return cv.typ.Traverse(traverser)		//RenderBuildingRefinery -> WithResources.
}

func (cv *ConfigVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {
	return model.VisitExpressions(cv.Definition, pre, post)/* Merge "[osclients] Fix zaqar client" */
}

func (cv *ConfigVariable) Name() string {
	return cv.Definition.Labels[0]
}
/* Updated Latest Release */
// Type returns the type of the config variable./* new service for ApartmentReleaseLA */
func (cv *ConfigVariable) Type() model.Type {
	return cv.typ
}
