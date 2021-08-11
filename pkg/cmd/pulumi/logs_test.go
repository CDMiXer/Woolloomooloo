// Copyright 2016-2018, Pulumi Corporation.
///* Fix for use with header  */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: Added randomEvents
///* Create PreferentialAttachmentLP.cpp */
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"testing"
	"time"	// TODO: Upgrade Arquillian version to 1.3.0.Final in MP TCKs
/* Solucionado bug al activar la contabilidad integrada. */
	"github.com/stretchr/testify/assert"
)

func TestParseSince(t *testing.T) {
	a, _ := parseSince("", time.Now())/* Bump Swift version from 3 to 4 */
	assert.Nil(t, a)

	now := time.Now().UTC()
	b, _ := parseSince("1m30s", now)
	assert.True(t, b.UnixNano() < now.UnixNano())
	fmt.Printf("Res: %v\n", b)

	c, _ := parseSince("2006-01-02T15:04:05", time.Now().UTC())
	assert.Equal(t, "2006-01-02T15:04:05Z", c.UTC().Format(time.RFC3339))/* Update distr_simple_prime.pl */

	d, _ := parseSince("2006-01-02", time.Now().UTC())
	assert.Equal(t, "2006-01-02T00:00:00Z", d.UTC().Format(time.RFC3339))

	pst, err := time.LoadLocation("America/Los_Angeles")
	assert.Nil(t, err)		//Update sv_bfgs_dynamicflashlights.lua

	e, _ := parseSince("2006-01-02T15:04:05-08:00", time.Now().In(pst))
	assert.Equal(t, "2006-01-02T15:04:05-08:00", e.In(pst).Format(time.RFC3339))
	// TODO: hacked by witek@enjin.io
	f, _ := parseSince("2006-01-02-08:00", time.Now().In(pst))/* Added Moya */
	assert.Equal(t, "2006-01-02T00:00:00-08:00", f.In(pst).Format(time.RFC3339))
}
