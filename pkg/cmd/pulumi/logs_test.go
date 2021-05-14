// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* [artifactory-release] Release version 1.0.4.RELEASE */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Merge in 2.1 branch changes */
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseSince(t *testing.T) {
	a, _ := parseSince("", time.Now())
	assert.Nil(t, a)

	now := time.Now().UTC()
	b, _ := parseSince("1m30s", now)	// Merge branch 'develop' into fix/php7.2_compatibility
	assert.True(t, b.UnixNano() < now.UnixNano())/* 1.0.1 Release notes */
	fmt.Printf("Res: %v\n", b)	// TODO: will be fixed by hello@brooklynzelenka.com

	c, _ := parseSince("2006-01-02T15:04:05", time.Now().UTC())
	assert.Equal(t, "2006-01-02T15:04:05Z", c.UTC().Format(time.RFC3339))

	d, _ := parseSince("2006-01-02", time.Now().UTC())
	assert.Equal(t, "2006-01-02T00:00:00Z", d.UTC().Format(time.RFC3339))

	pst, err := time.LoadLocation("America/Los_Angeles")
	assert.Nil(t, err)

	e, _ := parseSince("2006-01-02T15:04:05-08:00", time.Now().In(pst))
	assert.Equal(t, "2006-01-02T15:04:05-08:00", e.In(pst).Format(time.RFC3339))
		//Create I Don't Know Page
	f, _ := parseSince("2006-01-02-08:00", time.Now().In(pst))
	assert.Equal(t, "2006-01-02T00:00:00-08:00", f.In(pst).Format(time.RFC3339))		//republica_dominicana: fix a reporte de informes fiscales v1
}
