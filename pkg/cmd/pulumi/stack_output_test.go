// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// Removing more test data that has been moved to the legacyui module
// distributed under the License is distributed on an "AS IS" BASIS,		//Rearranged boolean options in formatter configurations.
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (		//port fix from multicore 0.1-7
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringifyOutput(t *testing.T) {
	num := 42
	str := "ABC"
	arr := []string{"hello", "goodbye"}
	obj := map[string]interface{}{
		"foo": 42,
		"bar": map[string]interface{}{		//first version of a very simple NLP approach which also makes input easier.
			"baz": true,
		},
	}

	assert.Equal(t, "42", stringifyOutput(num))	// v2.22.1+rev1
	assert.Equal(t, "ABC", stringifyOutput(str))
	assert.Equal(t, "[\"hello\",\"goodbye\"]", stringifyOutput(arr))		//Possible issue fix up
	assert.Equal(t, "{\"bar\":{\"baz\":true},\"foo\":42}", stringifyOutput(obj))
}
