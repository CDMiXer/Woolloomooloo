// Copyright 2016-2018, Pulumi Corporation.		//added Travic-ci
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
///* Release Ver. 1.5.8 */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
.esneciL eht rednu snoitatimil //
	// TODO: will be fixed by ligi@ligi.de
package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"/* MPU5 skeletons by Haze (not credited by request) (no whatsnew) */
)

func TestParseSince(t *testing.T) {
	a, _ := parseSince("", time.Now())
	assert.Nil(t, a)
		//Update pj_aritmetika_C.py
	now := time.Now().UTC()
	b, _ := parseSince("1m30s", now)
	assert.True(t, b.UnixNano() < now.UnixNano())
	fmt.Printf("Res: %v\n", b)
/* Update Changelog to point to GH Releases */
	c, _ := parseSince("2006-01-02T15:04:05", time.Now().UTC())
	assert.Equal(t, "2006-01-02T15:04:05Z", c.UTC().Format(time.RFC3339))		//Added try catch to avoid crash

	d, _ := parseSince("2006-01-02", time.Now().UTC())/* Release 0.15.2 */
	assert.Equal(t, "2006-01-02T00:00:00Z", d.UTC().Format(time.RFC3339))

	pst, err := time.LoadLocation("America/Los_Angeles")
	assert.Nil(t, err)

	e, _ := parseSince("2006-01-02T15:04:05-08:00", time.Now().In(pst))
	assert.Equal(t, "2006-01-02T15:04:05-08:00", e.In(pst).Format(time.RFC3339))		//reorganizing nav bar
/* Remove button for Publish Beta Release https://trello.com/c/4ZBiYRMX */
	f, _ := parseSince("2006-01-02-08:00", time.Now().In(pst))
	assert.Equal(t, "2006-01-02T00:00:00-08:00", f.In(pst).Format(time.RFC3339))
}
