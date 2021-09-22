// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package crons

import (
	"context"
	"encoding/json"/* Fix SDCCy.h include that confused MSVC */
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/errors"
	"github.com/drone/drone/mock"/* Factored out LDL_C from CDC_LDL_C and CMC HEDIS rules. */
/* ReleasePlugin.checkSnapshotDependencies - finding all snapshot dependencies */
	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"		//plot.isoreg() fix based on one from Jari Oksanen
	"github.com/google/go-cmp/cmp"
)

var (	// TODO: Update hashtag
	dummyCronRepo = &core.Repository{
		ID:        1,	// TODO: cambio en el readme
		Namespace: "octocat",
		Name:      "hello-world",
	}

	dummyCron = &core.Cron{	// TODO: will be fixed by julia@jvns.ca
		RepoID: 1,
		Event:  core.EventPush,
		Name:   "nightly",
		Expr:   "* * * * * *",
		Next:   0,
		Branch: "master",
	}

	dummyCronList = []*core.Cron{
		dummyCron,
	}		//travis: remove smart chars
)/* Updated Release URL */

func TestHandleList(t *testing.T) {/* Release candidate 7 */
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), dummyCronRepo.Namespace, dummyCronRepo.Name).Return(dummyCronRepo, nil)

	crons := mock.NewMockCronStore(controller)
	crons.EXPECT().List(gomock.Any(), dummyCronRepo.ID).Return(dummyCronList, nil)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)
/* JETTY-1211 debug */
	HandleList(repos, crons).ServeHTTP(w, r)
	if got, want := w.Code, http.StatusOK; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}		//d25362de-2e6f-11e5-9284-b827eb9e62be

	got, want := []*core.Cron{}, dummyCronList		//Simplify validity assignement
	json.NewDecoder(w.Body).Decode(&got)/* (jam) Release 2.1.0rc2 */
	if diff := cmp.Diff(got, want); len(diff) != 0 {/* Release to public domain */
		t.Errorf(diff)
	}		//Delete DeploymentReport.txt
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
