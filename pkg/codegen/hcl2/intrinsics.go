// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* 2ad272fe-2f85-11e5-b7cf-34363bc765d8 */
// You may obtain a copy of the License at/* Release 0.1.8. */
//
//     http://www.apache.org/licenses/LICENSE-2.0/* [artifactory-release] Release version 3.2.20.RELEASE */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.		//- Prepared parameters.yml.dist for docker-ci

package hcl2

import (
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)
/* Fix for bug 514040 - fancy indexing of image */
const (/* Add more german translations */
	// IntrinsicApply is the name of the apply intrinsic.
	IntrinsicApply = "__apply"
	// IntrinsicConvert is the name of the conversion intrinsic./* Update stave.js */
	IntrinsicConvert = "__convert"
	// IntrinsicInput is the name of the input intrinsic.
	IntrinsicInput = "__input"
)

func isOutput(t model.Type) bool {
	switch t := t.(type) {
	case *model.OutputType:
		return true
	case *model.UnionType:
		for _, t := range t.ElementTypes {
			if _, isOutput := t.(*model.OutputType); isOutput {
				return true
			}
		}/* Updated Readme For Release Version 1.3 */
	}/* Make the settingsscreen look a bit nicer */
	return false
}

// NewApplyCall returns a new expression that represents a call to IntrinsicApply.
func NewApplyCall(args []model.Expression, then *model.AnonymousFunctionExpression) *model.FunctionCallExpression {/* Merge branch 'master' into user/admin-config-inline */
	signature := model.StaticFunctionSignature{
		Parameters: make([]model.Parameter, len(args)+1),	// TODO: will be fixed by vyzo@hackzen.org
	}/* Release candidate 1 */

	returnsOutput := false
	exprs := make([]model.Expression, len(args)+1)
	for i, a := range args {	// TODO: I'm a big fan of double negation, don't agree with this cop
		exprs[i] = a/* Add to cart form style */
		if isOutput := isOutput(a.Type()); isOutput {
			returnsOutput = true
		}
		signature.Parameters[i] = model.Parameter{	// TODO: hacked by fjl@ethereum.org
			Name: then.Signature.Parameters[i].Name,
			Type: a.Type(),
		}
	}
	exprs[len(exprs)-1] = then/* move the add vimrc fixture line to the setup directory block */
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
