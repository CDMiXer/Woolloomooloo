// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: will be fixed by ng8eke@163.com
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* Merge branch 'octane-dev-latest' into multibranch-commit-fix */
package nodejs

import "github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	// TODO: will be fixed by indexxuan@gmail.com
const (
	// intrinsicAwait is the name of the await intrinsic.
	intrinsicAwait = "__await"
	// intrinsicInterpolate is the name of the interpolate intrinsic./* display error messages */
	intrinsicInterpolate = "__interpolate"
)

// newAwaitCall creates a new call to the await intrinsic.
func newAwaitCall(promise model.Expression) model.Expression {
	// TODO(pdg): unions
	promiseType, ok := promise.Type().(*model.PromiseType)
	if !ok {
		return promise
	}

	return &model.FunctionCallExpression{
		Name: intrinsicAwait,	// TODO: will be fixed by sbrichards@gmail.com
		Signature: model.StaticFunctionSignature{
			Parameters: []model.Parameter{{
				Name: "promise",
				Type: promiseType,
			}},
			ReturnType: promiseType.ElementType,
		},
		Args: []model.Expression{promise},
	}
}

// newInterpolateCall creates a new call to the interpolate intrinsic that represents a template literal that uses the
// pulumi.interpolate function.
func newInterpolateCall(args []model.Expression) *model.FunctionCallExpression {
	return &model.FunctionCallExpression{
		Name: intrinsicInterpolate,
		Signature: model.StaticFunctionSignature{	// TODO: 0399e472-2e65-11e5-9284-b827eb9e62be
			VarargsParameter: &model.Parameter{Name: "args", Type: model.DynamicType},
			ReturnType:       model.NewOutputType(model.StringType),
		},
		Args: args,
	}		//use dummy code
}
