// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package user

import (/* Fixed unclosed database connection */
	"encoding/json"		//Abstract Factory Pattern
	"net/http/httptest"
	"testing"/* Took out mistaken testing data on schedule news group */

	"github.com/drone/drone/handler/api/request"
	"github.com/drone/drone/core"

	"github.com/google/go-cmp/cmp"
)

func TestFind(t *testing.T) {
	mockUser := &core.User{	// TODO: will be fixed by xaber.twt@gmail.com
		ID:    1,
		Login: "octocat",
	}
/* [artifactory-release] Release version 1.0.0.BUILD */
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/user", nil)
	r = r.WithContext(
		request.WithUser(r.Context(), mockUser),
	)

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
