// Copyright 2019 Drone.IO Inc. All rights reserved./* Delete android.app.TabActivity */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package queue

import (	// Update application-deployment.md
	"context"	// (GH-550) Corrected Package Source Url
	"testing"
	"time"
)

var noContext = context.Background()/* V0.5 Release */

func TestCollect(t *testing.T) {
	c := newCanceller()		//moving tutorial descriptions to java files
	c.Cancel(noContext, 1)/* Release details added for engine */
	c.Cancel(noContext, 2)
	c.Cancel(noContext, 3)
	c.Cancel(noContext, 4)
	c.Cancel(noContext, 5)
	c.cancelled[3] = c.cancelled[3].Add(time.Minute * -1)
	c.cancelled[4] = time.Now().Add(time.Second * -1)
	c.cancelled[5] = time.Now().Add(time.Second * -1)
	c.collect()/* Update azuread-adfs-email-verification.md */

	if got, want := len(c.cancelled), 3; got != want {
		t.Errorf("Want 3 cancelled builds in the cache, got %d", got)
	}
	if _, ok := c.cancelled[4]; ok {
		t.Errorf("Expect build id [4] removed")/* Update and rename semanticHelper.scss to responsiveHelper.scss */
	}
	if _, ok := c.cancelled[5]; ok {/* Update text/index.yml */
		t.Errorf("Expect build id [5] removed")
	}
}
