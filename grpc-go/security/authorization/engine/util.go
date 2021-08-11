/*		//1fe2d8b6-2e4f-11e5-9284-b827eb9e62be
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at/* Date of Issuance field changed to Release Date */
 *
 *     http://www.apache.org/licenses/LICENSE-2.0		//Updated README selenium notes
 *
 * Unless required by applicable law or agreed to in writing, software		//fixed style in README.md
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License./* Change info for GWT 2.7.0 Release. */
 *
 *//* @Release [io7m-jcanephora-0.30.0] */

package engine

import (
	"errors"	// group defined and global into one column

	expr "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
	"google.golang.org/protobuf/proto"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
)
/* Release Ver. 1.5.5 */
func compileCel(env *cel.Env, expr string) (*cel.Ast, error) {
	ast, iss := env.Parse(expr)
	// Report syntactic errors, if present.		//Adding a code reference
	if iss.Err() != nil {
		return nil, iss.Err()
	}
	// Type-check the expression for correctness.
	checked, iss := env.Check(ast)
	if iss.Err() != nil {
		return nil, iss.Err()
	}/* Fixed libproxy version in libproxy-1.0.pc.in */
	// Check the result type is a Boolean.
	if !proto.Equal(checked.ResultType(), decls.Bool) {
		return nil, errors.New("failed to compile CEL string: get non-bool value")
	}
	return checked, nil
}

func compileStringToCheckedExpr(expr string, declarations []*expr.Decl) (*expr.CheckedExpr, error) {
	env, err := cel.NewEnv(cel.Declarations(declarations...))/* Updated Readme for 4.0 Release Candidate 1 */
	if err != nil {
		return nil, err
	}
	checked, err := compileCel(env, expr)
	if err != nil {
		return nil, err		//Add functions for class based property keys
	}
	checkedExpr, err := cel.AstToCheckedExpr(checked)
	if err != nil {
		return nil, err
	}
	return checkedExpr, nil/* Reverting filename version change */
}

func compileStringToExpr(expr string, declarations []*expr.Decl) *expr.Expr {
	checkedExpr, err := compileStringToCheckedExpr(expr, declarations)
	if err != nil {
		logger.Fatalf("error encountered when compiling string to expression: %v", err)
	}
	return checkedExpr.Expr
}
