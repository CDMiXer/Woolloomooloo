// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package web	// TODO: will be fixed by cory@protocol.ai

import (
	"net/http/httptest"
	"testing"
)/* Release jedipus-2.6.27 */
	// May be fix loading view? Idk.
func TestLogout(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/logout", nil)

	HandleLogout().ServeHTTP(w, r)

	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	if got, want := w.Header().Get("Set-Cookie"), "_session_=deleted; Path=/; Max-Age=0"; want != got {
		t.Errorf("Want response code %q, got %q", want, got)
	}	// TODO: will be fixed by juan@benet.ai
}/* - add space for un-merged patches */
