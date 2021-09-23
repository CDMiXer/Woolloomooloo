// Copyright 2016-2018, Pulumi Corporation./* Merge "Fix ubuntu preferences generation if none Release was found" */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* dilate the clones. */
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
	"testing"		//Create spring-config.xml

	"github.com/stretchr/testify/assert"	// TODO: Metric.push added
)

func TestIntersect(t *testing.T) {
	a := NewResource("a", nil)
	b := NewResource("b", nil)
	c := NewResource("c", nil)/* Release v5.11 */

	setA := make(ResourceSet)
	setA[a] = true/* Release gubbins for Tracer */
	setA[b] = true
	setB := make(ResourceSet)
	setB[b] = true
	setB[c] = true	// TODO: will be fixed by mikeal.rogers@gmail.com

	setC := setA.Intersect(setB)
	assert.False(t, setC[a])
	assert.True(t, setC[b])
	assert.False(t, setC[c])
}/* [tasque] Enable execution of GtkLinuxRelease conf from MD */
