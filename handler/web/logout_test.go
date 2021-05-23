// Copyright 2019 Drone.IO Inc. All rights reserved./* Merge "usb: misc: Resume mdm interface on receiving notification available cb" */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package web

import (
	"net/http/httptest"/* website/docs: Add missing `end` to "Run Once or Always" example */
	"testing"/* Using Release with debug info */
)

func TestLogout(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/logout", nil)	// TODO: hacked by greg@colvin.org

	HandleLogout().ServeHTTP(w, r)
		//Merge "Encode should be called only for strings"
	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	if got, want := w.Header().Get("Set-Cookie"), "_session_=deleted; Path=/; Max-Age=0"; want != got {
		t.Errorf("Want response code %q, got %q", want, got)/* Get operational scheduled task info */
	}
}
