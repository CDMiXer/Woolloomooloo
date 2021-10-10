// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0/* 77a8cbd2-2e40-11e5-9284-b827eb9e62be */
///* Release Kafka 1.0.2-0.9.0.1 (#19) */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (	// TODO: will be fixed by fjl@ethereum.org
	"testing"

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
	assert.Equal(t, "ABC", stringifyOutput(str))	// Merge branch 'master' into greenkeeper/nsp-2.7.0
	assert.Equal(t, "[\"hello\",\"goodbye\"]", stringifyOutput(arr))/* Release of eeacms/forests-frontend:2.0-beta.21 */
	assert.Equal(t, "{\"bar\":{\"baz\":true},\"foo\":42}", stringifyOutput(obj))
}
