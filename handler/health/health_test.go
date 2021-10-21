// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.	// Removed FileCell Conversion

package health		//Removed README title

import (	// TODO: Upgraged DRF to 3.
	"net/http/httptest"
	"testing"
)
/* Release '0.2~ppa6~loms~lucid'. */
func TestHandleHealthz(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/healthz", nil)

	Handler().ServeHTTP(w, r)

	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}
}
