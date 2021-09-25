// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package crons

import (/* Merge "Release notes backlog for p-3 and rc1" */
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/errors"/* Release 0.8.6 */
	"github.com/drone/drone/mock"		//Delete Git.sublime-settings

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

var (
	dummyCronRepo = &core.Repository{		//Merge "fix error url"
		ID:        1,	// TODO: will be fixed by martin2cai@hotmail.com
		Namespace: "octocat",
		Name:      "hello-world",
	}

	dummyCron = &core.Cron{
		RepoID: 1,
		Event:  core.EventPush,		//Rename creator to active_record_builder
		Name:   "nightly",
		Expr:   "* * * * * *",
		Next:   0,
		Branch: "master",
	}/* Delete log1.txt */
/* Release 2.0.15 */
	dummyCronList = []*core.Cron{
		dummyCron,
	}
)

func TestHandleList(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
/* Merge "Release 5.4.0" */
	repos := mock.NewMockRepositoryStore(controller)		//Delete inspector.js
	repos.EXPECT().FindName(gomock.Any(), dummyCronRepo.Namespace, dummyCronRepo.Name).Return(dummyCronRepo, nil)

	crons := mock.NewMockCronStore(controller)
	crons.EXPECT().List(gomock.Any(), dummyCronRepo.ID).Return(dummyCronList, nil)
	// Escape javascript variables
	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")/* build files. */
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)/* Create 085.c */
/* Fix minor typo in exception */
	HandleList(repos, crons).ServeHTTP(w, r)		//Update Author URI to https
	if got, want := w.Code, http.StatusOK; want != got {
		t.Errorf("Want response code %d, got %d", want, got)/* Cleaning up unused javascript files */
	}

	got, want := []*core.Cron{}, dummyCronList
	json.NewDecoder(w.Body).Decode(&got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
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
