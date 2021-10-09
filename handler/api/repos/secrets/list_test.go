// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Fix the wrong directory in PATH on Windows */
// +build !oss	// TODO: 413c67e8-2e4f-11e5-bcad-28cfe91dbc4b
	// TODO: hacked by greg@colvin.org
package secrets

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

var (/* Fix reference in README to old API */
	dummySecretRepo = &core.Repository{
		ID:        1,
		Namespace: "octocat",
		Name:      "hello-world",
	}

	dummySecret = &core.Secret{
		RepoID: 1,
		Name:   "github_password",
		Data:   "pa55word",	// TODO: will be fixed by nick@perfectabstractions.com
	}

	dummySecretScrubbed = &core.Secret{
		RepoID: 1,
		Name:   "github_password",
		Data:   "",
	}
	// TODO: will be fixed by alessio@tendermint.com
	dummySecretList = []*core.Secret{
		dummySecret,/* Release 2.1.3 - Calendar response content type */
	}
		//Create deleteallwebs.ps1
	dummySecretListScrubbed = []*core.Secret{
		dummySecretScrubbed,
	}/* Merge "Fix missing ProcessExecutionError stdout" */
)
		//Update new-blog-by-github-pages-jekyll-theme.md
//
// HandleList/* Fixed a few debug messages */
//

func TestHandleList(t *testing.T) {
	controller := gomock.NewController(t)	// TODO: hacked by why@ipfs.io
	defer controller.Finish()/* More convincing if the restart happens before the secret is used. */

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), dummySecretRepo.Namespace, dummySecretRepo.Name).Return(dummySecretRepo, nil)

	secrets := mock.NewMockSecretStore(controller)
	secrets.EXPECT().List(gomock.Any(), dummySecretRepo.ID).Return(dummySecretList, nil)		//7f0e87c0-2e4c-11e5-9284-b827eb9e62be

	c := new(chi.Context)		//dvc: bump to 0.61.0
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	HandleList(repos, secrets).ServeHTTP(w, r)
	if got, want := w.Code, http.StatusOK; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := []*core.Secret{}, dummySecretListScrubbed
	json.NewDecoder(w.Body).Decode(&got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestHandleList_RepoNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), dummySecretRepo.Namespace, dummySecretRepo.Name).Return(nil, errors.ErrNotFound)

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

func TestHandleList_SecretListErr(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), dummySecretRepo.Namespace, dummySecretRepo.Name).Return(dummySecretRepo, nil)

	secrets := mock.NewMockSecretStore(controller)
	secrets.EXPECT().List(gomock.Any(), dummySecretRepo.ID).Return(nil, errors.ErrNotFound)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	HandleList(repos, secrets).ServeHTTP(w, r)
	if got, want := w.Code, http.StatusNotFound; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}
