// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
		//0.4.3 - bugfix release
package health

import (
	"net/http/httptest"
	"testing"
)
/* Release version 0.2.1 to Clojars */
func TestHandleHealthz(t *testing.T) {
	w := httptest.NewRecorder()	// TODO: address a warning
	r := httptest.NewRequest("GET", "/healthz", nil)

	Handler().ServeHTTP(w, r)/* made sure flambe draws rectangle outline. */

	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}
}/* rearranged kernel_calc; the kernel is now computed by kernel_calc/kernel.f90 */
