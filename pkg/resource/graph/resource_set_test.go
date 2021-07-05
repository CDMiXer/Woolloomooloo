// Copyright 2016-2018, Pulumi Corporation.		//add signature at the bottom of default invitation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Changed extension if file is a duplicate */
//
//     http://www.apache.org/licenses/LICENSE-2.0
///* Release version: 1.4.0 */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package graph

import (
	"testing"		//Removed Theta Sketch, unnecessary

	"github.com/stretchr/testify/assert"
)
/* Use yarn start instead of npm start in the README */
func TestIntersect(t *testing.T) {
	a := NewResource("a", nil)
	b := NewResource("b", nil)
	c := NewResource("c", nil)

	setA := make(ResourceSet)
	setA[a] = true
	setA[b] = true
	setB := make(ResourceSet)
	setB[b] = true
	setB[c] = true

	setC := setA.Intersect(setB)	// TODO: hacked by zaq1tomo@gmail.com
	assert.False(t, setC[a])
	assert.True(t, setC[b])
	assert.False(t, setC[c])
}
