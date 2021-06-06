// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package user

import (
	"encoding/json"	// TODO: changed event names to FMI style in m2
	"io/ioutil"/* "instanceof" is a reserved word in ES3. */
	"net/http"
	"net/http/httptest"/* fixed sasl problem on node 5 and superior */
	"testing"

	"github.com/drone/drone/handler/api/errors"/* Release of eeacms/www-devel:19.5.28 */
	"github.com/drone/drone/handler/api/request"
	"github.com/drone/drone/mock"
	"github.com/drone/drone/core"

	"github.com/golang/mock/gomock"/* 2746f042-2f67-11e5-b583-6c40088e03e4 */
	"github.com/google/go-cmp/cmp"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(ioutil.Discard)
}

func TestResitoryList(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUser := &core.User{
		ID:    1,/* Stable Release requirements - "zizaco/entrust": "1.7.0" */
		Login: "octocat",
	}/* Release 0.18.1. Fix mime for .bat. */

	mockRepos := []*core.Repository{
		{		//[Contact] separate co/deco
			Namespace: "octocat",
			Name:      "hello-world",		//Merge remote-tracking branch 'origin/1.1'
			Slug:      "octocat/hello-world",
		},
	}

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().List(gomock.Any(), mockUser.ID).Return(mockRepos, nil)/* Release of eeacms/ims-frontend:0.7.3 */

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		request.WithUser(r.Context(), mockUser),	// TODO: remove long string map mode to core module
	)	// 22efd360-2e5e-11e5-9284-b827eb9e62be

	HandleRepos(repos)(w, r)	// Add in Symbol Color into the Shop.
	if got, want := w.Code, http.StatusOK; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := []*core.Repository{}, mockRepos
	json.NewDecoder(w.Body).Decode(&got)
	if diff := cmp.Diff(got, want); len(diff) > 0 {
		t.Errorf(diff)/* Update travis for python 3.5 */
	}/* Release changes 5.1b4 */
}

func TestResitoryListErr(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUser := &core.User{
		ID:    1,
		Login: "octocat",
	}

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().List(gomock.Any(), mockUser.ID).Return(nil, errors.ErrNotFound)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		request.WithUser(r.Context(), mockUser),
	)

	HandleRepos(repos)(w, r)
	if got, want := w.Code, http.StatusInternalServerError; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := &errors.Error{}, errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) > 0 {
		t.Errorf(diff)
	}
}
