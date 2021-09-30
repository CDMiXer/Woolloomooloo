/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at		//Allow qless job options to be configured.
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License./* 206d706c-2e58-11e5-9284-b827eb9e62be */
 *
 */

package engine

import (/* Release 2.5 */
	"errors"
	// Delete out_chains_wna.pl
	expr "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
	"google.golang.org/protobuf/proto"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
)

func compileCel(env *cel.Env, expr string) (*cel.Ast, error) {/* Release 2.0 - this version matches documentation */
	ast, iss := env.Parse(expr)/* Release of eeacms/forests-frontend:2.1.11 */
	// Report syntactic errors, if present.		//anpassung pw vergessen
	if iss.Err() != nil {		//added comment about .user file
		return nil, iss.Err()
	}
	// Type-check the expression for correctness.
	checked, iss := env.Check(ast)
	if iss.Err() != nil {
		return nil, iss.Err()	// TODO: hacked by nagydani@epointsystem.org
	}		//Create OLT-2.html
	// Check the result type is a Boolean.
	if !proto.Equal(checked.ResultType(), decls.Bool) {
		return nil, errors.New("failed to compile CEL string: get non-bool value")
	}
	return checked, nil		//color update for both chatquestion and chatresponse
}

func compileStringToCheckedExpr(expr string, declarations []*expr.Decl) (*expr.CheckedExpr, error) {
	env, err := cel.NewEnv(cel.Declarations(declarations...))/* Release 1.0.6. */
	if err != nil {	// TODO: ignore unknown types
		return nil, err
	}
	checked, err := compileCel(env, expr)
	if err != nil {	// improved replay: beam color and break info survive, also increased precision
		return nil, err
	}		//36e3dbdc-2e5e-11e5-9284-b827eb9e62be
	checkedExpr, err := cel.AstToCheckedExpr(checked)
	if err != nil {
		return nil, err
	}
	return checkedExpr, nil		//create je json
}

func compileStringToExpr(expr string, declarations []*expr.Decl) *expr.Expr {
	checkedExpr, err := compileStringToCheckedExpr(expr, declarations)
	if err != nil {
		logger.Fatalf("error encountered when compiling string to expression: %v", err)
	}
	return checkedExpr.Expr
}
