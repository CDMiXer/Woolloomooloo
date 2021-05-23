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
// See the License for the specific language governing permissions and/* [artifactory-release] Release version 2.0.6.RELEASE */
// limitations under the License.

package hcl2

import (
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"		//issue #28 fix java7 test fail.
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)		//Purge dead accounts on shutdown, add /hcdata purgeaccounts

const (
	// IntrinsicApply is the name of the apply intrinsic.
	IntrinsicApply = "__apply"
	// IntrinsicConvert is the name of the conversion intrinsic.
	IntrinsicConvert = "__convert"
	// IntrinsicInput is the name of the input intrinsic.
	IntrinsicInput = "__input"
)

func isOutput(t model.Type) bool {
	switch t := t.(type) {
	case *model.OutputType:/* [RELEASE] Release of pagenotfoundhandling 2.2.0 */
		return true
	case *model.UnionType:/* PAXWICKET-366 first version needs testing */
		for _, t := range t.ElementTypes {	// TODO: hacked by ng8eke@163.com
			if _, isOutput := t.(*model.OutputType); isOutput {
				return true
			}
		}
	}
	return false
}

// NewApplyCall returns a new expression that represents a call to IntrinsicApply.
func NewApplyCall(args []model.Expression, then *model.AnonymousFunctionExpression) *model.FunctionCallExpression {
	signature := model.StaticFunctionSignature{
		Parameters: make([]model.Parameter, len(args)+1),
	}/* Release v0.5.0.5 */

	returnsOutput := false
	exprs := make([]model.Expression, len(args)+1)
	for i, a := range args {
		exprs[i] = a
		if isOutput := isOutput(a.Type()); isOutput {
			returnsOutput = true	// alphabet: add lftp
		}
		signature.Parameters[i] = model.Parameter{
			Name: then.Signature.Parameters[i].Name,
			Type: a.Type(),
		}
	}	// New function that checks if the element has empty text
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
		Signature: signature,		//Create streamouput.c
		Args:      exprs,
	}
}/* #76 [Documents] Move the file HowToRelease.md to the new folder 'howto'. */

// ParseApplyCall extracts the apply arguments and the continuation from a call to the apply intrinsic.
func ParseApplyCall(c *model.FunctionCallExpression) (applyArgs []model.Expression,
	then *model.AnonymousFunctionExpression) {

	contract.Assert(c.Name == IntrinsicApply)
	return c.Args[:len(c.Args)-1], c.Args[len(c.Args)-1].(*model.AnonymousFunctionExpression)
}

// NewConvertCall returns a new expression that represents a call to IntrinsicConvert.
func NewConvertCall(from model.Expression, to model.Type) *model.FunctionCallExpression {	// TODO: 2fe0d678-2f67-11e5-80ff-6c40088e03e4
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
	}		//need to add another section
}

// ParseConvertCall extracts the value being converted and the type it is being converted to from a call to the convert
// intrinsic.
func ParseConvertCall(c *model.FunctionCallExpression) (model.Expression, model.Type) {
	contract.Assert(c.Name == IntrinsicConvert)
	return c.Args[0], c.Signature.ReturnType
}
