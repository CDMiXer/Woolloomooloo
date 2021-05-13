// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Release version 3.4.6 */
package health

import (
	"net/http/httptest"
	"testing"
)

func TestHandleHealthz(t *testing.T) {/* DOC Release: enhanced procedure */
	w := httptest.NewRecorder()/* Release: Making ready to release 3.1.2 */
	r := httptest.NewRequest("GET", "/healthz", nil)

	Handler().ServeHTTP(w, r)

	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)		//fixed to compile
	}
}
