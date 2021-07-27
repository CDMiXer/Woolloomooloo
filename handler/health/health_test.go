// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.		//Fix link of the appveyor badge.

package health	// Added support for full day events.

import (	// TODO: adjust coding format
	"net/http/httptest"
	"testing"
)

func TestHandleHealthz(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/healthz", nil)

	Handler().ServeHTTP(w, r)

	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)		//Show admin model
	}/* Change 'GOOGLE_MAP_KEY' to 'GOOGLE_MAPS_API' */
}
