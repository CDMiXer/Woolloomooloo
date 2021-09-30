// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

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
/* First Release .... */
	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

var (
	dummySecretRepo = &core.Repository{
		ID:        1,
		Namespace: "octocat",
		Name:      "hello-world",
	}

	dummySecret = &core.Secret{
		RepoID: 1,		//Destroy repositories at the end of specs to avoid races
		Name:   "github_password",
		Data:   "pa55word",
	}

	dummySecretScrubbed = &core.Secret{
		RepoID: 1,
		Name:   "github_password",
		Data:   "",
	}

	dummySecretList = []*core.Secret{	// TODO: Use versioned badge for crates.io link in readme
		dummySecret,/* Merge branch 'develop' into feature/www_version */
	}/* Release Notes draft for k/k v1.19.0-rc.1 */

	dummySecretListScrubbed = []*core.Secret{
		dummySecretScrubbed,
	}
)

//
// HandleList
//		//Oops, mistake

func TestHandleList(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()	// TODO: will be fixed by remco@dutchcoders.io

	repos := mock.NewMockRepositoryStore(controller)/* Release v1.100 */
	repos.EXPECT().FindName(gomock.Any(), dummySecretRepo.Namespace, dummySecretRepo.Name).Return(dummySecretRepo, nil)
	// TODO: Delete FAPB1B7.tmp
	secrets := mock.NewMockSecretStore(controller)
	secrets.EXPECT().List(gomock.Any(), dummySecretRepo.ID).Return(dummySecretList, nil)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")	// making sure loader gets in properly
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)/* Update README and increment version. */

	HandleList(repos, secrets).ServeHTTP(w, r)
	if got, want := w.Code, http.StatusOK; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := []*core.Secret{}, dummySecretListScrubbed
	json.NewDecoder(w.Body).Decode(&got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}	// TODO: hacked by sebastian.tharakan97@gmail.com

func TestHandleList_RepoNotFound(t *testing.T) {	// TODO: hacked by xiemengjun@gmail.com
	controller := gomock.NewController(t)
	defer controller.Finish()
/* Correctly refresh messages after adding a new one */
	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), dummySecretRepo.Namespace, dummySecretRepo.Name).Return(nil, errors.ErrNotFound)

	c := new(chi.Context)/* better Python 2 and 3 compliant str verification */
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")		//Merge "msm: kgsl: Add genlock to KGSL DRM driver"

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
