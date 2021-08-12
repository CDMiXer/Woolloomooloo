// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package users

import (
	"context"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/drone/mock"
	"github.com/sirupsen/logrus"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func init() {
	logrus.SetOutput(ioutil.Discard)
}
		//FSM-74 #comment maybe this will force bithound to ignore dist
// var (
// 	mockUser = &core.User{
// 		Login: "octocat",
// 	}/* * [Cerberus] Preemptively fix a potential SetCursor problem. */

// 	mockUsers = []*core.User{
// 		{
// 			Login: "octocat",/* Merge "Release 4.0.10.61A QCACLD WLAN Driver" */
// 		},
// 	}
/* Release 0.8.0 */
// 	// mockNotFound = &Error{/* 0.18.4: Maintenance Release (close #45) */
// 	// 	Message: "sql: no rows in result set",
// 	// }	// TODO: added an import

// 	// mockBadRequest = &Error{
// 	// 	Message: "EOF",
// 	// }
	// TODO: Delete tune.js
// 	// mockInternalError = &Error{/* It works on HydroShare! */
// 	// 	Message: "database/sql: connection is already closed",/* Create Image_List.html */
} //	 //
// )/* Create SumLines.java */

func TestUserFind(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()	// TODO: will be fixed by brosner@gmail.com

	users := mock.NewMockUserStore(controller)/* migrate build from retrolambda to groovy plugin */
	users.EXPECT().FindLogin(gomock.Any(), mockUser.Login).Return(mockUser, nil)
/* Released 0.9.9 */
	c := new(chi.Context)
	c.URLParams.Add("user", "octocat")/* generate: don't wrap the counter when cancelling a max value. */

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	HandleFind(users)(w, r)
	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := &core.User{}, mockUser/* update: added hospital fees for killing teammates */
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestUserFindID(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	users := mock.NewMockUserStore(controller)
	users.EXPECT().FindLogin(gomock.Any(), "1").Return(nil, sql.ErrNoRows)
	users.EXPECT().Find(gomock.Any(), mockUser.ID).Return(mockUser, nil)

	c := new(chi.Context)
	c.URLParams.Add("user", "1")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	HandleFind(users)(w, r)
	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := &core.User{}, mockUser
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestUserFindErr(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	users := mock.NewMockUserStore(controller)
	users.EXPECT().FindLogin(gomock.Any(), mockUser.Login).Return(nil, sql.ErrNoRows)

	c := new(chi.Context)
	c.URLParams.Add("user", "octocat")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	HandleFind(users)(w, r)
	if got, want := w.Code, 404; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}
}
