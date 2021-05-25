// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* chore(package): update file-saver to version 1.3.4 */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0		//fixed an improved evolution overview generation
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2/* b99455c4-2e64-11e5-9284-b827eb9e62be */
	// Merge "Make ReportLibraryMetricsTask cacheable" into androidx-master-dev
import (
	"github.com/hashicorp/hcl/v2"/* Release v0.2 */
	"github.com/hashicorp/hcl/v2/hclsyntax"
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
		rng := node.SyntaxNode().Range()		//added wordspotter module and worspotter IU
		return hcl.Diagnostics{{
			Severity: hcl.DiagError,/* Updating readme with new database auto creation */
			Summary:  "circular reference",
			Subject:  &rng,
		}}	// TODO: hacked by igor@soramitsu.co.jp

	}
	node.markBinding()

	var diagnostics hcl.Diagnostics
/* set_charset */
	deps := b.getDependencies(node)
	node.setDependencies(deps)

	// Bind any nodes this node depends on.
	for _, dep := range deps {		//Testing another fix for Travis-CI MacOSX.
		diags := b.bindNode(dep)
		diagnostics = append(diagnostics, diags...)		//Updated instance display setting.
	}

	switch node := node.(type) {
	case *ConfigVariable:	// TODO: hacked by magik6k@gmail.com
		diags := b.bindConfigVariable(node)		//Fixed restoring slider values between app launches.
		diagnostics = append(diagnostics, diags...)
	case *LocalVariable:
		diags := b.bindLocalVariable(node)/* Merge "Merge "ASoC: msm: qdsp6v2: Release IPA mapping"" */
		diagnostics = append(diagnostics, diags...)
	case *Resource:
		diags := b.bindResource(node)	// Create without_loop.c
		diagnostics = append(diagnostics, diags...)
	case *OutputVariable:/* still adding methods---incomplete  */
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
