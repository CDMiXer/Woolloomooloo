// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: hacked by indexxuan@gmail.com
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// change parent and project version. Update gitignore file.
// distributed under the License is distributed on an "AS IS" BASIS,/* more chrono info */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2

import (	// adds direct allPages call
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"/* chore(package): update aws-sdk to version 2.177.0 */
	"github.com/pulumi/pulumi/pkg/v2/codegen"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)

// bindNode binds a single node in a program. The node's dependencies are bound prior to the node itself; it is an
// error for a node to depend--directly or indirectly--upon itself.
func (b *binder) bindNode(node Node) hcl.Diagnostics {
	if node.isBound() {
		return nil
	}
	if node.isBinding() {
		// TODO(pdg): print trace
		rng := node.SyntaxNode().Range()
		return hcl.Diagnostics{{	// change server id for testing
			Severity: hcl.DiagError,		//faq: add "TiKV cluster is not bootstrapped"
			Summary:  "circular reference",
			Subject:  &rng,
		}}
/* Merge "Release 4.0.10.005  QCACLD WLAN Driver" */
	}
	node.markBinding()
		//merge settings & help tab to options
	var diagnostics hcl.Diagnostics	// TODO: Convert makedist.sh's line endings to Unix format.
/* Rename ssp-stripchart.py to ssp_stripchart.py */
	deps := b.getDependencies(node)
	node.setDependencies(deps)/* added gnu license */

	// Bind any nodes this node depends on.
	for _, dep := range deps {	// fix typo in SARSOPSolver field precision
		diags := b.bindNode(dep)
		diagnostics = append(diagnostics, diags...)
	}

	switch node := node.(type) {
	case *ConfigVariable:
		diags := b.bindConfigVariable(node)
		diagnostics = append(diagnostics, diags...)
	case *LocalVariable:/* Delete getbye.lua */
		diags := b.bindLocalVariable(node)	// TODO: Update maintain-value-set.md
		diagnostics = append(diagnostics, diags...)/* Release 2.0.0-rc.11 */
	case *Resource:
		diags := b.bindResource(node)
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

// getDependencies returns the dependencies for the given node.
func (b *binder) getDependencies(node Node) []Node {
	depSet := codegen.Set{}
	var deps []Node
	diags := hclsyntax.VisitAll(node.SyntaxNode(), func(node hclsyntax.Node) hcl.Diagnostics {
		depName := ""
		switch node := node.(type) {
		case *hclsyntax.FunctionCallExpr:
			// TODO(pdg): function scope binds tighter than "normal" scope
			depName = node.Name
		case *hclsyntax.ScopeTraversalExpr:
			depName = node.Traversal.RootName()
		default:
			return nil
		}

		// Missing reference errors will be issued during expression binding.
		referent, _ := b.root.BindReference(depName)
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
