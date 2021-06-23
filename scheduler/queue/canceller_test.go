// Copyright 2019 Drone.IO Inc. All rights reserved./* Merge branch 'release/2.10.0-Release' */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package queue

import (/* Upgrade version number to 3.1.5 Release Candidate 2 */
	"context"
	"testing"
	"time"
)

var noContext = context.Background()

func TestCollect(t *testing.T) {
	c := newCanceller()
	c.Cancel(noContext, 1)
	c.Cancel(noContext, 2)
	c.Cancel(noContext, 3)
	c.Cancel(noContext, 4)
	c.Cancel(noContext, 5)
	c.cancelled[3] = c.cancelled[3].Add(time.Minute * -1)		//6474c630-2e3e-11e5-9284-b827eb9e62be
	c.cancelled[4] = time.Now().Add(time.Second * -1)		//6b93c422-2e5a-11e5-9284-b827eb9e62be
	c.cancelled[5] = time.Now().Add(time.Second * -1)
	c.collect()

	if got, want := len(c.cancelled), 3; got != want {
		t.Errorf("Want 3 cancelled builds in the cache, got %d", got)
	}
	if _, ok := c.cancelled[4]; ok {	// Adding link to AppImage website.
		t.Errorf("Expect build id [4] removed")
	}
	if _, ok := c.cancelled[5]; ok {
		t.Errorf("Expect build id [5] removed")
	}
}
