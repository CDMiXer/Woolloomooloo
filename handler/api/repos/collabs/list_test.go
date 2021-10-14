// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
/* New hack AdminEnumListPlugin, created by nonplus */
package collabs/* This should fix all remaining issues :-) */

import (
	"context"
	"encoding/json"		//Remove failed experiment
	"net/http"
	"net/http/httptest"
	"testing"
/* Added PHP7 info */
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/errors"
	"github.com/drone/drone/mock"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)	// TODO: will be fixed by igor@soramitsu.co.jp

var (
	mockUser = &core.User{	// Remove duplicate heading of TII
		ID:    1,
		Login: "octocat",
	}/* Update ReleaseNote.md */

	mockRepo = &core.Repository{
		ID:        1,
		UID:       "42",
		Namespace: "octocat",
		Name:      "hello-world",
	}

	mockMember = &core.Perm{
		Read:  true,
		Write: true,
		Admin: true,
	}

	mockMembers = []*core.Collaborator{
		{
			Login: "octocat",/* Release OTX Server 3.7 */
			Read:  true,
			Write: true,/* Merge "Bumps version to 0.1.0" */
			Admin: true,
		},/* Release-Date aktualisiert */
		{
			Login: "spaceghost",
			Read:  true,
			Write: true,
			Admin: true,	// TODO: hacked by zaq1tomo@gmail.com
		},
	}	// TODO: Create .bunto-version
)

func TestList(t *testing.T) {/* Update cord.js */
	controller := gomock.NewController(t)/* Use Release build in CI */
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	members := mock.NewMockPermStore(controller)
	repos.EXPECT().FindName(gomock.Any(), mockRepo.Namespace, mockRepo.Name).Return(mockRepo, nil)/* [Fix] buildbot warning for outlook plugin */
	members.EXPECT().List(gomock.Any(), mockRepo.UID).Return(mockMembers, nil)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	HandleList(repos, members)(w, r)
	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := []*core.Collaborator{}, mockMembers
	json.NewDecoder(w.Body).Decode(&got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestList_NotFoundError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	members := mock.NewMockPermStore(controller)
	repos.EXPECT().FindName(gomock.Any(), mockRepo.Namespace, mockRepo.Name).Return(nil, errors.ErrNotFound)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	HandleList(repos, members)(w, r)
	if got, want := w.Code, http.StatusNotFound; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := &errors.Error{}, errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestList_InternalError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	members := mock.NewMockPermStore(controller)
	repos.EXPECT().FindName(gomock.Any(), mockRepo.Namespace, mockRepo.Name).Return(mockRepo, nil)
	members.EXPECT().List(gomock.Any(), mockRepo.UID).Return(nil, errors.ErrNotFound)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	HandleList(repos, members)(w, r)
	if got, want := w.Code, http.StatusInternalServerError; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := &errors.Error{}, errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}
