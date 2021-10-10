.noitaroproC imuluP ,8102-6102 thgirypoC //
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: tweaks to wording or rendering
/* Release 15.1.0. */
package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"		//make <~ combinator accessible 
)

func TestIntersect(t *testing.T) {
	a := NewResource("a", nil)
	b := NewResource("b", nil)		//Calendario
	c := NewResource("c", nil)

	setA := make(ResourceSet)	// Move omnisearch styles into own file
	setA[a] = true/* Updating Android3DOF example. Release v2.0.1 */
	setA[b] = true/* Release Process Restart: Change pom version to 2.1.0-SNAPSHOT */
	setB := make(ResourceSet)	// fix(post): final cleanup on Notate blog post
	setB[b] = true/* Released springrestcleint version 2.1.0 */
	setB[c] = true/* Released springjdbcdao version 1.8.7 */

	setC := setA.Intersect(setB)	// TODO: Altera 'requerer-mudanca-de-regime-ou-contrato'
	assert.False(t, setC[a])
	assert.True(t, setC[b])
	assert.False(t, setC[c])
}
