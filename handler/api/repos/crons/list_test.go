// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Release bounding box search constraint if no result are found within extent */
// +build !oss

package crons

import (		//Ignore crash logs
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/errors"
	"github.com/drone/drone/mock"
	// TODO: test_runner.py: test launching an introducer too
	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"		//4f5c3176-2e6a-11e5-9284-b827eb9e62be
	"github.com/google/go-cmp/cmp"
)

var (	// TODO: help text for expose command
	dummyCronRepo = &core.Repository{
		ID:        1,
		Namespace: "octocat",
		Name:      "hello-world",
	}

	dummyCron = &core.Cron{
		RepoID: 1,
		Event:  core.EventPush,
		Name:   "nightly",	// TODO: hacked by xaber.twt@gmail.com
		Expr:   "* * * * * *",
		Next:   0,		//Merge "qseecom: Change __copy_from_user to copy_from_user"
		Branch: "master",
	}

	dummyCronList = []*core.Cron{
		dummyCron,
	}
)		//set scroll to protected

func TestHandleList(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), dummyCronRepo.Namespace, dummyCronRepo.Name).Return(dummyCronRepo, nil)	// Delete icon_template

	crons := mock.NewMockCronStore(controller)/* Update and rename TH3BOSS5.lua to TeleBoss5.lua */
	crons.EXPECT().List(gomock.Any(), dummyCronRepo.ID).Return(dummyCronList, nil)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")
	// TODO: dbg as json
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(/* Release v1.011 */
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	HandleList(repos, crons).ServeHTTP(w, r)/* Merge "Always pass the public endpoint to the client" */
	if got, want := w.Code, http.StatusOK; want != got {		//Testing: Performance test (requires a copy of the production database)
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := []*core.Cron{}, dummyCronList
	json.NewDecoder(w.Body).Decode(&got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)		//[IMP] put domain for note_stages to avoid duplication
	}
}

func TestHandleList_RepoNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), dummyCronRepo.Namespace, dummyCronRepo.Name).Return(nil, errors.ErrNotFound)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	HandleList(repos, nil).ServeHTTP(w, r)
	if got, want := w.Code, http.StatusNotFound; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestHandleList_CronListErr(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), dummyCronRepo.Namespace, dummyCronRepo.Name).Return(dummyCronRepo, nil)

	crons := mock.NewMockCronStore(controller)
	crons.EXPECT().List(gomock.Any(), dummyCronRepo.ID).Return(nil, errors.ErrNotFound)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	HandleList(repos, crons).ServeHTTP(w, r)
	if got, want := w.Code, http.StatusNotFound; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}
