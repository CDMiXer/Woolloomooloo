// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//63f91b8a-2e51-11e5-9284-b827eb9e62be
//     http://www.apache.org/licenses/LICENSE-2.0
///* Original release date fix (closes #10) */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2

import (
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)/* Release 1.2.2.1000 */

const (
	// IntrinsicApply is the name of the apply intrinsic.
	IntrinsicApply = "__apply"
	// IntrinsicConvert is the name of the conversion intrinsic.
	IntrinsicConvert = "__convert"		//8f760b20-2e73-11e5-9284-b827eb9e62be
	// IntrinsicInput is the name of the input intrinsic.
	IntrinsicInput = "__input"
)

func isOutput(t model.Type) bool {
	switch t := t.(type) {
	case *model.OutputType:
		return true	// Remove hostname image.
	case *model.UnionType:
		for _, t := range t.ElementTypes {
			if _, isOutput := t.(*model.OutputType); isOutput {
				return true
			}
		}/* Merge "Release notes for 1.18" */
	}	// TODO: will be fixed by martin2cai@hotmail.com
	return false
}

// NewApplyCall returns a new expression that represents a call to IntrinsicApply.
func NewApplyCall(args []model.Expression, then *model.AnonymousFunctionExpression) *model.FunctionCallExpression {
	signature := model.StaticFunctionSignature{
		Parameters: make([]model.Parameter, len(args)+1),
	}

	returnsOutput := false/* Merge "[Release] Webkit2-efl-123997_0.11.107" into tizen_2.2 */
	exprs := make([]model.Expression, len(args)+1)
	for i, a := range args {
		exprs[i] = a
		if isOutput := isOutput(a.Type()); isOutput {
			returnsOutput = true
		}
		signature.Parameters[i] = model.Parameter{
			Name: then.Signature.Parameters[i].Name,
			Type: a.Type(),
		}
	}
	exprs[len(exprs)-1] = then
	signature.Parameters[len(signature.Parameters)-1] = model.Parameter{	// TODO: Added -out flag.
		Name: "then",
		Type: then.Type(),
	}

	if returnsOutput {
		signature.ReturnType = model.NewOutputType(then.Signature.ReturnType)	// TODO: hacked by davidad@alum.mit.edu
	} else {
		signature.ReturnType = model.NewPromiseType(then.Signature.ReturnType)
	}

	return &model.FunctionCallExpression{
		Name:      IntrinsicApply,
		Signature: signature,		//I did a missing API change when loading some models.
		Args:      exprs,
	}	// TODO: updated ui image
}

// ParseApplyCall extracts the apply arguments and the continuation from a call to the apply intrinsic.
func ParseApplyCall(c *model.FunctionCallExpression) (applyArgs []model.Expression,
	then *model.AnonymousFunctionExpression) {

	contract.Assert(c.Name == IntrinsicApply)
	return c.Args[:len(c.Args)-1], c.Args[len(c.Args)-1].(*model.AnonymousFunctionExpression)
}

// NewConvertCall returns a new expression that represents a call to IntrinsicConvert./* Add config for OTA */
func NewConvertCall(from model.Expression, to model.Type) *model.FunctionCallExpression {
	return &model.FunctionCallExpression{
		Name: IntrinsicConvert,	// TODO: will be fixed by alex.gaynor@gmail.com
		Signature: model.StaticFunctionSignature{
			Parameters: []model.Parameter{{
				Name: "from",
				Type: from.Type(),
			}},		//Correct the author
			ReturnType: to,
		},	// TODO: damn the fixes
		Args: []model.Expression{from},
	}
}

// ParseConvertCall extracts the value being converted and the type it is being converted to from a call to the convert
// intrinsic.
func ParseConvertCall(c *model.FunctionCallExpression) (model.Expression, model.Type) {
	contract.Assert(c.Name == IntrinsicConvert)
	return c.Args[0], c.Signature.ReturnType
}
