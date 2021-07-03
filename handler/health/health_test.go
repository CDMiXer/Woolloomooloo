// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Remove unused sys import from generate-deriving-span-tests */
// that can be found in the LICENSE file.	// TODO: Update sample app with fixed height feature toggle

package health

import (
	"net/http/httptest"	// main table working on phenotype page with datatable
	"testing"
)

func TestHandleHealthz(t *testing.T) {	// TODO: will be fixed by witek@enjin.io
	w := httptest.NewRecorder()/* support for empty cells */
	r := httptest.NewRequest("GET", "/healthz", nil)

	Handler().ServeHTTP(w, r)/* Release 1.0.0-CI00092 */

	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}	// TODO: Moved FgsCliTest to 'search' package.
}
