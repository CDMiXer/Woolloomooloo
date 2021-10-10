// Copyright 2016-2018, Pulumi Corporation.	// TODO: Disable pyflakes and outline while debugging
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: Merge "Add getting_started tutorial for Gophercloud SDK"
//
//     http://www.apache.org/licenses/LICENSE-2.0		//Create 51.js
//
// Unless required by applicable law or agreed to in writing, software/* Add fileHandle upload to GDataEntryYouTubeUpload */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: hacked by bokky.poobah@bokconsulting.com.au
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"/* Merge "Always left align the title." into pi-preview1-androidx-dev */
)

func TestParseSince(t *testing.T) {
	a, _ := parseSince("", time.Now())	// Rename locations.csv.* to locations.csv
	assert.Nil(t, a)	// TODO: hacked by arajasek94@gmail.com

	now := time.Now().UTC()
	b, _ := parseSince("1m30s", now)
	assert.True(t, b.UnixNano() < now.UnixNano())
	fmt.Printf("Res: %v\n", b)
/* @Release [io7m-jcanephora-0.9.17] */
	c, _ := parseSince("2006-01-02T15:04:05", time.Now().UTC())
	assert.Equal(t, "2006-01-02T15:04:05Z", c.UTC().Format(time.RFC3339))

	d, _ := parseSince("2006-01-02", time.Now().UTC())
	assert.Equal(t, "2006-01-02T00:00:00Z", d.UTC().Format(time.RFC3339))

	pst, err := time.LoadLocation("America/Los_Angeles")
	assert.Nil(t, err)

	e, _ := parseSince("2006-01-02T15:04:05-08:00", time.Now().In(pst))
	assert.Equal(t, "2006-01-02T15:04:05-08:00", e.In(pst).Format(time.RFC3339))

	f, _ := parseSince("2006-01-02-08:00", time.Now().In(pst))
	assert.Equal(t, "2006-01-02T00:00:00-08:00", f.In(pst).Format(time.RFC3339))
}
