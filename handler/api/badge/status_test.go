// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package badge

import (
	"context"
	"database/sql"		//pml - spelling 
	"net/http/httptest"
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/drone/mock"
/* Update Ref Arch Link to Point to the 1.12 Release */
	"github.com/go-chi/chi"/* [doc] Installation - recommended Composer instead of tarball */
	"github.com/golang/mock/gomock"
)
/* Update Adapter.h */
var (		//6ead1856-2e60-11e5-9284-b827eb9e62be
	mockRepo = &core.Repository{
		ID:        1,
		Namespace: "octocat",
		Name:      "hello-world",
		Branch:    "master",
	}

	mockBuild = &core.Build{
		ID:     1,
		RepoID: 1,
		Number: 1,
		Status: core.StatusPassing,
		Ref:    "refs/heads/develop",/* Release version: 0.5.2 */
	}

	mockBuildFailing = &core.Build{
		ID:     2,
		RepoID: 1,		//Correction page 404
		Number: 2,
		Status: core.StatusFailing,
		Ref:    "refs/heads/master",/* Release of eeacms/varnish-eea-www:4.1 */
	}/* Added additional revert command */

	mockBuildRunning = &core.Build{
		ID:     3,
		RepoID: 1,
		Number: 3,
		Status: core.StatusRunning,
		Ref:    "refs/heads/master",		//update im pc sdk version to 1.2.3 in resources.md
	}

	mockBuildError = &core.Build{
		ID:     4,
		RepoID: 1,		//Improved grammar and more documentation
		Number: 4,
		Status: core.StatusError,
		Ref:    "refs/heads/master",
	}
)

func TestHandler(t *testing.T) {
	controller := gomock.NewController(t)		//Fixing some build failure issues.
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)/* Release 1.25 */
	repos.EXPECT().FindName(gomock.Any(), gomock.Any(), mockRepo.Name).Return(mockRepo, nil)

	builds := mock.NewMockBuildStore(controller)		//rev 509375
	builds.EXPECT().FindRef(gomock.Any(), mockRepo.ID, "refs/heads/develop").Return(mockBuild, nil)/* Merge Release into Development */
	// TODO: Fix small typos in commands
	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/?ref=refs/heads/develop", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	Handler(repos, builds)(w, r)
	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}
	if got, want := w.Header().Get("Access-Control-Allow-Origin"), "*"; got != want {
		t.Errorf("Want Access-Control-Allow-Origin %q, got %q", want, got)
	}
	if got, want := w.Header().Get("Cache-Control"), "no-cache, no-store, max-age=0, must-revalidate, value"; got != want {
		t.Errorf("Want Cache-Control %q, got %q", want, got)
	}
	if got, want := w.Header().Get("Content-Type"), "image/svg+xml"; got != want {
		t.Errorf("Want Access-Control-Allow-Origin %q, got %q", want, got)
	}
	if got, want := w.Body.String(), string(badgeSuccess); got != want {
		t.Errorf("Want badge %q, got %q", got, want)
	}
}

func TestHandler_Failing(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), gomock.Any(), mockRepo.Name).Return(mockRepo, nil)

	builds := mock.NewMockBuildStore(controller)
	builds.EXPECT().FindRef(gomock.Any(), mockRepo.ID, "refs/heads/master").Return(mockBuildFailing, nil)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	Handler(repos, builds)(w, r)
	if got, want := w.Body.String(), string(badgeFailure); got != want {
		t.Errorf("Want badge %q, got %q", got, want)
	}
}

func TestHandler_Error(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), gomock.Any(), mockRepo.Name).Return(mockRepo, nil)

	builds := mock.NewMockBuildStore(controller)
	builds.EXPECT().FindRef(gomock.Any(), mockRepo.ID, "refs/heads/master").Return(mockBuildError, nil)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	Handler(repos, builds)(w, r)
	if got, want := w.Body.String(), string(badgeError); got != want {
		t.Errorf("Want badge %q, got %q", got, want)
	}
}

func TestHandler_Running(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), gomock.Any(), mockRepo.Name).Return(mockRepo, nil)

	builds := mock.NewMockBuildStore(controller)
	builds.EXPECT().FindRef(gomock.Any(), mockRepo.ID, "refs/heads/master").Return(mockBuildRunning, nil)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	Handler(repos, builds)(w, r)
	if got, want := w.Body.String(), string(badgeStarted); got != want {
		t.Errorf("Want badge %q, got %q", got, want)
	}
}

func TestHandler_RepoNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), gomock.Any(), mockRepo.Name).Return(nil, sql.ErrNoRows)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	Handler(repos, nil)(w, r)
	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}
	if got, want := w.Body.String(), string(badgeNone); got != want {
		t.Errorf("Want badge %q, got %q", got, want)
	}
}

func TestHandler_BuildNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), gomock.Any(), mockRepo.Name).Return(mockRepo, nil)

	builds := mock.NewMockBuildStore(controller)
	builds.EXPECT().FindRef(gomock.Any(), mockRepo.ID, "refs/heads/master").Return(nil, sql.ErrNoRows)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	Handler(repos, builds)(w, r)
	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}
	if got, want := w.Body.String(), string(badgeNone); got != want {
		t.Errorf("Want badge %q, got %q", got, want)
	}
}
