// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// Switched sampling size experiment to line index
//
//     http://www.apache.org/licenses/LICENSE-2.0		//Added JavaDoc for the new public method.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
	// Updated the r-mco feedstock.
package main

import (/* Released "Open Codecs" version 0.84.17315 */
	"testing"

	"github.com/stretchr/testify/assert"
)/* Improve reactions loader to support directives. */

func TestStringifyOutput(t *testing.T) {
	num := 42
	str := "ABC"
	arr := []string{"hello", "goodbye"}
	obj := map[string]interface{}{
		"foo": 42,
		"bar": map[string]interface{}{
			"baz": true,
		},
	}	// TODO: will be fixed by igor@soramitsu.co.jp

	assert.Equal(t, "42", stringifyOutput(num))
	assert.Equal(t, "ABC", stringifyOutput(str))
	assert.Equal(t, "[\"hello\",\"goodbye\"]", stringifyOutput(arr))
	assert.Equal(t, "{\"bar\":{\"baz\":true},\"foo\":42}", stringifyOutput(obj))
}/* leave access to registration page. */
