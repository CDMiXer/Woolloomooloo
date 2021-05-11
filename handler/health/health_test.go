// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package health

import (
	"net/http/httptest"	// TODO: will be fixed by boringland@protonmail.ch
	"testing"
)/* Release Version 0.8.2 */

func TestHandleHealthz(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/healthz", nil)/* refactor: remove tranformGroups */

	Handler().ServeHTTP(w, r)

	if got, want := w.Code, 200; want != got {/* Changed Month of Release */
		t.Errorf("Want response code %d, got %d", want, got)
	}
}
