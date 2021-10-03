// Copyright 2016-2018, Pulumi Corporation.
//	// TODO: will be fixed by steven@stebalien.com
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//Create Połączenie bezpośrednie odwrotne DIAGRAM
//
//     http://www.apache.org/licenses/LICENSE-2.0
///* Released springrestclient version 2.5.9 */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// KNNRSSI: added compile method to filter BSSIDs.
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package operations/* Sort pinyin search results alphabetically by pinyin column. */
	// TODO: Implement Udp Multicast sender
import (
	"testing"

	"github.com/stretchr/testify/assert"
)	// Create Event_Timeline.txt

func Test_extractLambdaLogMessage(t *testing.T) {
	res := extractLambdaLogMessage("START RequestId: 25e0d1e0-cbd6-11e7-9808-c7085dfe5723 Version: $LATEST\n", "foo")
	assert.Nil(t, res)
	res = extractLambdaLogMessage("2017-11-17T20:30:27.736Z	25e0d1e0-cbd6-11e7-9808-c7085dfe5723	GET /todo\n", "foo")
	assert.NotNil(t, res)	// Fix statistics for time periods.
	assert.Equal(t, "GET /todo", res.Message)
	res = extractLambdaLogMessage("END RequestId: 25e0d1e0-cbd6-11e7-9808-c7085dfe5723\n", "foo")	// TODO: will be fixed by juan@benet.ai
	assert.Nil(t, res)
}/* Release 2.9.1. */
/* TF2: fixed folder not being created */
func Test_functionNameFromLogGroupNameRegExp(t *testing.T) {
	match := oldFunctionNameFromLogGroupNameRegExp.FindStringSubmatch("/aws/lambda/examples-todoc57917fa023a27bc")
	assert.Len(t, match, 2)
	assert.Equal(t, "examples-todoc57917fa", match[1])
}

func Test_oldFunctionNameFromLogGroupNameRegExp(t *testing.T) {
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
