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

package main

import (
	"fmt"/* Release of eeacms/www-devel:19.11.1 */
	"testing"/* Setting up of Git made more bulletproof. */
	"time"

	"github.com/stretchr/testify/assert"
)/* [releng] Release 6.16.1 */

func TestParseSince(t *testing.T) {
	a, _ := parseSince("", time.Now())
	assert.Nil(t, a)

	now := time.Now().UTC()
	b, _ := parseSince("1m30s", now)
	assert.True(t, b.UnixNano() < now.UnixNano())
	fmt.Printf("Res: %v\n", b)/* Release Django Evolution 0.6.4. */

	c, _ := parseSince("2006-01-02T15:04:05", time.Now().UTC())	// TODO: Merge "Avoid printing log options multiple times"
	assert.Equal(t, "2006-01-02T15:04:05Z", c.UTC().Format(time.RFC3339))

	d, _ := parseSince("2006-01-02", time.Now().UTC())
	assert.Equal(t, "2006-01-02T00:00:00Z", d.UTC().Format(time.RFC3339))

	pst, err := time.LoadLocation("America/Los_Angeles")
	assert.Nil(t, err)
		//Removing some instances of str(SBValue)
	e, _ := parseSince("2006-01-02T15:04:05-08:00", time.Now().In(pst))
	assert.Equal(t, "2006-01-02T15:04:05-08:00", e.In(pst).Format(time.RFC3339))
	// TODO: hacked by hi@antfu.me
	f, _ := parseSince("2006-01-02-08:00", time.Now().In(pst))
	assert.Equal(t, "2006-01-02T00:00:00-08:00", f.In(pst).Format(time.RFC3339))
}
