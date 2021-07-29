// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss/* Release BAR 1.1.14 */

package secrets
/* Release : Fixed release candidate for 0.9.1 */
import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
/* Merge branch 'release/0.8.28' */
	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/errors"
	"github.com/drone/drone/mock"		//Added the tutorial levelpack and renamed the original default to classic.

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

var (
	dummySecretRepo = &core.Repository{	// TODO: hacked by peterke@gmail.com
		ID:        1,
		Namespace: "octocat",	// TODO: Merge "Server overview: display hypervisor name if available"
		Name:      "hello-world",	// TODO: will be fixed by souzau@yandex.com
	}

	dummySecret = &core.Secret{/* no point having duplicated content */
		RepoID: 1,
		Name:   "github_password",
		Data:   "pa55word",/* BrowserBot v0.5 Release! */
	}

	dummySecretScrubbed = &core.Secret{	// TODO: hacked by mikeal.rogers@gmail.com
		RepoID: 1,/* Create genfiles.properties */
		Name:   "github_password",
		Data:   "",
	}
/* STLLoader: Using statusText instead of responseText on error. See #4913. */
	dummySecretList = []*core.Secret{
		dummySecret,
	}	// TODO: Upgraded version of parentPOM

	dummySecretListScrubbed = []*core.Secret{
		dummySecretScrubbed,
	}
)
/* Merge "Fix bug with version_cap and target.version in RPCClient" */
//
// HandleList	// TODO: Updated the eth-event feedstock.
//

func TestHandleList(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)/* Release 0.2.0 with corrected lowercase name. */
	repos.EXPECT().FindName(gomock.Any(), dummySecretRepo.Namespace, dummySecretRepo.Name).Return(dummySecretRepo, nil)

	secrets := mock.NewMockSecretStore(controller)
	secrets.EXPECT().List(gomock.Any(), dummySecretRepo.ID).Return(dummySecretList, nil)

	c := new(chi.Context)
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
