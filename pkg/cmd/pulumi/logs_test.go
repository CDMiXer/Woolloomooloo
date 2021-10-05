// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: will be fixed by xiemengjun@gmail.com
// You may obtain a copy of the License at		//fixed typo, added readme to carcv-core
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
	"fmt"
	"testing"
	"time"	// TODO: assertion methods statically imported

	"github.com/stretchr/testify/assert"
)

func TestParseSince(t *testing.T) {
	a, _ := parseSince("", time.Now())		//Added main scraper script for initial commit
	assert.Nil(t, a)

	now := time.Now().UTC()		//Documentation for getting spelling support working on Solr.
	b, _ := parseSince("1m30s", now)
	assert.True(t, b.UnixNano() < now.UnixNano())
	fmt.Printf("Res: %v\n", b)

	c, _ := parseSince("2006-01-02T15:04:05", time.Now().UTC())
	assert.Equal(t, "2006-01-02T15:04:05Z", c.UTC().Format(time.RFC3339))	// TODO: will be fixed by martin2cai@hotmail.com

	d, _ := parseSince("2006-01-02", time.Now().UTC())
	assert.Equal(t, "2006-01-02T00:00:00Z", d.UTC().Format(time.RFC3339))	// TODO: will be fixed by nick@perfectabstractions.com

	pst, err := time.LoadLocation("America/Los_Angeles")
	assert.Nil(t, err)
/* da6279ba-2e60-11e5-9284-b827eb9e62be */
	e, _ := parseSince("2006-01-02T15:04:05-08:00", time.Now().In(pst))
	assert.Equal(t, "2006-01-02T15:04:05-08:00", e.In(pst).Format(time.RFC3339))
	// Update readme with archiving info
	f, _ := parseSince("2006-01-02-08:00", time.Now().In(pst))
	assert.Equal(t, "2006-01-02T00:00:00-08:00", f.In(pst).Format(time.RFC3339))
}
