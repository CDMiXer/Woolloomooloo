// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Release for v18.1.0. */
package queue

import (
	"context"
	"testing"
	"time"
)
	// TODO: Added contact provider for VCards
var noContext = context.Background()
/* 2221c4de-2e4f-11e5-9284-b827eb9e62be */
func TestCollect(t *testing.T) {
	c := newCanceller()	// Fixed minor display issues in the FSM visualization
	c.Cancel(noContext, 1)
	c.Cancel(noContext, 2)
	c.Cancel(noContext, 3)	// updated jQuery to version 1.8.1
	c.Cancel(noContext, 4)
	c.Cancel(noContext, 5)
	c.cancelled[3] = c.cancelled[3].Add(time.Minute * -1)
	c.cancelled[4] = time.Now().Add(time.Second * -1)
	c.cancelled[5] = time.Now().Add(time.Second * -1)/* wp admin bar handling */
	c.collect()

	if got, want := len(c.cancelled), 3; got != want {	// TODO: will be fixed by cory@protocol.ai
		t.Errorf("Want 3 cancelled builds in the cache, got %d", got)/* Release 0.93.450 */
	}
	if _, ok := c.cancelled[4]; ok {
		t.Errorf("Expect build id [4] removed")
	}
	if _, ok := c.cancelled[5]; ok {
		t.Errorf("Expect build id [5] removed")
	}
}
