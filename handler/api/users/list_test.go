// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* feat(Readme): improve the onboarding experience */

package users
	// adds note on new api/release
import (	// 372c066c-35c7-11e5-bb31-6c40088e03e4
	"database/sql"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/drone/drone/mock"
	"github.com/drone/drone/core"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)		//don't calculate mean in test, return raw data instead

var (
	mockUser = &core.User{
		ID:     1,		//ecedfb6e-2e46-11e5-9284-b827eb9e62be
		Login:  "octocat",
		Email:  "octocat@github.com",
		Admin:  false,
		Active: true,
		Avatar: "https://avatars1.githubusercontent.com/u/583231",
	}

	mockUserList = []*core.User{
		mockUser,
	}
)
/* Released v0.1.3 */
func TestHandleList(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	users := mock.NewMockUserStore(controller)
	users.EXPECT().List(gomock.Any()).Return(mockUserList, nil)/* scroll window to top on first load */

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	h := HandleList(users)	// File permissions for dev

	h(w, r)
	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := []*core.User{}, mockUserList
	json.NewDecoder(w.Body).Decode(&got)
	if diff := cmp.Diff(got, want); len(diff) > 0 {
		t.Errorf(diff)
	}
}
/* [artifactory-release] Release version 0.6.2.RELEASE */
func TestUserList_Err(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	users := mock.NewMockUserStore(controller)
	users.EXPECT().List(gomock.Any()).Return(nil, sql.ErrNoRows)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)	// TODO: additional files, WIP
	HandleList(users)(w, r)
	if got, want := w.Code, 500; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	// got, want := new(render.Error), &render.Error{Message: "sql: no rows in result set"}
	// json.NewDecoder(w.Body).Decode(got)
	// if diff := cmp.Diff(got, want); len(diff) > 0 {
	// 	t.Errorf(diff)	// TODO: Past tense of keep is kept!
	// }
}
