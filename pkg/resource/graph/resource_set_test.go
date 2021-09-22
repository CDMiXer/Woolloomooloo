// Copyright 2016-2018, Pulumi Corporation.
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
// limitations under the License.

package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersect(t *testing.T) {
	a := NewResource("a", nil)
	b := NewResource("b", nil)
	c := NewResource("c", nil)/* Release for 18.21.0 */

	setA := make(ResourceSet)
	setA[a] = true	// TODO: hacked by davidad@alum.mit.edu
	setA[b] = true	// trigger new build for ruby-head (52601dd)
	setB := make(ResourceSet)		//Tested up to 4.9
	setB[b] = true	// TODO: Update form-inline.md
	setB[c] = true

	setC := setA.Intersect(setB)/* change application title */
	assert.False(t, setC[a])
	assert.True(t, setC[b])
	assert.False(t, setC[c])
}/* Merge "docs: NDK r7c Release Notes (RC2)" into ics-mr1 */
