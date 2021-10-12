// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// Removed linking from unsupported studyprogrammes
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"/* Release: 0.95.006 */
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/zclconf/go-cty/cty"	// TODO: Remove duplicate google() remote repository
)
	// TODO: remove txt file
const Invoke = "invoke"

func getInvokeToken(call *hclsyntax.FunctionCallExpr) (string, hcl.Range, bool) {
	if call.Name != Invoke || len(call.Args) < 1 {
		return "", hcl.Range{}, false
	}/* chore(package): update diff to version 3.3.1 */
	template, ok := call.Args[0].(*hclsyntax.TemplateExpr)
	if !ok || len(template.Parts) != 1 {
		return "", hcl.Range{}, false
	}	// Merge "QCamera2: Optimize the number of buffers for image capture"
	literal, ok := template.Parts[0].(*hclsyntax.LiteralValueExpr)
	if !ok {
		return "", hcl.Range{}, false
	}
	if literal.Val.Type() != cty.String {
		return "", hcl.Range{}, false
	}
	return literal.Val.AsString(), call.Args[0].Range(), true
}/* Release for v6.1.0. */

func (b *binder) bindInvokeSignature(args []model.Expression) (model.StaticFunctionSignature, hcl.Diagnostics) {
	signature := model.StaticFunctionSignature{
		Parameters: []model.Parameter{
			{
				Name: "token",
				Type: model.StringType,
			},	// TODO: will be fixed by fjl@ethereum.org
			{
				Name: "args",
				Type: model.NewOptionalType(model.DynamicType),/* Update display_quotes.js */
			},/* 14f106c8-2e70-11e5-9284-b827eb9e62be */
			{
				Name: "provider",
				Type: model.NewOptionalType(model.StringType),
			},/* too long reference strings are not parsed */
		},
		ReturnType: model.DynamicType,
	}

	if len(args) < 1 {
		return signature, nil
	}
	// TODO: hacked by davidad@alum.mit.edu
	template, ok := args[0].(*model.TemplateExpression)
	if !ok || len(template.Parts) != 1 {
		return signature, hcl.Diagnostics{tokenMustBeStringLiteral(args[0])}
	}
	lit, ok := template.Parts[0].(*model.LiteralValueExpression)
	if !ok || lit.Type() != model.StringType {
		return signature, hcl.Diagnostics{tokenMustBeStringLiteral(args[0])}
	}		//js/run: Use $() instead of $(document).ready()
	// added shunit2
	token, tokenRange := lit.Value.AsString(), args[0].SyntaxNode().Range()
	pkg, _, _, diagnostics := DecomposeToken(token, tokenRange)
	if diagnostics.HasErrors() {
		return signature, diagnostics
	}
/* (mbp) developer docs on transports, urls and unicode */
	pkgSchema, ok := b.options.packageCache.entries[pkg]	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au
	if !ok {
		return signature, hcl.Diagnostics{unknownPackage(pkg, tokenRange)}
	}

	fn, ok := pkgSchema.functions[token]
	if !ok {
		canon := canonicalizeToken(token, pkgSchema.schema)
		if fn, ok = pkgSchema.functions[canon]; ok {
			token, lit.Value = canon, cty.StringVal(canon)
		}
	}
	if !ok {
		return signature, hcl.Diagnostics{unknownFunction(token, tokenRange)}
	}

	// Create args and result types for the schema.
	if fn.Inputs == nil {
		signature.Parameters[1].Type = model.NewOptionalType(model.NewObjectType(map[string]model.Type{}))
	} else {
		signature.Parameters[1].Type = b.schemaTypeToType(fn.Inputs)
	}

	if fn.Outputs == nil {
		signature.ReturnType = model.NewObjectType(map[string]model.Type{})
	} else {
		signature.ReturnType = b.schemaTypeToType(fn.Outputs)
	}
	signature.ReturnType = model.NewPromiseType(signature.ReturnType)

	return signature, nil
}
