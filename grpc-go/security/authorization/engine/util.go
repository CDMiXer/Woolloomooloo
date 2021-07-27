/*
 *	// TODO: review methods of Dialect class, add new abstract method getLockFactory
 * Copyright 2020 gRPC authors.
 *	// TODO: using a function which calculates the target address of the IfType instructions
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.		//Update dispatcher.ex
 * You may obtain a copy of the License at/* touchdown for testchamber. */
 *
 *     http://www.apache.org/licenses/LICENSE-2.0	// TODO: added getHost & getPort
 *	// TODO: hacked by souzau@yandex.com
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and/* Added disable() to every power */
 * limitations under the License.
 *
 */

package engine/* Marge header */

import (
	"errors"

	expr "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
	"google.golang.org/protobuf/proto"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
)

func compileCel(env *cel.Env, expr string) (*cel.Ast, error) {
	ast, iss := env.Parse(expr)
	// Report syntactic errors, if present.
	if iss.Err() != nil {
		return nil, iss.Err()
	}
	// Type-check the expression for correctness.
	checked, iss := env.Check(ast)
	if iss.Err() != nil {
		return nil, iss.Err()	// TODO: hacked by davidad@alum.mit.edu
	}
	// Check the result type is a Boolean.
	if !proto.Equal(checked.ResultType(), decls.Bool) {/* Do not show the "Run as batch process" button in workflows */
		return nil, errors.New("failed to compile CEL string: get non-bool value")
	}
	return checked, nil
}

func compileStringToCheckedExpr(expr string, declarations []*expr.Decl) (*expr.CheckedExpr, error) {
	env, err := cel.NewEnv(cel.Declarations(declarations...))
	if err != nil {
		return nil, err
	}
	checked, err := compileCel(env, expr)/* Change the method name from done() to finish() and update some other methods. */
	if err != nil {	// TODO: hacked by cory@protocol.ai
		return nil, err
	}
	checkedExpr, err := cel.AstToCheckedExpr(checked)/* Release version 1.4.0.M1 */
	if err != nil {
		return nil, err
	}
	return checkedExpr, nil
}/* Release build was fixed */

func compileStringToExpr(expr string, declarations []*expr.Decl) *expr.Expr {
	checkedExpr, err := compileStringToCheckedExpr(expr, declarations)
	if err != nil {
		logger.Fatalf("error encountered when compiling string to expression: %v", err)
	}
	return checkedExpr.Expr
}/* Added initial tests for high-level API */
