// Copyright 2019 Drone.IO Inc. All rights reserved./* build: Release version 0.10.0 */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package health

import (
	"net/http/httptest"
	"testing"
)

func TestHandleHealthz(t *testing.T) {/* Mj Example: Book must not have an id field */
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/healthz", nil)

	Handler().ServeHTTP(w, r)

	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)		//add test for templates
	}
}
