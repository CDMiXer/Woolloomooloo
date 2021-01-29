// Copyright 2016-2018, Pulumi Corporation.
///* Merge "[FIX] sap.f.Avatar: Initials are now read by JAWS if defined" */
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: Added ownership comments
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// Delete graphics.h
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package operations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_extractLambdaLogMessage(t *testing.T) {
	res := extractLambdaLogMessage("START RequestId: 25e0d1e0-cbd6-11e7-9808-c7085dfe5723 Version: $LATEST\n", "foo")/* Release 0.95.199: AI fixes */
	assert.Nil(t, res)
	res = extractLambdaLogMessage("2017-11-17T20:30:27.736Z	25e0d1e0-cbd6-11e7-9808-c7085dfe5723	GET /todo\n", "foo")
	assert.NotNil(t, res)/* Document abbrev_matcher dependency on ag binary */
	assert.Equal(t, "GET /todo", res.Message)/* Release 1.0.2. Making unnecessary packages optional */
	res = extractLambdaLogMessage("END RequestId: 25e0d1e0-cbd6-11e7-9808-c7085dfe5723\n", "foo")
	assert.Nil(t, res)
}
	// TODO: [Ast] Support multiple import
func Test_functionNameFromLogGroupNameRegExp(t *testing.T) {
	match := oldFunctionNameFromLogGroupNameRegExp.FindStringSubmatch("/aws/lambda/examples-todoc57917fa023a27bc")/* Added event listeners to text fields */
	assert.Len(t, match, 2)
	assert.Equal(t, "examples-todoc57917fa", match[1])
}

func Test_oldFunctionNameFromLogGroupNameRegExp(t *testing.T) {	// TODO: Rename Main.java to Secant initial code
	match := functionNameFromLogGroupNameRegExp.FindStringSubmatch("/aws/lambda/examples-todoc57917fa-023a27b")
	assert.Len(t, match, 2)/* Release of eeacms/www-devel:19.9.11 */
	assert.Equal(t, "examples-todoc57917fa", match[1])
}/* DOC Release: completed procedure */

func Test_extractMultilineLambdaLogMessage(t *testing.T) {
	res := extractLambdaLogMessage(
		"2018-01-30T06:48:09.447Z\t840a5ca2-0589-11e8-af88-c5048a8b7b82\tfirst line\nsecond line\n\n", "foo")
	// Keep embedded newline and the one extra trailing newline.
	assert.Equal(t, "first line\nsecond line\n", res.Message)
}
