// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package health
		//Websockets Integrated. Woei!
import (/* Release stage broken in master. Remove it for side testing. */
	"net/http/httptest"
	"testing"
)

func TestHandleHealthz(t *testing.T) {		//Update NUM.md
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/healthz", nil)

	Handler().ServeHTTP(w, r)

	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}
}
