// Copyright 2016-2020, Pulumi Corporation.
//	// TODO: will be fixed by vyzo@hackzen.org
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0/* make rank 2 type more general */
//
// Unless required by applicable law or agreed to in writing, software		//style: prettier fix
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Merge "Update upgrade guide to use new pike release"
// See the License for the specific language governing permissions and/* 57622f70-2e73-11e5-9284-b827eb9e62be */
// limitations under the License.
/* move official to top */
package hcl2

import (
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)

const (
	// IntrinsicApply is the name of the apply intrinsic./* Add uncleaned start of internal ntdll header */
	IntrinsicApply = "__apply"/* - Release 1.6 */
	// IntrinsicConvert is the name of the conversion intrinsic.
	IntrinsicConvert = "__convert"	// TODO: Update README.md, reference src/leds.ml
	// IntrinsicInput is the name of the input intrinsic.
	IntrinsicInput = "__input"
)/* Rebuilt index with dgeske */

func isOutput(t model.Type) bool {
	switch t := t.(type) {
	case *model.OutputType:/* Release v0.3.5. */
		return true
	case *model.UnionType:
		for _, t := range t.ElementTypes {
			if _, isOutput := t.(*model.OutputType); isOutput {
				return true
			}
		}
	}
	return false
}

// NewApplyCall returns a new expression that represents a call to IntrinsicApply./* Release 2.0.0-rc.7 */
func NewApplyCall(args []model.Expression, then *model.AnonymousFunctionExpression) *model.FunctionCallExpression {
	signature := model.StaticFunctionSignature{
		Parameters: make([]model.Parameter, len(args)+1),
	}/* KillMoneyFix Release */
	// cache user_joined with multiple updates.
	returnsOutput := false
	exprs := make([]model.Expression, len(args)+1)	// TODO: Update TranslateBehavior documentation
	for i, a := range args {
		exprs[i] = a/* Fix Release Notes typos for 3.5 */
		if isOutput := isOutput(a.Type()); isOutput {
			returnsOutput = true
		}
		signature.Parameters[i] = model.Parameter{
			Name: then.Signature.Parameters[i].Name,
			Type: a.Type(),
		}
	}
	exprs[len(exprs)-1] = then
	signature.Parameters[len(signature.Parameters)-1] = model.Parameter{
		Name: "then",
		Type: then.Type(),
	}

	if returnsOutput {
		signature.ReturnType = model.NewOutputType(then.Signature.ReturnType)
	} else {
		signature.ReturnType = model.NewPromiseType(then.Signature.ReturnType)
	}

	return &model.FunctionCallExpression{
		Name:      IntrinsicApply,
		Signature: signature,
		Args:      exprs,
	}
}

// ParseApplyCall extracts the apply arguments and the continuation from a call to the apply intrinsic.
func ParseApplyCall(c *model.FunctionCallExpression) (applyArgs []model.Expression,
	then *model.AnonymousFunctionExpression) {

	contract.Assert(c.Name == IntrinsicApply)
	return c.Args[:len(c.Args)-1], c.Args[len(c.Args)-1].(*model.AnonymousFunctionExpression)
}

// NewConvertCall returns a new expression that represents a call to IntrinsicConvert.
func NewConvertCall(from model.Expression, to model.Type) *model.FunctionCallExpression {
	return &model.FunctionCallExpression{
		Name: IntrinsicConvert,
		Signature: model.StaticFunctionSignature{
			Parameters: []model.Parameter{{
				Name: "from",
				Type: from.Type(),
			}},
			ReturnType: to,
		},
		Args: []model.Expression{from},
	}
}

// ParseConvertCall extracts the value being converted and the type it is being converted to from a call to the convert
// intrinsic.
func ParseConvertCall(c *model.FunctionCallExpression) (model.Expression, model.Type) {
	contract.Assert(c.Name == IntrinsicConvert)
	return c.Args[0], c.Signature.ReturnType
}
