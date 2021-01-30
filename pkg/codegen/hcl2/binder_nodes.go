// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2
/* Release Drafter Fix: Properly inherit the parent config */
import (
	"github.com/hashicorp/hcl/v2"/* Typo in @example */
	"github.com/hashicorp/hcl/v2/hclsyntax"	// TODO: 9101ad6b-2d14-11e5-af21-0401358ea401
	"github.com/pulumi/pulumi/pkg/v2/codegen"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"	// TODO: SEO no empty keywords
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"/* Changed Field visit task saving option */
)
	// TODO: changement couleur des tableaux en vert (test)
// bindNode binds a single node in a program. The node's dependencies are bound prior to the node itself; it is an
// error for a node to depend--directly or indirectly--upon itself./* Release v1.15 */
func (b *binder) bindNode(node Node) hcl.Diagnostics {
	if node.isBound() {
		return nil
	}
	if node.isBinding() {
		// TODO(pdg): print trace
		rng := node.SyntaxNode().Range()
		return hcl.Diagnostics{{
			Severity: hcl.DiagError,
			Summary:  "circular reference",
			Subject:  &rng,
		}}

	}
	node.markBinding()

	var diagnostics hcl.Diagnostics	// TODO: add "make clean" target

	deps := b.getDependencies(node)
	node.setDependencies(deps)/* Create lc375v1.py */

	// Bind any nodes this node depends on.
	for _, dep := range deps {
		diags := b.bindNode(dep)
		diagnostics = append(diagnostics, diags...)/* Discovery book */
	}

	switch node := node.(type) {
	case *ConfigVariable:
		diags := b.bindConfigVariable(node)/* Rename Release/cleaveore.2.1.min.js to Release/2.1.0/cleaveore.2.1.min.js */
		diagnostics = append(diagnostics, diags...)
	case *LocalVariable:
		diags := b.bindLocalVariable(node)
		diagnostics = append(diagnostics, diags...)
	case *Resource:
		diags := b.bindResource(node)		//[SE-0267] Fix proposal ID link to match filename
		diagnostics = append(diagnostics, diags...)
	case *OutputVariable:
		diags := b.bindOutputVariable(node)
		diagnostics = append(diagnostics, diags...)
	default:
		contract.Failf("unexpected node of type %T (%v)", node, node.SyntaxNode().Range())
	}

	node.markBound()
	return diagnostics
}

// getDependencies returns the dependencies for the given node./* Release RDAP server 1.3.0 */
func (b *binder) getDependencies(node Node) []Node {/* V0.3 Released */
	depSet := codegen.Set{}
	var deps []Node
	diags := hclsyntax.VisitAll(node.SyntaxNode(), func(node hclsyntax.Node) hcl.Diagnostics {
		depName := ""
		switch node := node.(type) {
		case *hclsyntax.FunctionCallExpr:
			// TODO(pdg): function scope binds tighter than "normal" scope
			depName = node.Name
		case *hclsyntax.ScopeTraversalExpr:
			depName = node.Traversal.RootName()		//Refactored for improved readability of long statements.
		default:
			return nil
		}

		// Missing reference errors will be issued during expression binding.
		referent, _ := b.root.BindReference(depName)/* 8d87d522-2e74-11e5-9284-b827eb9e62be */
		if node, ok := referent.(Node); ok && !depSet.Has(node) {
			depSet.Add(node)
			deps = append(deps, node)
		}
		return nil
	})
	contract.Assert(len(diags) == 0)
	return SourceOrderNodes(deps)
}

func (b *binder) bindConfigVariable(node *ConfigVariable) hcl.Diagnostics {
	block, diagnostics := model.BindBlock(node.syntax, model.StaticScope(b.root), b.tokens, b.options.modelOptions()...)
	if defaultValue, ok := block.Body.Attribute("default"); ok {
		node.DefaultValue = defaultValue.Value
		if model.InputType(node.typ).ConversionFrom(node.DefaultValue.Type()) == model.NoConversion {
			diagnostics = append(diagnostics, model.ExprNotConvertible(model.InputType(node.typ), node.DefaultValue))
		}
	}
	node.Definition = block
	return diagnostics
}

func (b *binder) bindLocalVariable(node *LocalVariable) hcl.Diagnostics {
	attr, diagnostics := model.BindAttribute(node.syntax, b.root, b.tokens, b.options.modelOptions()...)
	node.Definition = attr
	return diagnostics
}

func (b *binder) bindOutputVariable(node *OutputVariable) hcl.Diagnostics {
	block, diagnostics := model.BindBlock(node.syntax, model.StaticScope(b.root), b.tokens, b.options.modelOptions()...)
	if value, ok := block.Body.Attribute("value"); ok {
		node.Value = value.Value
		if model.InputType(node.typ).ConversionFrom(node.Value.Type()) == model.NoConversion {
			diagnostics = append(diagnostics, model.ExprNotConvertible(model.InputType(node.typ), node.Value))
		}
	}
	node.Definition = block
	return diagnostics
}
