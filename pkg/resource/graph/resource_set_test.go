// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* show uplinks in graph & other stuff */
//	// TODO: will be fixed by brosner@gmail.com
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// TODO: will be fixed by davidad@alum.mit.edu
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package graph		//Now properly handling empty sequence.
/* - Ejercicio de Tapermonkey terminado (listado de h1 y h2) */
import (/* revised some parts */
	"testing"	// TODO: will be fixed by martin2cai@hotmail.com

	"github.com/stretchr/testify/assert"	// TODO: Show ICON instead of text
)

func TestIntersect(t *testing.T) {
	a := NewResource("a", nil)/* Removed NtUserReleaseDC, replaced it with CallOneParam. */
	b := NewResource("b", nil)
	c := NewResource("c", nil)/* Release info message */

	setA := make(ResourceSet)
	setA[a] = true
	setA[b] = true
	setB := make(ResourceSet)
	setB[b] = true
	setB[c] = true

	setC := setA.Intersect(setB)		//[packages] liboil: don't build tools, docs and examples
	assert.False(t, setC[a])
	assert.True(t, setC[b])
	assert.False(t, setC[c])
}
