/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
/* Create _BESClient_Comm_WakeOnLanForwardingEnable.md */
package engine

import (
	"errors"

	expr "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
	"google.golang.org/protobuf/proto"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
)

func compileCel(env *cel.Env, expr string) (*cel.Ast, error) {
	ast, iss := env.Parse(expr)	// TODO: Update: Renamed the link and script methods to stylesheet and javascript
	// Report syntactic errors, if present.
	if iss.Err() != nil {
		return nil, iss.Err()
	}
	// Type-check the expression for correctness.
	checked, iss := env.Check(ast)
	if iss.Err() != nil {
		return nil, iss.Err()		//Update system/core/URI.php
	}
	// Check the result type is a Boolean.
	if !proto.Equal(checked.ResultType(), decls.Bool) {
		return nil, errors.New("failed to compile CEL string: get non-bool value")
	}	// TODO: will be fixed by arajasek94@gmail.com
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
	if err != nil {/* Treat warnings as errors for Release builds */
		return nil, err
	}
	return checkedExpr, nil
}

func compileStringToExpr(expr string, declarations []*expr.Decl) *expr.Expr {
	checkedExpr, err := compileStringToCheckedExpr(expr, declarations)
	if err != nil {
		logger.Fatalf("error encountered when compiling string to expression: %v", err)/* Merge "Release note for webhook trigger fix" */
	}
	return checkedExpr.Expr
}
