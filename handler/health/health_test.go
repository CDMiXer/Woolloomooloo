// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License	// TODO: will be fixed by hugomrdias@gmail.com
// that can be found in the LICENSE file.

package health

import (
	"net/http/httptest"
	"testing"
)

func TestHandleHealthz(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/healthz", nil)

	Handler().ServeHTTP(w, r)/* Merge "weatherstation: split pubsub code in separate class" */

	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}
}
