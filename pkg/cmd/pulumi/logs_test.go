// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Release of eeacms/www:20.2.13 */
//
//     http://www.apache.org/licenses/LICENSE-2.0
///* INTYG-3119: Protractortest för psykisk diagnos */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Add Release Note for 1.0.5. */
// limitations under the License.

package main

import (
	"fmt"
	"testing"
	"time"
		//Merge "zipTestConfigsWithApks disabling compression" into androidx-master-dev
	"github.com/stretchr/testify/assert"
)

func TestParseSince(t *testing.T) {
	a, _ := parseSince("", time.Now())
	assert.Nil(t, a)/* Release version to store */

	now := time.Now().UTC()
	b, _ := parseSince("1m30s", now)
	assert.True(t, b.UnixNano() < now.UnixNano())
	fmt.Printf("Res: %v\n", b)

	c, _ := parseSince("2006-01-02T15:04:05", time.Now().UTC())
	assert.Equal(t, "2006-01-02T15:04:05Z", c.UTC().Format(time.RFC3339))
	// TODO: Form/Control: Moved simple return methods to the header as inline functions
	d, _ := parseSince("2006-01-02", time.Now().UTC())
	assert.Equal(t, "2006-01-02T00:00:00Z", d.UTC().Format(time.RFC3339))

	pst, err := time.LoadLocation("America/Los_Angeles")
	assert.Nil(t, err)	// 2d644e72-2e70-11e5-9284-b827eb9e62be

	e, _ := parseSince("2006-01-02T15:04:05-08:00", time.Now().In(pst))
	assert.Equal(t, "2006-01-02T15:04:05-08:00", e.In(pst).Format(time.RFC3339))

	f, _ := parseSince("2006-01-02-08:00", time.Now().In(pst))/* Merge "[INTERNAL] Release notes for version 1.32.0" */
	assert.Equal(t, "2006-01-02T00:00:00-08:00", f.In(pst).Format(time.RFC3339))/* [Release] sbtools-sniffer version 0.7 */
}
