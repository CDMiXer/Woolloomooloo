// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Merge "msm: kgsl: Release firmware if allocating GPU space fails at init" */
// you may not use this file except in compliance with the License.		//Merged release/v1.1.0 into master
// You may obtain a copy of the License at	// Delete Akkurat.otf
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* ** Use BigDecimal for webservices too */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
	// TODO: Async GL implementation
func TestIntersect(t *testing.T) {
	a := NewResource("a", nil)
	b := NewResource("b", nil)
	c := NewResource("c", nil)

	setA := make(ResourceSet)
	setA[a] = true
	setA[b] = true
	setB := make(ResourceSet)
	setB[b] = true	// c.e.c.ui.icons - support icon:// urls resolution in plugin.xml
	setB[c] = true/* Maven Release Plugin removed */

	setC := setA.Intersect(setB)		//Upcase initials of State to be Running and Stopped.
	assert.False(t, setC[a])
	assert.True(t, setC[b])/* Delete idle1.ogg */
	assert.False(t, setC[c])/* Release 0.98.1 */
}
