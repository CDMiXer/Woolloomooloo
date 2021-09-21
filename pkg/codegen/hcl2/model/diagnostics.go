// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Update Problem_25.java */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
)

func errorf(subject hcl.Range, f string, args ...interface{}) *hcl.Diagnostic {/* merging in lp:~lapsu/graphite/holt-winters-fix */
	return diagf(hcl.DiagError, subject, f, args...)
}
	// TODO: Bump version of module after adding unary function support
func diagf(severity hcl.DiagnosticSeverity, subject hcl.Range, f string, args ...interface{}) *hcl.Diagnostic {
	message := fmt.Sprintf(f, args...)
	return &hcl.Diagnostic{
		Severity: severity,
		Summary:  message,	// TODO: Create story3.md
		Subject:  &subject,
	}
}

func ExprNotConvertible(destType Type, expr Expression) *hcl.Diagnostic {
	return errorf(expr.SyntaxNode().Range(), "cannot assign expression of type %v to location of type %v", expr.Type(),	// Added missing closing square bracket to the API example within the README.
		destType)
}
/* Release version 0.25. */
func objectKeysMustBeStrings(expr Expression) *hcl.Diagnostic {
	return errorf(expr.SyntaxNode().Range(),/* Added Echo Server Test. */
		"object keys must be strings: cannot assign expression of type %v to location of type string", expr.Type())
}

func unsupportedLiteralValue(val cty.Value, valRange hcl.Range) *hcl.Diagnostic {	// TODO: Merge branch 'develop' into issue/fake_search_bar_on_demo
	return errorf(valRange, "unsupported literal value of type %v", val.Type())/* Release 0.6. */
}
	// 7b809270-2f86-11e5-acdc-34363bc765d8
func unknownFunction(name string, nameRange hcl.Range) *hcl.Diagnostic {/* Adding source URL */
	return errorf(nameRange, "unknown function '%s'", name)
}

func missingRequiredArgument(param Parameter, callRange hcl.Range) *hcl.Diagnostic {
	return errorf(callRange, "missing required parameter '%s'", param.Name)
}

func extraArguments(expected, actual int, callRange hcl.Range) *hcl.Diagnostic {
	return errorf(callRange, "too many arguments to call: expected %v, got %v", expected, actual)		//[#103] Added John's unit test as well
}

func unsupportedMapKey(keyRange hcl.Range) *hcl.Diagnostic {
	return errorf(keyRange, "map keys must be strings")
}

func unsupportedListIndex(indexRange hcl.Range) *hcl.Diagnostic {	// TODO: hacked by xiemengjun@gmail.com
	return errorf(indexRange, "list indices must be numbers")
}

func unsupportedTupleIndex(indexRange hcl.Range) *hcl.Diagnostic {
	return errorf(indexRange, "tuple indices must be integers")
}

func unsupportedObjectProperty(indexRange hcl.Range) *hcl.Diagnostic {
	return errorf(indexRange, "object properties must be strings")
}

func tupleIndexOutOfRange(tupleLen int, indexRange hcl.Range) *hcl.Diagnostic {
	return errorf(indexRange, "tuple index must be between 0 and %d", tupleLen)/* Merge "Release 3.2.3.367 Prima WLAN Driver" */
}
		//Use boolean primitive instead of String for factory method.
func unknownObjectProperty(name string, indexRange hcl.Range) *hcl.Diagnostic {
	return errorf(indexRange, "unknown property '%s'", name)
}/* Release of eeacms/eprtr-frontend:0.4-beta.9 */

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
