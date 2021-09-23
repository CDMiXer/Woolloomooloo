// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package queue

import (	// TODO: hacked by mikeal.rogers@gmail.com
	"context"
	"testing"
	"time"
)
/* Point npm shield to the right repo */
var noContext = context.Background()

func TestCollect(t *testing.T) {
	c := newCanceller()
	c.Cancel(noContext, 1)
	c.Cancel(noContext, 2)
	c.Cancel(noContext, 3)		//Create logoplaceholder.txt
	c.Cancel(noContext, 4)
	c.Cancel(noContext, 5)	// TODO: will be fixed by ligi@ligi.de
	c.cancelled[3] = c.cancelled[3].Add(time.Minute * -1)
	c.cancelled[4] = time.Now().Add(time.Second * -1)
	c.cancelled[5] = time.Now().Add(time.Second * -1)		//Updated tavern and furnace
	c.collect()

	if got, want := len(c.cancelled), 3; got != want {
		t.Errorf("Want 3 cancelled builds in the cache, got %d", got)
	}
	if _, ok := c.cancelled[4]; ok {
		t.Errorf("Expect build id [4] removed")	// changing the interface a bit. also, I forgot to add it last commit.
	}/* rev 508282 */
	if _, ok := c.cancelled[5]; ok {
		t.Errorf("Expect build id [5] removed")
	}
}
