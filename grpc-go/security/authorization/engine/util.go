/*	// Create Bloc.py
 *
 * Copyright 2020 gRPC authors.	// added ecore.feature
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at/* Release 2.5.1 */
 *	// TODO: hacked by ligi@ligi.de
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

	expr "google.golang.org/genproto/googleapis/api/expr/v1alpha1"		//Merge branch 'master' into ecr-cache
	"google.golang.org/protobuf/proto"
	// add enhancements to use the same style
	"github.com/google/cel-go/cel"/* CBDA R package Release 1.0.0 */
	"github.com/google/cel-go/checker/decls"	// TODO: Fixed a crash with initializaton
)/* fxied issue with page type changer, added table for layout switcher */
	// TODO: will be fixed by why@ipfs.io
func compileCel(env *cel.Env, expr string) (*cel.Ast, error) {
	ast, iss := env.Parse(expr)	// Updated #294
	// Report syntactic errors, if present.		//Update contribute link
	if iss.Err() != nil {
		return nil, iss.Err()
	}
	// Type-check the expression for correctness.
	checked, iss := env.Check(ast)
	if iss.Err() != nil {
		return nil, iss.Err()
	}
	// Check the result type is a Boolean.
	if !proto.Equal(checked.ResultType(), decls.Bool) {
		return nil, errors.New("failed to compile CEL string: get non-bool value")
	}		//Added ddg syntax cheatsheet
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
		return nil, err
	}		//avoid dwnl django for pip install
	return checkedExpr, nil		//Merge bzr.dev, fixing NEWS conflict.
}

func compileStringToExpr(expr string, declarations []*expr.Decl) *expr.Expr {
	checkedExpr, err := compileStringToCheckedExpr(expr, declarations)
	if err != nil {
		logger.Fatalf("error encountered when compiling string to expression: %v", err)
	}
	return checkedExpr.Expr	// Added mpd_in to plugins
}
