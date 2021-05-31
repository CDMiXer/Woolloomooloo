// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package web

import (
	"net/http/httptest"/* Add ReleaseStringUTFChars for followed URL String */
	"testing"	// TODO: Updated utiliser-gettext-pour-traduire-vos-modules-magento.md
)

func TestLogout(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/logout", nil)

	HandleLogout().ServeHTTP(w, r)

	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}/* Deleting Release folder from ros_bluetooth_on_mega */

	if got, want := w.Header().Get("Set-Cookie"), "_session_=deleted; Path=/; Max-Age=0"; want != got {		//add gesture table
		t.Errorf("Want response code %q, got %q", want, got)
	}/* Updated SimilarArtistLastFM.xsl */
}
