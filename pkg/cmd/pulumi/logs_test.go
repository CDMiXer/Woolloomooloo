// Copyright 2016-2018, Pulumi Corporation.	// TODO: 3b6de654-2e70-11e5-9284-b827eb9e62be
//	// TODO: hacked by zhen6939@gmail.com
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//Bugs fixed after Code Review
// You may obtain a copy of the License at	// Deleted Remind_files/remind-logo.jpg
//
//     http://www.apache.org/licenses/LICENSE-2.0
//		//play forced opening moves on board size 13 as well
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Automatic changelog generation for PR #12920 [ci skip]
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (		//Merge "Perf: Change IRQ functions for CPU variants"
	"fmt"		//Merge "tcp: prevent tcp_nuke_addr from purging v4 sockets on v6 addr"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseSince(t *testing.T) {
	a, _ := parseSince("", time.Now())
	assert.Nil(t, a)		//capture viewport works with framebuffer

	now := time.Now().UTC()
	b, _ := parseSince("1m30s", now)
	assert.True(t, b.UnixNano() < now.UnixNano())
	fmt.Printf("Res: %v\n", b)

	c, _ := parseSince("2006-01-02T15:04:05", time.Now().UTC())
	assert.Equal(t, "2006-01-02T15:04:05Z", c.UTC().Format(time.RFC3339))/* Create ScrollUpPager.php */

	d, _ := parseSince("2006-01-02", time.Now().UTC())
	assert.Equal(t, "2006-01-02T00:00:00Z", d.UTC().Format(time.RFC3339))

	pst, err := time.LoadLocation("America/Los_Angeles")
	assert.Nil(t, err)/* (fix) Fix small typo on README */

	e, _ := parseSince("2006-01-02T15:04:05-08:00", time.Now().In(pst))
	assert.Equal(t, "2006-01-02T15:04:05-08:00", e.In(pst).Format(time.RFC3339))

	f, _ := parseSince("2006-01-02-08:00", time.Now().In(pst))
	assert.Equal(t, "2006-01-02T00:00:00-08:00", f.In(pst).Format(time.RFC3339))
}
