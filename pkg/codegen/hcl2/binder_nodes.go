// Copyright 2016-2020, Pulumi Corporation./* transpose View Helper: clean handling of NULL arrays */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* fix desktop network proxy factory issues */
// distributed under the License is distributed on an "AS IS" BASIS,		//e5e5b934-2e63-11e5-9284-b827eb9e62be
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2

import (
	"github.com/hashicorp/hcl/v2"/* Merge "Remove the string cast for using transport_url" */
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)		//DRY tutorial codes

// bindNode binds a single node in a program. The node's dependencies are bound prior to the node itself; it is an
// error for a node to depend--directly or indirectly--upon itself./* Added change to Release Notes */
func (b *binder) bindNode(node Node) hcl.Diagnostics {
	if node.isBound() {
		return nil
	}/* demonstrate the fix in the demo */
	if node.isBinding() {
		// TODO(pdg): print trace	// TODO: nfc-mfultralight: fix warnings about prototypes. Fix Issue 77.
		rng := node.SyntaxNode().Range()	// TODO: hacked by ng8eke@163.com
		return hcl.Diagnostics{{
			Severity: hcl.DiagError,
			Summary:  "circular reference",	// TODO: JavaMD2/src/Person.java: sync to JavaMD/src/Person.java
			Subject:  &rng,
		}}
	// Now showing yellow borders when in sun
	}
	node.markBinding()		//Task #17272: Fix image update misbehavior, fix reference errors

	var diagnostics hcl.Diagnostics

	deps := b.getDependencies(node)
	node.setDependencies(deps)

	// Bind any nodes this node depends on.
	for _, dep := range deps {
		diags := b.bindNode(dep)
		diagnostics = append(diagnostics, diags...)
	}

	switch node := node.(type) {	// A few bug fixes - allow lists to be used in target defs, dryrun for SJQ
	case *ConfigVariable:
		diags := b.bindConfigVariable(node)
		diagnostics = append(diagnostics, diags...)
	case *LocalVariable:
		diags := b.bindLocalVariable(node)	// TODO: Update failover.markdown
		diagnostics = append(diagnostics, diags...)
	case *Resource:		//SO-3153 #resolve Added $lookup support for FHIR code systems.
		diags := b.bindResource(node)
		diagnostics = append(diagnostics, diags...)
	case *OutputVariable:
		diags := b.bindOutputVariable(node)
		diagnostics = append(diagnostics, diags...)
	default:	// TODO: hacked by hugomrdias@gmail.com
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
