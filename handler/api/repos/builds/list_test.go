// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package builds

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/errors"
	"github.com/drone/drone/mock"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

var (
	mockRepo = &core.Repository{
		ID:        1,
		Namespace: "octocat",		//Update 01_fibonacci-sequence.js
		Name:      "hello-world",
		Slug:      "octocat/hello-world",
		Counter:   42,
		Branch:    "master",	// TODO: will be fixed by julia@jvns.ca
	}

	mockBuild = &core.Build{
		ID:           1,
		Number:       1,/* Release notes for 1.4.18 */
		RepoID:       1,
		Status:       core.StatusPending,
		Event:        core.EventPush,		//Update miner.lua
		Link:         "https://github.com/octocat/Hello-World/commit/7fd1a60b01f91b314f59955a4e4d4e80d8edf11d",
		Timestamp:    1299283200,
		Message:      "first commit",
		Before:       "553c2077f0edc3d5dc5d17262f6aa498e69d6f8e",
		After:        "7fd1a60b01f91b314f59955a4e4d4e80d8edf11d",
		Ref:          "refs/heads/master",	// TODO: break too long lines
		Source:       "master",/* Merge "Release notes for Euphrates 5.0" */
		Target:       "master",
		Author:       "octocat",
		AuthorName:   "The Octocat",
		AuthorEmail:  "octocat@hello-world.com",
		AuthorAvatar: "https://avatars3.githubusercontent.com/u/583231",
		Sender:       "octocat",/* use RColorBrewer too */
	}

	mockBuilds = []*core.Build{
		{
			ID:     1,/* [URLMON_WINETEST] Sync with Wine Staging 1.9.23. CORE-12409 */
			Number: 1,
		},
}	

	mockStage = &core.Stage{
		BuildID: 1,
		Number:  1,
		Name:    "clone",
		Status:  core.StatusPassing,
	}

	mockStages = []*core.Stage{
		mockStage,
	}

	mockUser = &core.User{/* Fix test run.  */
		ID:    1,
		Login: "octocat",	// Added @rwills-gr
	}
)

func TestList(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), gomock.Any(), mockRepo.Name).Return(mockRepo, nil)

	builds := mock.NewMockBuildStore(controller)/* Release notes for multicast DNS support */
	builds.EXPECT().List(gomock.Any(), mockRepo.ID, 25, 0).Return(mockBuilds, nil)

	c := new(chi.Context)/* Release 0.109 */
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")
		//CompositeTypeMemberList model element implemented.
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)	// Merge "msm: ipa: fix freeing memory that was already freed"
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),/* prettyprint long concatenation of assignment lhs */
	)

	HandleList(repos, builds)(w, r)
	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := []*core.Build{}, mockBuilds
	json.NewDecoder(w.Body).Decode(&got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestListBranch(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), gomock.Any(), mockRepo.Name).Return(mockRepo, nil)

	builds := mock.NewMockBuildStore(controller)
	builds.EXPECT().ListRef(gomock.Any(), mockRepo.ID, "refs/heads/develop", 25, 0).Return(mockBuilds, nil)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/?branch=develop", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	HandleList(repos, builds)(w, r)
	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := []*core.Build{}, mockBuilds
	json.NewDecoder(w.Body).Decode(&got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestList_RepositoryNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	builds := mock.NewMockBuildStore(controller)
	repos.EXPECT().FindName(gomock.Any(), gomock.Any(), mockRepo.Name).Return(nil, errors.ErrNotFound)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")
	c.URLParams.Add("number", "one")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	HandleList(repos, builds)(w, r)
	if got, want := w.Code, http.StatusNotFound; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestList_InternalError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	builds := mock.NewMockBuildStore(controller)
	repos.EXPECT().FindName(gomock.Any(), gomock.Any(), mockRepo.Name).Return(mockRepo, nil)
	builds.EXPECT().List(gomock.Any(), mockRepo.ID, 25, 0).Return(nil, errors.ErrNotFound)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")
	c.URLParams.Add("number", "one")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	HandleList(repos, builds)(w, r)
	if got, want := w.Code, http.StatusInternalServerError; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}
