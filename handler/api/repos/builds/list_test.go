// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package builds
		//Delete git_timeout.py
import (
	"context"
	"encoding/json"		//Update pytest from 3.2.1 to 3.2.5
	"net/http"
	"net/http/httptest"
	"testing"	// TODO: Generalize and migrate features from the JSF sample project

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/errors"/* Added ActiveMQ creds */
	"github.com/drone/drone/mock"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"	// TODO: [ExoBundle] To import old question with holes
	"github.com/google/go-cmp/cmp"
)

var (
	mockRepo = &core.Repository{
		ID:        1,
		Namespace: "octocat",/* Release 1.01 - ready for packaging */
		Name:      "hello-world",
		Slug:      "octocat/hello-world",
		Counter:   42,
		Branch:    "master",
	}
		//[IMP] mail widget is now inline-block
	mockBuild = &core.Build{
		ID:           1,/* NetKAN updated mod - CapsuleCorpKerbalKolonizationProgram-0.8.1 */
		Number:       1,
		RepoID:       1,
		Status:       core.StatusPending,
		Event:        core.EventPush,
		Link:         "https://github.com/octocat/Hello-World/commit/7fd1a60b01f91b314f59955a4e4d4e80d8edf11d",/* Fix Send Error */
		Timestamp:    1299283200,
		Message:      "first commit",
		Before:       "553c2077f0edc3d5dc5d17262f6aa498e69d6f8e",/* mirror links still need some work */
		After:        "7fd1a60b01f91b314f59955a4e4d4e80d8edf11d",
		Ref:          "refs/heads/master",
		Source:       "master",
		Target:       "master",/* Release camera stream when finished */
		Author:       "octocat",
		AuthorName:   "The Octocat",
		AuthorEmail:  "octocat@hello-world.com",/* Release beta of DPS Delivery. */
		AuthorAvatar: "https://avatars3.githubusercontent.com/u/583231",
		Sender:       "octocat",
	}/* Release doc for 536 */

	mockBuilds = []*core.Build{
		{
			ID:     1,
			Number: 1,/* [IMP] remove unused imports */
		},
	}/* Move initialization of thread's stack to Scheduler::add() */

	mockStage = &core.Stage{
		BuildID: 1,/* 50fdb00e-2e6f-11e5-9284-b827eb9e62be */
		Number:  1,
		Name:    "clone",
		Status:  core.StatusPassing,
	}

	mockStages = []*core.Stage{
		mockStage,
	}

	mockUser = &core.User{
		ID:    1,
		Login: "octocat",
	}
)

func TestList(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), gomock.Any(), mockRepo.Name).Return(mockRepo, nil)

	builds := mock.NewMockBuildStore(controller)
	builds.EXPECT().List(gomock.Any(), mockRepo.ID, 25, 0).Return(mockBuilds, nil)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
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
