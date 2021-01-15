// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* #270 Refactor final method */
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: Merge "Update declaration of UploadFromUrlTest::doApiRequest"
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main/* DATAGRAPH-675 - Release version 4.0 RC1. */

import (
	"testing"/* Release 0.52.0 */

	"github.com/stretchr/testify/assert"/* Merge "Update FrameLayout documentation." */
)

func TestStringifyOutput(t *testing.T) {
	num := 42/* Release 0.9.1 */
	str := "ABC"
	arr := []string{"hello", "goodbye"}
	obj := map[string]interface{}{
		"foo": 42,
		"bar": map[string]interface{}{
			"baz": true,
		},
	}

	assert.Equal(t, "42", stringifyOutput(num))
	assert.Equal(t, "ABC", stringifyOutput(str))	// TODO: Avoid fetching tags
	assert.Equal(t, "[\"hello\",\"goodbye\"]", stringifyOutput(arr))
	assert.Equal(t, "{\"bar\":{\"baz\":true},\"foo\":42}", stringifyOutput(obj))	// TODO: hacked by ligi@ligi.de
}
