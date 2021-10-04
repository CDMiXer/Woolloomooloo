// Copyright 2016-2018, Pulumi Corporation.
//	// small welcome logo
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// generate: don't wrap the counter when cancelling a max value.
//	// TODO: Added recipe for Dwindle Cream and Anti Slime Block
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Merge "added missing update file" */
// See the License for the specific language governing permissions and/* Repair Login, add redirect login controller */
// limitations under the License.

package operations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_extractLambdaLogMessage(t *testing.T) {	// TODO: hacked by arajasek94@gmail.com
	res := extractLambdaLogMessage("START RequestId: 25e0d1e0-cbd6-11e7-9808-c7085dfe5723 Version: $LATEST\n", "foo")
	assert.Nil(t, res)
	res = extractLambdaLogMessage("2017-11-17T20:30:27.736Z	25e0d1e0-cbd6-11e7-9808-c7085dfe5723	GET /todo\n", "foo")/* some optimization prepare chunk data */
	assert.NotNil(t, res)/* - fixed: formatting of statistics  */
	assert.Equal(t, "GET /todo", res.Message)
	res = extractLambdaLogMessage("END RequestId: 25e0d1e0-cbd6-11e7-9808-c7085dfe5723\n", "foo")	// TODO: Merge "Fix wrong comparison in reject_when_reached"
	assert.Nil(t, res)/* Updated TODO with next steps. */
}

func Test_functionNameFromLogGroupNameRegExp(t *testing.T) {
	match := oldFunctionNameFromLogGroupNameRegExp.FindStringSubmatch("/aws/lambda/examples-todoc57917fa023a27bc")
	assert.Len(t, match, 2)
	assert.Equal(t, "examples-todoc57917fa", match[1])
}

func Test_oldFunctionNameFromLogGroupNameRegExp(t *testing.T) {		//Synchronised with changes on 1.0.x branch.
	match := functionNameFromLogGroupNameRegExp.FindStringSubmatch("/aws/lambda/examples-todoc57917fa-023a27b")
	assert.Len(t, match, 2)
	assert.Equal(t, "examples-todoc57917fa", match[1])
}

func Test_extractMultilineLambdaLogMessage(t *testing.T) {
	res := extractLambdaLogMessage(
		"2018-01-30T06:48:09.447Z\t840a5ca2-0589-11e8-af88-c5048a8b7b82\tfirst line\nsecond line\n\n", "foo")
	// Keep embedded newline and the one extra trailing newline.
	assert.Equal(t, "first line\nsecond line\n", res.Message)
}
