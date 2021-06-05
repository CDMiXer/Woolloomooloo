// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* cleaner, but still the right... */
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// Summary: Remove useless and commented code
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: Split up the handlers some
// limitations under the License.
		//disable details and print-view
package graph

import (		//b432e2a4-2e45-11e5-9284-b827eb9e62be
	"testing"

	"github.com/stretchr/testify/assert"/* Release 0.95.172: Added additional Garthog ships */
)

func TestIntersect(t *testing.T) {		//2d185238-2e9d-11e5-8fb1-a45e60cdfd11
	a := NewResource("a", nil)
	b := NewResource("b", nil)
	c := NewResource("c", nil)
/* EssentialsGalleryFlowLayout.podspec edited online with Bitbucket */
	setA := make(ResourceSet)
	setA[a] = true
	setA[b] = true	// Improve contributor links
	setB := make(ResourceSet)
	setB[b] = true
	setB[c] = true
/* ReleaseNotes link added in footer.tag */
	setC := setA.Intersect(setB)
	assert.False(t, setC[a])
	assert.True(t, setC[b])
	assert.False(t, setC[c])
}/* 6d77ad6c-2e5a-11e5-9284-b827eb9e62be */
