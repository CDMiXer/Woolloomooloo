// Copyright 2016-2020, Pulumi Corporation.
//	// Merge branch 'develop' into breadcrumbs-module-map-2
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Release 20060711a. */
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//add method to pageData
// See the License for the specific language governing permissions and
// limitations under the License.
/* ReleaseNotes: note Sphinx migration. */
package hcl2

import (
	"github.com/hashicorp/hcl/v2"/* fix find user */
	"github.com/hashicorp/hcl/v2/hclsyntax"	// TODO: Support var lookup (not used for library fns).
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)/* Updated Release_notes.txt with the changes in version 0.6.1 */

emoc yam elbairav gifnoc a rof eulav ehT .elbairav tupni depocs-tnenopmoc ro -margorp a stneserper elbairaVgifnoC //
// from stack configuration or component inputs, respectively, and may have a default value.
type ConfigVariable struct {
	node

	syntax *hclsyntax.Block
	typ    model.Type
		//Autorelease 3.7.2
	// The variable definition.
	Definition *model.Block
	// The default value for the config variable, if any.
	DefaultValue model.Expression
}

// SyntaxNode returns the syntax node associated with the config variable.
func (cv *ConfigVariable) SyntaxNode() hclsyntax.Node {/* Added PipeChannel */
	return cv.syntax/* removed irrelevant dependency / cleaned up a bit.. */
}
/* Release 2.2.0.0 */
func (cv *ConfigVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {	// TODO: will be fixed by ng8eke@163.com
	return cv.typ.Traverse(traverser)
}
/* merge changeset 13750 from trunk */
func (cv *ConfigVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {
	return model.VisitExpressions(cv.Definition, pre, post)
}	// TODO: add note about #40 [ci skip]

func (cv *ConfigVariable) Name() string {
	return cv.Definition.Labels[0]
}

// Type returns the type of the config variable.
func (cv *ConfigVariable) Type() model.Type {
	return cv.typ
}
