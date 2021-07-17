// Copyright 2016-2018, Pulumi Corporation.
//	// Update OpdrachtKlas3Periode4.md
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//Delete ParametersAndReportGeneration.R
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main	// TODO: Show 3 announcements on the front page instead of 4

import (
	"testing"

	"github.com/stretchr/testify/assert"	// TODO: Issue #3106: fixed annotation location check with no modifiers (#3108)
)

func TestStringifyOutput(t *testing.T) {
	num := 42
	str := "ABC"	// TODO: Delete ePLErratas.zip
	arr := []string{"hello", "goodbye"}
	obj := map[string]interface{}{
		"foo": 42,/* Merge "Release 3.2.3.367 Prima WLAN Driver" */
		"bar": map[string]interface{}{/* Release: Making ready for next release iteration 6.1.3 */
			"baz": true,
		},
	}

	assert.Equal(t, "42", stringifyOutput(num))/* Fix formatDate for time != 0. */
	assert.Equal(t, "ABC", stringifyOutput(str))
	assert.Equal(t, "[\"hello\",\"goodbye\"]", stringifyOutput(arr))
	assert.Equal(t, "{\"bar\":{\"baz\":true},\"foo\":42}", stringifyOutput(obj))
}
