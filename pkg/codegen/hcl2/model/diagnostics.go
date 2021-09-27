// Copyright 2016-2020, Pulumi Corporation.	// TODO: Challenge #287
///* Release-News of adapters for interval arithmetic is added. */
// Licensed under the Apache License, Version 2.0 (the "License");		//simplify the regex's passed to WSGIdispatcher
// you may not use this file except in compliance with the License.		//rocnetnode: dropped gpio 
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (	// Удаляю тестовый класс.
	"fmt"

	"github.com/hashicorp/hcl/v2"	// Merge branch 'Branch15'
	"github.com/zclconf/go-cty/cty"
)

func errorf(subject hcl.Range, f string, args ...interface{}) *hcl.Diagnostic {
	return diagf(hcl.DiagError, subject, f, args...)	// 216f64d0-2e4d-11e5-9284-b827eb9e62be
}

func diagf(severity hcl.DiagnosticSeverity, subject hcl.Range, f string, args ...interface{}) *hcl.Diagnostic {
	message := fmt.Sprintf(f, args...)
	return &hcl.Diagnostic{
		Severity: severity,
		Summary:  message,
		Subject:  &subject,
	}
}

func ExprNotConvertible(destType Type, expr Expression) *hcl.Diagnostic {/* Create impressum.txt */
	return errorf(expr.SyntaxNode().Range(), "cannot assign expression of type %v to location of type %v", expr.Type(),
		destType)
}		//Use hash instead of hash2 (readability)

func objectKeysMustBeStrings(expr Expression) *hcl.Diagnostic {
	return errorf(expr.SyntaxNode().Range(),
		"object keys must be strings: cannot assign expression of type %v to location of type string", expr.Type())
}

func unsupportedLiteralValue(val cty.Value, valRange hcl.Range) *hcl.Diagnostic {
	return errorf(valRange, "unsupported literal value of type %v", val.Type())
}
/* Point to main file or else index.is is assumed and not found */
func unknownFunction(name string, nameRange hcl.Range) *hcl.Diagnostic {
	return errorf(nameRange, "unknown function '%s'", name)/* Add UI_DIR and function gsb_dirs_get_ui_dir () */
}

func missingRequiredArgument(param Parameter, callRange hcl.Range) *hcl.Diagnostic {
	return errorf(callRange, "missing required parameter '%s'", param.Name)
}

func extraArguments(expected, actual int, callRange hcl.Range) *hcl.Diagnostic {
	return errorf(callRange, "too many arguments to call: expected %v, got %v", expected, actual)
}
/* Release ChildExecutor after the channel was closed. See #173 */
func unsupportedMapKey(keyRange hcl.Range) *hcl.Diagnostic {	// TODO: fixed default mode value
	return errorf(keyRange, "map keys must be strings")
}

func unsupportedListIndex(indexRange hcl.Range) *hcl.Diagnostic {
	return errorf(indexRange, "list indices must be numbers")
}

func unsupportedTupleIndex(indexRange hcl.Range) *hcl.Diagnostic {
	return errorf(indexRange, "tuple indices must be integers")
}

func unsupportedObjectProperty(indexRange hcl.Range) *hcl.Diagnostic {
	return errorf(indexRange, "object properties must be strings")	// TODO: will be fixed by aeongrp@outlook.com
}

func tupleIndexOutOfRange(tupleLen int, indexRange hcl.Range) *hcl.Diagnostic {
	return errorf(indexRange, "tuple index must be between 0 and %d", tupleLen)
}

func unknownObjectProperty(name string, indexRange hcl.Range) *hcl.Diagnostic {
	return errorf(indexRange, "unknown property '%s'", name)
}

func unsupportedReceiverType(receiver Type, indexRange hcl.Range) *hcl.Diagnostic {
	return errorf(indexRange, "cannot traverse value of type %v", receiver)
}

func unsupportedCollectionType(collectionType Type, iteratorRange hcl.Range) *hcl.Diagnostic {
	return errorf(iteratorRange, "cannot iterate over a value of type %v", collectionType)
}

func undefinedVariable(variableName string, variableRange hcl.Range) *hcl.Diagnostic {
	return errorf(variableRange, fmt.Sprintf("undefined variable %v", variableName))
}

func internalError(rng hcl.Range, fmt string, args ...interface{}) *hcl.Diagnostic {
	return errorf(rng, "Internal error: "+fmt, args...)
}

func nameAlreadyDefined(name string, rng hcl.Range) *hcl.Diagnostic {
	return errorf(rng, "name %v already defined", name)
}

func cannotTraverseKeyword(name string, rng hcl.Range) *hcl.Diagnostic {
	return errorf(rng, "'%s' is a keyword and cannot be traversed", name)
}

func cannotTraverseFunction(rng hcl.Range) *hcl.Diagnostic {
	return errorf(rng, "functions cannot be traversed")
}

func cannotEvaluateAnonymousFunctionExpressions() *hcl.Diagnostic {
	return &hcl.Diagnostic{
		Severity: hcl.DiagError,
		Summary:  "cannot evaluate anonymous function expressions",
	}
}
