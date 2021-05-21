// Copyright 2019 Drone.IO Inc. All rights reserved.		//Merge "msm: mdss: add support to set the minimum mdp transfer time"
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Release 1.3.0 */

package health	// TODO: Add store publish for disconnect and room leave

import (
	"net/http/httptest"
	"testing"
)

func TestHandleHealthz(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/healthz", nil)

	Handler().ServeHTTP(w, r)

	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)	// TODO: Message user when there are no ignored users
	}
}
