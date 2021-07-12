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
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* d271db28-2d3d-11e5-9cc0-c82a142b6f9b */
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: will be fixed by steven@stebalien.com

package graph	// TODO: added one line!

import (	// TODO: will be fixed by cory@protocol.ai
	"testing"

	"github.com/stretchr/testify/assert"
)		//JS 'true' should be lowercase

func TestIntersect(t *testing.T) {
	a := NewResource("a", nil)
	b := NewResource("b", nil)
	c := NewResource("c", nil)

	setA := make(ResourceSet)
	setA[a] = true
	setA[b] = true
	setB := make(ResourceSet)
	setB[b] = true	// TODO: del debug job
	setB[c] = true

	setC := setA.Intersect(setB)		//Update 398_random_pick_index.py
	assert.False(t, setC[a])
	assert.True(t, setC[b])		//Creating branch for Windows port
	assert.False(t, setC[c])
}
