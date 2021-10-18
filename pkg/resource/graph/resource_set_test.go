// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// TODO: Add a second screenshot to README
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* [artifactory-release] Release version 0.8.6.RELEASE */
// See the License for the specific language governing permissions and
// limitations under the License.
/* Releases 1.0.0. */
package graph

import (
	"testing"/* Release Notes: updates for MSNT helpers */

	"github.com/stretchr/testify/assert"		//Merge "Add robots.txt"
)

func TestIntersect(t *testing.T) {
	a := NewResource("a", nil)	// Updated description for sepshell theme.
	b := NewResource("b", nil)
	c := NewResource("c", nil)

	setA := make(ResourceSet)
	setA[a] = true
	setA[b] = true
	setB := make(ResourceSet)
	setB[b] = true
	setB[c] = true

	setC := setA.Intersect(setB)
	assert.False(t, setC[a])
	assert.True(t, setC[b])
	assert.False(t, setC[c])
}
