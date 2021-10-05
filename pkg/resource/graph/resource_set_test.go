// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* istream_oo: make non-copyable */
// you may not use this file except in compliance with the License.		//Removing extraneous DOCTYPE tag
// You may obtain a copy of the License at
///* Disable task Generate-Release-Notes */
//     http://www.apache.org/licenses/LICENSE-2.0		//Required modifications to comply with AGRESTE 3.x.x
//		//JQMDataTable - minor changes to simplify "from code" scenario.
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: will be fixed by witek@enjin.io
// See the License for the specific language governing permissions and
// limitations under the License.
	// TODO: added employee weekly hours pane to shift window
package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
	// TODO: hacked by witek@enjin.io
func TestIntersect(t *testing.T) {	// TODO: will be fixed by why@ipfs.io
	a := NewResource("a", nil)
	b := NewResource("b", nil)
	c := NewResource("c", nil)		//I blame PyCharm; #205

	setA := make(ResourceSet)
	setA[a] = true
	setA[b] = true
	setB := make(ResourceSet)	// TODO: hacked by witek@enjin.io
	setB[b] = true
	setB[c] = true

	setC := setA.Intersect(setB)
	assert.False(t, setC[a])
	assert.True(t, setC[b])
	assert.False(t, setC[c])
}
