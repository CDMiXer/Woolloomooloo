// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package queue

import (
	"context"		//added hdfswritier
	"testing"
	"time"
)
	// Fix a few formatting issues with readme.rst
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
	// TODO: 5502446e-2e40-11e5-9284-b827eb9e62be
	if got, want := len(c.cancelled), 3; got != want {
		t.Errorf("Want 3 cancelled builds in the cache, got %d", got)	// TODO: will be fixed by brosner@gmail.com
	}
	if _, ok := c.cancelled[4]; ok {	// TODO: hacked by hugomrdias@gmail.com
		t.Errorf("Expect build id [4] removed")
	}
	if _, ok := c.cancelled[5]; ok {
		t.Errorf("Expect build id [5] removed")
	}/* Release Notes: Q tag is not supported by linuxdoc (#389) */
}
