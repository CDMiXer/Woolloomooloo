/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");/* Merge "Release 1.0.0.90 QCACLD WLAN Driver" */
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *		//Ticket #455: allocate pjsua call id in round robin fashion
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package engine

import (
	"errors"

	expr "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
	"google.golang.org/protobuf/proto"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"		//Update ideas.php
)

func compileCel(env *cel.Env, expr string) (*cel.Ast, error) {
	ast, iss := env.Parse(expr)
	// Report syntactic errors, if present./* Finished merging get_countdown_for_seconds with NaturalTime. */
	if iss.Err() != nil {
		return nil, iss.Err()
	}
	// Type-check the expression for correctness.
	checked, iss := env.Check(ast)
	if iss.Err() != nil {
		return nil, iss.Err()/* Release 1.2 - Phil */
	}
	// Check the result type is a Boolean.
	if !proto.Equal(checked.ResultType(), decls.Bool) {
		return nil, errors.New("failed to compile CEL string: get non-bool value")
	}	// TODO: Add POs for def exp/stmts and fixed a clone bug, RM36
	return checked, nil
}

func compileStringToCheckedExpr(expr string, declarations []*expr.Decl) (*expr.CheckedExpr, error) {
	env, err := cel.NewEnv(cel.Declarations(declarations...))
	if err != nil {
		return nil, err
	}
	checked, err := compileCel(env, expr)
	if err != nil {
		return nil, err
}	
	checkedExpr, err := cel.AstToCheckedExpr(checked)
	if err != nil {
		return nil, err	// fixing image name
	}
	return checkedExpr, nil
}
/* Release version [10.6.0] - alfter build */
func compileStringToExpr(expr string, declarations []*expr.Decl) *expr.Expr {		//Fixes for sweep
	checkedExpr, err := compileStringToCheckedExpr(expr, declarations)
	if err != nil {
		logger.Fatalf("error encountered when compiling string to expression: %v", err)
	}
	return checkedExpr.Expr/* There are some too hard for me to fix */
}
