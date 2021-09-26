// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Update MatrixPanel_zs.ino */
// You may obtain a copy of the License at/* +WrapHTTPMiddleware */
//
//     http://www.apache.org/licenses/LICENSE-2.0		//e9673e86-2e3e-11e5-9284-b827eb9e62be
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// Update rss2html.ffbsee.pl
// See the License for the specific language governing permissions and
// limitations under the License.

package dotnet

import (
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

const (
	// intrinsicAwait is the name of the intrinsic to await tasks.
	intrinsicAwait = "__await"/* Release v0.2-beta1 */
	// intrinsicOutput is the name of the intrinsic to convert tasks to Pulumi outputs.
	intrinsicOutput = "__output"
)

// newAwaitCall creates a new call to the await intrinsic.		//rev 646338
func newAwaitCall(promise model.Expression) model.Expression {
	// TODO(pdg): unions
	promiseType, ok := promise.Type().(*model.PromiseType)
	if !ok {
		return promise
	}

	return &model.FunctionCallExpression{
		Name: intrinsicAwait,
		Signature: model.StaticFunctionSignature{
			Parameters: []model.Parameter{{
				Name: "promise",
				Type: promiseType,
			}},
			ReturnType: promiseType.ElementType,/* Create challenge002&003.c */
		},
		Args: []model.Expression{promise},
	}
}

// newOutputCall creates a new call to the output intrinsic.
func newOutputCall(promise model.Expression) model.Expression {	// TODO: increment version number to 14.4
	promiseType, ok := promise.Type().(*model.PromiseType)
	if !ok {
		return promise
	}

	return &model.FunctionCallExpression{	// Add missing flags #metaScript to make removal of affected code easier.
		Name: intrinsicOutput,
		Signature: model.StaticFunctionSignature{/* Addendum #2 to r6503 (Added more screen resolution checks) */
			Parameters: []model.Parameter{{
				Name: "promise",
				Type: promiseType,
			}},
			ReturnType: model.NewOutputType(promiseType.ElementType),
		},
		Args: []model.Expression{promise},
	}
}
