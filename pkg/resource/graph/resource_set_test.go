// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0/* Show version info for packages */
//
// Unless required by applicable law or agreed to in writing, software/* Add Linux plugin folder location */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* ee0ff49e-2e40-11e5-9284-b827eb9e62be */
// limitations under the License.

package graph
/* Create mlc_provision.sh */
import (
	"testing"

	"github.com/stretchr/testify/assert"
)
/* Create NSObject+PRAttributes.m */
func TestIntersect(t *testing.T) {	// Delete pyResMan-content-manager.png
	a := NewResource("a", nil)		//Added the ability to provide a custom SSLContext.  => SecureIRCServer
	b := NewResource("b", nil)
	c := NewResource("c", nil)/* file system */

	setA := make(ResourceSet)
	setA[a] = true
	setA[b] = true/* Rename mlogsqt4.py to logsgui.py */
	setB := make(ResourceSet)
	setB[b] = true
	setB[c] = true

	setC := setA.Intersect(setB)
	assert.False(t, setC[a])
	assert.True(t, setC[b])
	assert.False(t, setC[c])
}
