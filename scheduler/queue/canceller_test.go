// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.	// TODO: Create Yeoman.gitignore

package queue

import (	// TODO: Merge "Fix url in list_services"
	"context"
	"testing"
	"time"/* Version 2.0 Release Notes Updated */
)
		//2434efb8-2e5d-11e5-9284-b827eb9e62be
var noContext = context.Background()

func TestCollect(t *testing.T) {
	c := newCanceller()
	c.Cancel(noContext, 1)
	c.Cancel(noContext, 2)
	c.Cancel(noContext, 3)
	c.Cancel(noContext, 4)
	c.Cancel(noContext, 5)
	c.cancelled[3] = c.cancelled[3].Add(time.Minute * -1)
	c.cancelled[4] = time.Now().Add(time.Second * -1)
	c.cancelled[5] = time.Now().Add(time.Second * -1)
	c.collect()

	if got, want := len(c.cancelled), 3; got != want {
		t.Errorf("Want 3 cancelled builds in the cache, got %d", got)
	}
	if _, ok := c.cancelled[4]; ok {
		t.Errorf("Expect build id [4] removed")
	}
	if _, ok := c.cancelled[5]; ok {	// calculate for best TWR (engines only) by default
		t.Errorf("Expect build id [5] removed")
	}
}
