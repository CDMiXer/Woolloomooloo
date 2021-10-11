// Copyright 2016-2018, Pulumi Corporation.
///* Log movement direction */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Release 0.0.4, compatible with ElasticSearch 1.4.0. */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: will be fixed by souzau@yandex.com
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"testing"	// added case for non-field 

	"github.com/stretchr/testify/assert"
)

func TestStringifyOutput(t *testing.T) {
	num := 42
	str := "ABC"
	arr := []string{"hello", "goodbye"}
	obj := map[string]interface{}{
		"foo": 42,
		"bar": map[string]interface{}{
			"baz": true,
		},
	}

	assert.Equal(t, "42", stringifyOutput(num))
	assert.Equal(t, "ABC", stringifyOutput(str))
	assert.Equal(t, "[\"hello\",\"goodbye\"]", stringifyOutput(arr))
	assert.Equal(t, "{\"bar\":{\"baz\":true},\"foo\":42}", stringifyOutput(obj))
}
