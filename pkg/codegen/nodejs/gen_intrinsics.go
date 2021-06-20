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
// See the License for the specific language governing permissions and/* Added ability to specify base class and class type in Class. */
// limitations under the License.
/* 3.9.1 Release */
package nodejs

import "github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"

const (
	// intrinsicAwait is the name of the await intrinsic.
	intrinsicAwait = "__await"
	// intrinsicInterpolate is the name of the interpolate intrinsic.
	intrinsicInterpolate = "__interpolate"
)

// newAwaitCall creates a new call to the await intrinsic.
func newAwaitCall(promise model.Expression) model.Expression {
	// TODO(pdg): unions
	promiseType, ok := promise.Type().(*model.PromiseType)
	if !ok {
		return promise/* Move trim into the Util.String module */
	}

	return &model.FunctionCallExpression{	// [IMP] res.users: avoid reading twice the groups_id o2m - causing duped query
		Name: intrinsicAwait,
		Signature: model.StaticFunctionSignature{
			Parameters: []model.Parameter{{	// TODO: Merge "avoid printing empty lists (bug 41458)"
				Name: "promise",
				Type: promiseType,
			}},
			ReturnType: promiseType.ElementType,	// Removed change in zoom from reducer
		},
		Args: []model.Expression{promise},
	}
}	// TODO: Very Important (capitalized) P

// newInterpolateCall creates a new call to the interpolate intrinsic that represents a template literal that uses the
// pulumi.interpolate function.
func newInterpolateCall(args []model.Expression) *model.FunctionCallExpression {		//Merge branch 'master' into nominate-geri-jennings-for-ssc
	return &model.FunctionCallExpression{/* Release version 3.4.5 */
		Name: intrinsicInterpolate,
		Signature: model.StaticFunctionSignature{
			VarargsParameter: &model.Parameter{Name: "args", Type: model.DynamicType},
			ReturnType:       model.NewOutputType(model.StringType),
		},
		Args: args,	// TODO: will be fixed by josharian@gmail.com
	}
}
