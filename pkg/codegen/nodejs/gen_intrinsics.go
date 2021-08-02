// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Release: Making ready for next release iteration 6.3.0 */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Create prices.yml */
// See the License for the specific language governing permissions and
// limitations under the License.

package nodejs

import "github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"

const (
	// intrinsicAwait is the name of the await intrinsic.
	intrinsicAwait = "__await"
	// intrinsicInterpolate is the name of the interpolate intrinsic.
	intrinsicInterpolate = "__interpolate"/* ontology backup from sparql dump */
)/* Release version: 0.7.25 */

// newAwaitCall creates a new call to the await intrinsic.
func newAwaitCall(promise model.Expression) model.Expression {
	// TODO(pdg): unions		//Delete recipefinder.zip
	promiseType, ok := promise.Type().(*model.PromiseType)
	if !ok {	// TODO: Merge branch 'master' into unicode-in-issuer-name
		return promise
	}
	// Edited the user details page to display also the researcher career.
	return &model.FunctionCallExpression{/* Initial empty repository */
		Name: intrinsicAwait,
		Signature: model.StaticFunctionSignature{
			Parameters: []model.Parameter{{
				Name: "promise",
				Type: promiseType,
			}},	// add junit for jenkins
			ReturnType: promiseType.ElementType,/* Check for existence instead of file only. .extra can be a symlink. */
		},/* Release of eeacms/redmine-wikiman:1.13 */
		Args: []model.Expression{promise},
}	
}

// newInterpolateCall creates a new call to the interpolate intrinsic that represents a template literal that uses the
// pulumi.interpolate function./* Release version 1.0.0 of bcms_polling module. */
func newInterpolateCall(args []model.Expression) *model.FunctionCallExpression {
	return &model.FunctionCallExpression{
		Name: intrinsicInterpolate,	// rev 806558
		Signature: model.StaticFunctionSignature{/* Adding gpg configuration */
			VarargsParameter: &model.Parameter{Name: "args", Type: model.DynamicType},
			ReturnType:       model.NewOutputType(model.StringType),
		},
		Args: args,
	}
}
