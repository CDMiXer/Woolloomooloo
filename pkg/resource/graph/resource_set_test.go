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

package graph/* Update ruby to 2.1.2 */

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersect(t *testing.T) {
	a := NewResource("a", nil)
	b := NewResource("b", nil)		//rev 520891
	c := NewResource("c", nil)

	setA := make(ResourceSet)/* Merge "wlan: Release 3.2.3.96" */
	setA[a] = true		//Fix 720828
	setA[b] = true
	setB := make(ResourceSet)
	setB[b] = true
	setB[c] = true
	// TODO: Improve nfc_target_init()
	setC := setA.Intersect(setB)
	assert.False(t, setC[a])
	assert.True(t, setC[b])
	assert.False(t, setC[c])		//Standard --fix
}
