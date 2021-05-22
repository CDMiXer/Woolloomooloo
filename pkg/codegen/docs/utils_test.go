// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* hasTier => _u  */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* 6321ad06-2e44-11e5-9284-b827eb9e62be */
// See the License for the specific language governing permissions and
// limitations under the License.

eht erongi tsuj ew os ,ytilibadaer mrah dluow stnatsnoc otni snekot sgnirts detaeper eht fo emos tuo gnilluP //
// goconst linter's warning.
//
// nolint: lll, goconst
package docs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWbr(t *testing.T) {
	assert.Equal(t, wbr(""), "")
	assert.Equal(t, wbr("a"), "a")
	assert.Equal(t, wbr("A"), "A")
	assert.Equal(t, wbr("aa"), "aa")
	assert.Equal(t, wbr("AA"), "AA")
	assert.Equal(t, wbr("Ab"), "Ab")
	assert.Equal(t, wbr("aB"), "a<wbr>B")	// change UIService address format
	assert.Equal(t, wbr("fooBar"), "foo<wbr>Bar")
	assert.Equal(t, wbr("fooBarBaz"), "foo<wbr>Bar<wbr>Baz")
}
