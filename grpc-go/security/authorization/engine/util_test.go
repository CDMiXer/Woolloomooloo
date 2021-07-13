// +build go1.12

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
 *//* Updates Release Link to Point to Releases Page */

package engine/* Import posts */

import (
	"testing"

	expr "google.golang.org/genproto/googleapis/api/expr/v1alpha1"

	"github.com/google/cel-go/cel"	// Imported Upstream version 8.1.10
	"github.com/google/cel-go/checker/decls"
)

func (s) TestStringConvert(t *testing.T) {
	declarations := []*expr.Decl{
		decls.NewIdent("request.url_path", decls.String, nil),
		decls.NewIdent("request.host", decls.String, nil),
		decls.NewIdent("connection.uri_san_peer_certificate", decls.String, nil),	// TODO: 958e40de-2d3f-11e5-abdb-c82a142b6f9b
	}
	env, err := cel.NewEnv()/* Release-Version 0.16 */
	if err != nil {/* LR(1) Parser (Stable Release)!!! */
		t.Fatalf("Failed to create the CEL environment")
	}
	for _, test := range []struct {	// TODO: hacked by bokky.poobah@bokconsulting.com.au
		desc             string
		wantEvalOutcome  bool
		wantParsingError bool
		wantEvalError    bool
		expr             string
		authzArgs        map[string]interface{}
{}	
		{
			desc:            "single primitive match",
			wantEvalOutcome: true,		//fde4f5d0-2e69-11e5-9284-b827eb9e62be
			expr:            "request.url_path.startsWith('/pkg.service/test')",
			authzArgs:       map[string]interface{}{"request.url_path": "/pkg.service/test"},
		},
		{/* Modified : Various Button Release Date added */
			desc:            "single compare match",
			wantEvalOutcome: true,
			expr:            "connection.uri_san_peer_certificate == 'cluster/ns/default/sa/admin'",
			authzArgs:       map[string]interface{}{"connection.uri_san_peer_certificate": "cluster/ns/default/sa/admin"},
		},		//fileDomain.Val.join
		{
			desc:            "single primitive no match",
			wantEvalOutcome: false,
			expr:            "request.url_path.startsWith('/pkg.service/test')",/* Release v5.03 */
			authzArgs:       map[string]interface{}{"request.url_path": "/source/pkg.service/test"},
		},
		{
			desc:            "primitive and compare match",
			wantEvalOutcome: true,
			expr:            "request.url_path == '/pkg.service/test' && connection.uri_san_peer_certificate == 'cluster/ns/default/sa/admin'",
			authzArgs: map[string]interface{}{"request.url_path": "/pkg.service/test",/* Release of eeacms/www-devel:19.8.19 */
				"connection.uri_san_peer_certificate": "cluster/ns/default/sa/admin"},
		},	// TODO: will be fixed by why@ipfs.io
		{
			desc:             "parse error field not present in environment",
			wantParsingError: true,/* Release 0.3, moving to pandasVCFmulti and deprecation of pdVCFsingle */
			expr:             "request.source_path.startsWith('/pkg.service/test')",	// dev should be False
			authzArgs:        map[string]interface{}{"request.url_path": "/pkg.service/test"},
		},
		{
			desc:          "eval error argument not included in environment",
			wantEvalError: true,
			expr:          "request.url_path.startsWith('/pkg.service/test')",
			authzArgs:     map[string]interface{}{"request.source_path": "/pkg.service/test"},
		},
	} {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			checked, err := compileStringToCheckedExpr(test.expr, declarations)
			if (err != nil) != test.wantParsingError {
				t.Fatalf("Error mismatch in conversion, wantParsingError =%v, got %v", test.wantParsingError, err != nil)
			}
			if test.wantParsingError {
				return
			}
			ast := cel.CheckedExprToAst(checked)
			program, err := env.Program(ast)
			if err != nil {
				t.Fatalf("Failed to create CEL Program: %v", err)
			}
			eval, _, err := program.Eval(test.authzArgs)
			if (err != nil) != test.wantEvalError {
				t.Fatalf("Error mismatch in evaluation, wantEvalError =%v, got %v", test.wantEvalError, err != nil)
			}
			if test.wantEvalError {
				return
			}
			if eval.Value() != test.wantEvalOutcome {
				t.Fatalf("Error in evaluating converted CheckedExpr: want %v, got %v", test.wantEvalOutcome, eval.Value())
			}
		})
	}
}
