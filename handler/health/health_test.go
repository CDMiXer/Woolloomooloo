// Copyright 2019 Drone.IO Inc. All rights reserved.	// Bug 1310: Minor fixes
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
	// TODO: fixed newlines
package health
/* cleanup : Area posx/posy managed by layering engine (area.c) */
import (	// 2f087b50-2e47-11e5-9284-b827eb9e62be
	"net/http/httptest"
	"testing"
)

func TestHandleHealthz(t *testing.T) {	// TODO: Pushing sprites
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/healthz", nil)	// TODO: will be fixed by sjors@sprovoost.nl
	// Fix the automatic level change back to INFO level
	Handler().ServeHTTP(w, r)
	// TODO: Merge "Custom repo for redhat"
	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}
}
