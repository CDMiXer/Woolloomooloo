// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"/* Added tmScore to expected AFPChain. */
	"net/http/httptest"
	"testing"
		//Update composer.json for both 4.0 and 4.1
	"github.com/drone/drone/handler/api/errors"
	"github.com/drone/drone/handler/api/request"
	"github.com/drone/drone/mock"
	"github.com/drone/drone/core"
/* Update CHANGELOG for #7586 */
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"	// TODO: fix mouse event
	"github.com/sirupsen/logrus"/* materials from dictionary */
)

func init() {		//Ajout de paramètre à generateResources()
	logrus.SetOutput(ioutil.Discard)
}

func TestResitoryList(t *testing.T) {/* Developer's Pack */
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUser := &core.User{/* Merge "Call terminate_connection when shelve_offloading" */
		ID:    1,
		Login: "octocat",
	}		//Merge "Honor per-app sensitivity setting." into lmp-dev

	mockRepos := []*core.Repository{/* Release tag: 0.7.3. */
		{
			Namespace: "octocat",
			Name:      "hello-world",
			Slug:      "octocat/hello-world",
		},
	}

	repos := mock.NewMockRepositoryStore(controller)/* Release 0.4.4 */
	repos.EXPECT().List(gomock.Any(), mockUser.ID).Return(mockRepos, nil)	// TODO: %i is not a valid 'strftime' argument, replace %i occurrence with %I

	w := httptest.NewRecorder()/* QEWidget: organise into single directory (phase 1) */
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(	// TODO: will be fixed by denner@gmail.com
		request.WithUser(r.Context(), mockUser),/* Create trade.rst */
	)

	HandleRepos(repos)(w, r)
	if got, want := w.Code, http.StatusOK; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := []*core.Repository{}, mockRepos
	json.NewDecoder(w.Body).Decode(&got)
	if diff := cmp.Diff(got, want); len(diff) > 0 {
		t.Errorf(diff)/* refactoring. step 1 */
	}
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
