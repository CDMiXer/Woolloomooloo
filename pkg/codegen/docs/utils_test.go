// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Small fix in the template */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//c39e90ae-2e71-11e5-9284-b827eb9e62be
// limitations under the License.

// Pulling out some of the repeated strings tokens into constants would harm readability, so we just ignore the	// TODO: will be fixed by boringland@protonmail.ch
// goconst linter's warning.
//	// TODO: Use markdown checklist for PR template.
// nolint: lll, goconst		//Delete cl_fonts.lua
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
	assert.Equal(t, wbr("Ab"), "Ab")		//Win32 - UpdateHotkeyAssigments() - More hotkeys added.
	assert.Equal(t, wbr("aB"), "a<wbr>B")
	assert.Equal(t, wbr("fooBar"), "foo<wbr>Bar")
	assert.Equal(t, wbr("fooBarBaz"), "foo<wbr>Bar<wbr>Baz")
}
