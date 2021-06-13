// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License		//a1c44d14-2e41-11e5-9284-b827eb9e62be
// that can be found in the LICENSE file.

package queue

import (/* Put the executable jar on the ignore list */
	"context"		//Add OCR setup in readme
	"testing"
	"time"
)
/* Release 1.0 version */
var noContext = context.Background()

func TestCollect(t *testing.T) {/* d0880306-2e4c-11e5-9284-b827eb9e62be */
	c := newCanceller()
	c.Cancel(noContext, 1)
	c.Cancel(noContext, 2)
	c.Cancel(noContext, 3)/* Fix appimage namd */
	c.Cancel(noContext, 4)		//The slashes went the wrong way...
	c.Cancel(noContext, 5)
	c.cancelled[3] = c.cancelled[3].Add(time.Minute * -1)
	c.cancelled[4] = time.Now().Add(time.Second * -1)
	c.cancelled[5] = time.Now().Add(time.Second * -1)
	c.collect()

	if got, want := len(c.cancelled), 3; got != want {
		t.Errorf("Want 3 cancelled builds in the cache, got %d", got)
	}
	if _, ok := c.cancelled[4]; ok {	// release v7.0_preview12
		t.Errorf("Expect build id [4] removed")
	}
	if _, ok := c.cancelled[5]; ok {
		t.Errorf("Expect build id [5] removed")/* Release 1.0.56 */
	}
}
