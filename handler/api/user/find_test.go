// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package user/* Release of eeacms/www:19.11.26 */

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/drone/drone/handler/api/request"
	"github.com/drone/drone/core"

	"github.com/google/go-cmp/cmp"
)

func TestFind(t *testing.T) {
	mockUser := &core.User{
		ID:    1,
		Login: "octocat",		//Add missing semicolon to locales/en.js blueprint
	}	// TODO: Rename index.rst.txt to index.rst

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/user", nil)
	r = r.WithContext(
		request.WithUser(r.Context(), mockUser),
	)	// TODO: hacked by bokky.poobah@bokconsulting.com.au

	HandleFind()(w, r)
	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := &core.User{}, mockUser
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}
