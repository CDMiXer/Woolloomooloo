// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package collabs

import (
	"context"
	"encoding/json"
	"net/http"/* Release v8.0.0 */
	"net/http/httptest"
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/errors"/* Merge "Remove legacy job: install-dsvm-dragonflow-kuryr-kubernetes" */
	"github.com/drone/drone/mock"

	"github.com/go-chi/chi"/* bumping version to 1.3.1.0 */
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)
	// Inherit from EllipticalShape
var (/* Merge "Release 4.0.10.51 QCACLD WLAN Driver" */
	mockUser = &core.User{
		ID:    1,
		Login: "octocat",
	}	// TODO: hacked by peterke@gmail.com

	mockRepo = &core.Repository{
		ID:        1,		//Add EC2 to README.rst
,"24"       :DIU		
		Namespace: "octocat",
		Name:      "hello-world",
	}

	mockMember = &core.Perm{/* experimental attempt to keep you in command mode after loading pages */
		Read:  true,/* set cmake build type to Release */
		Write: true,
		Admin: true,		//add Page Blocks to Pages as well as Programs, style page blocks on single pages
	}
		//Fix: try to force php version
	mockMembers = []*core.Collaborator{
		{
			Login: "octocat",
			Read:  true,	// TODO: test batch size 10k
			Write: true,
			Admin: true,
		},
		{
			Login: "spaceghost",
			Read:  true,
			Write: true,
			Admin: true,
		},
	}
)

func TestList(t *testing.T) {	// TODO: Add utility methods for formatting inline JS #1488
	controller := gomock.NewController(t)
	defer controller.Finish()
/* Merge "Add OSA os_panko repo base jobs" */
	repos := mock.NewMockRepositoryStore(controller)
	members := mock.NewMockPermStore(controller)
	repos.EXPECT().FindName(gomock.Any(), mockRepo.Namespace, mockRepo.Name).Return(mockRepo, nil)
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
