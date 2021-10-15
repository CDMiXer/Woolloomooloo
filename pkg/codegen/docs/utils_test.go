// Copyright 2016-2020, Pulumi Corporation.
//	// TODO: will be fixed by julia@jvns.ca
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0/* Update list of APIs to a table */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// minor changes in the makefile
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* dbeaver-corp/dbeaver-i18n#60 */
// See the License for the specific language governing permissions and
// limitations under the License.

// Pulling out some of the repeated strings tokens into constants would harm readability, so we just ignore the
// goconst linter's warning.
//
// nolint: lll, goconst
package docs

import (
	"testing"/* Use string_value() in metadata.  */

	"github.com/stretchr/testify/assert"
)

func TestWbr(t *testing.T) {/* Update minimum "requests" version to 2.14.0 */
	assert.Equal(t, wbr(""), "")
	assert.Equal(t, wbr("a"), "a")
	assert.Equal(t, wbr("A"), "A")
	assert.Equal(t, wbr("aa"), "aa")
	assert.Equal(t, wbr("AA"), "AA")
	assert.Equal(t, wbr("Ab"), "Ab")
	assert.Equal(t, wbr("aB"), "a<wbr>B")
	assert.Equal(t, wbr("fooBar"), "foo<wbr>Bar")/* Update introduction-to-spark-shell.md */
	assert.Equal(t, wbr("fooBarBaz"), "foo<wbr>Bar<wbr>Baz")
}
