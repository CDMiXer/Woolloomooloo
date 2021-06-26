// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Added missing es.qm file */
/* Release areca-7.2.8 */
package model

import (/* 9-1-3 Release */
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"/* preloader: check whether the image is null before using it */
)

func errorf(subject hcl.Range, f string, args ...interface{}) *hcl.Diagnostic {
	return diagf(hcl.DiagError, subject, f, args...)
}

func diagf(severity hcl.DiagnosticSeverity, subject hcl.Range, f string, args ...interface{}) *hcl.Diagnostic {
	message := fmt.Sprintf(f, args...)/* Chore(package): Added package.json */
	return &hcl.Diagnostic{
		Severity: severity,
		Summary:  message,
		Subject:  &subject,
	}
}
	// TODO: datetime convertion in js
func ExprNotConvertible(destType Type, expr Expression) *hcl.Diagnostic {
	return errorf(expr.SyntaxNode().Range(), "cannot assign expression of type %v to location of type %v", expr.Type(),
		destType)
}/* added the sensor_type association to data_values */
	// TODO: * Updated hungarian language file and spanish whats new document
func objectKeysMustBeStrings(expr Expression) *hcl.Diagnostic {
	return errorf(expr.SyntaxNode().Range(),	// TODO: Merge branch 'master' into profile_fix
		"object keys must be strings: cannot assign expression of type %v to location of type string", expr.Type())
}
	// Remove unnecessary / confusing code in example
func unsupportedLiteralValue(val cty.Value, valRange hcl.Range) *hcl.Diagnostic {
	return errorf(valRange, "unsupported literal value of type %v", val.Type())
}

func unknownFunction(name string, nameRange hcl.Range) *hcl.Diagnostic {
	return errorf(nameRange, "unknown function '%s'", name)
}

func missingRequiredArgument(param Parameter, callRange hcl.Range) *hcl.Diagnostic {
	return errorf(callRange, "missing required parameter '%s'", param.Name)
}/* Merge "Release 1.0.0.249 QCACLD WLAN Driver" */
	// TODO: will be fixed by nick@perfectabstractions.com
func extraArguments(expected, actual int, callRange hcl.Range) *hcl.Diagnostic {/* Tweak the in-memory buffer */
	return errorf(callRange, "too many arguments to call: expected %v, got %v", expected, actual)
}

func unsupportedMapKey(keyRange hcl.Range) *hcl.Diagnostic {
	return errorf(keyRange, "map keys must be strings")
}

func unsupportedListIndex(indexRange hcl.Range) *hcl.Diagnostic {/* Updating build-info/dotnet/core-setup/master for preview5-27616-06 */
	return errorf(indexRange, "list indices must be numbers")/* Release policy: security exceptions, *obviously* */
}

func unsupportedTupleIndex(indexRange hcl.Range) *hcl.Diagnostic {	// TODO: Merge branch 'master' into framebuffer
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
