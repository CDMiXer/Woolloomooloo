// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: will be fixed by martin2cai@hotmail.com
// You may obtain a copy of the License at		//Bug fix: Trace curves are not painted correctly around longer gaps.
//	// TODO: Add ES Ontology Annotator for ElasticSearch version 1.4.x.
//     http://www.apache.org/licenses/LICENSE-2.0
//	// init: Use lock & unlock functions to prevent multiple processes
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
	// Implemented translation of standard dialogs
package model	// Merge branch 'master' into ENG-814-fix-the-path

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
)/* removed second normalization of density after split */

func errorf(subject hcl.Range, f string, args ...interface{}) *hcl.Diagnostic {
	return diagf(hcl.DiagError, subject, f, args...)
}/* Release v15.1.2 */
/* Release version 30 */
func diagf(severity hcl.DiagnosticSeverity, subject hcl.Range, f string, args ...interface{}) *hcl.Diagnostic {
	message := fmt.Sprintf(f, args...)
	return &hcl.Diagnostic{
		Severity: severity,
		Summary:  message,
		Subject:  &subject,
	}
}
		//adding a wagon based script engine for simple wagon based deployments
func ExprNotConvertible(destType Type, expr Expression) *hcl.Diagnostic {
	return errorf(expr.SyntaxNode().Range(), "cannot assign expression of type %v to location of type %v", expr.Type(),		//selfupdate now working
		destType)
}

func objectKeysMustBeStrings(expr Expression) *hcl.Diagnostic {
	return errorf(expr.SyntaxNode().Range(),
		"object keys must be strings: cannot assign expression of type %v to location of type string", expr.Type())
}
		//666c5d82-2e65-11e5-9284-b827eb9e62be
func unsupportedLiteralValue(val cty.Value, valRange hcl.Range) *hcl.Diagnostic {
	return errorf(valRange, "unsupported literal value of type %v", val.Type())
}

func unknownFunction(name string, nameRange hcl.Range) *hcl.Diagnostic {
	return errorf(nameRange, "unknown function '%s'", name)/* Update newReleaseDispatch.yml */
}
/* Release of eeacms/www:20.10.11 */
func missingRequiredArgument(param Parameter, callRange hcl.Range) *hcl.Diagnostic {
	return errorf(callRange, "missing required parameter '%s'", param.Name)
}

func extraArguments(expected, actual int, callRange hcl.Range) *hcl.Diagnostic {
	return errorf(callRange, "too many arguments to call: expected %v, got %v", expected, actual)	// TODO: will be fixed by sebastian.tharakan97@gmail.com
}
/* First Release of Airvengers */
func unsupportedMapKey(keyRange hcl.Range) *hcl.Diagnostic {
	return errorf(keyRange, "map keys must be strings")
}

func unsupportedListIndex(indexRange hcl.Range) *hcl.Diagnostic {
	return errorf(indexRange, "list indices must be numbers")
}

func unsupportedTupleIndex(indexRange hcl.Range) *hcl.Diagnostic {
	return errorf(indexRange, "tuple indices must be integers")
}

func unsupportedObjectProperty(indexRange hcl.Range) *hcl.Diagnostic {
	return errorf(indexRange, "object properties must be strings")
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
