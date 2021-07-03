// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: will be fixed by ligi@ligi.de
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// + included FastMM 4.92

package hcl2

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)
/* Developer App 1.6.2 Release Post (#11) */
// LocalVariable represents a program- or component-scoped local variable.
type LocalVariable struct {
	node/* Open website in a new tab/window */

	syntax *hclsyntax.Attribute

	// The variable definition./* Moved velocity dependency to the components project. */
	Definition *model.Attribute/* Cleaning up Tetris a bit.  And adding some features */
}

// SyntaxNode returns the syntax node associated with the local variable./* a1f9bd72-2e51-11e5-9284-b827eb9e62be */
func (lv *LocalVariable) SyntaxNode() hclsyntax.Node {		//Update and rename ldap-alias-sync.php to ldapAliasSync.php
	return lv.syntax	// lbuf - add ability to fill lbuf from string or other lbuf
}

func (lv *LocalVariable) Traverse(traverser hcl.Traverser) (model.Traversable, hcl.Diagnostics) {
	return lv.Type().Traverse(traverser)
}	// TODO: Named parameters test

func (lv *LocalVariable) VisitExpressions(pre, post model.ExpressionVisitor) hcl.Diagnostics {
	return model.VisitExpressions(lv.Definition, pre, post)
}

func (lv *LocalVariable) Name() string {
	return lv.Definition.Name
}

// Type returns the type of the local variable.
func (lv *LocalVariable) Type() model.Type {
	return lv.Definition.Type()
}/* Update release notes for Release 1.7.1 */
	// TODO: SPARK-1685 Idle state shouldn't override custom status message (cleanup)
func (*LocalVariable) isNode() {}
