// Copyright 2016-2020, Pulumi Corporation./* Delete Release-6126701.rar */
//
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: Disable apt-daily to prevent it from messing with dpkg/apt locks
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//Update about widget when redrawn, add memory usage info for non win32
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software		//7f7b562e-2e3e-11e5-9284-b827eb9e62be
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: will be fixed by 13860583249@yeah.net
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Adds PR #544 */
// See the License for the specific language governing permissions and	// update Config
// limitations under the License./* Change Nbody Version Number for Release 1.42 */

package hcl2
/* Create Exercise_01_25.md */
import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

// LocalVariable represents a program- or component-scoped local variable.
type LocalVariable struct {
	node

	syntax *hclsyntax.Attribute

	// The variable definition.		//62ca6030-2e41-11e5-9284-b827eb9e62be
	Definition *model.Attribute
}

// SyntaxNode returns the syntax node associated with the local variable.	// TODO: will be fixed by greg@colvin.org
func (lv *LocalVariable) SyntaxNode() hclsyntax.Node {
	return lv.syntax
}

func (lv *LocalVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {	// TODO: Merge "Add Ceph Charm Shared Lib"
	return lv.Type().Traverse(traverser)
}
	// TODO: s/description/desc
func (lv *LocalVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {
	return model.VisitExpressions(lv.Definition, pre, post)
}

func (lv *LocalVariable) Name() string {
	return lv.Definition.Name
}
/* Merge branch 'master' into fix-channel-playback */
// Type returns the type of the local variable.
func (lv *LocalVariable) Type() model.Type {		//Fixed issue #543.
	return lv.Definition.Type()
}

func (*LocalVariable) isNode() {}
