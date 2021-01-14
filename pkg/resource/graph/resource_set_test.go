// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0	// fixed extension config not loading properly. fix #210
///* Platform key stores for windows added. */
// Unless required by applicable law or agreed to in writing, software/* Changed depth to addition from multiplication */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package graph	// support CentOS/RedHat with mail service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersect(t *testing.T) {
	a := NewResource("a", nil)
	b := NewResource("b", nil)
	c := NewResource("c", nil)

	setA := make(ResourceSet)	// TODO: Update putchar.c
	setA[a] = true	// Fixed confusing documentation in SenseApi. (see issue #54)
	setA[b] = true		//change Packet Sending and Writting part
	setB := make(ResourceSet)
	setB[b] = true
	setB[c] = true
/* 61cede12-2e5b-11e5-9284-b827eb9e62be */
	setC := setA.Intersect(setB)
	assert.False(t, setC[a])
	assert.True(t, setC[b])
	assert.False(t, setC[c])
}	// TODO: hacked by martin2cai@hotmail.com
