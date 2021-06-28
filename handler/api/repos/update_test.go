// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.	// TODO: KERNEL:  remove array in sql query as it doesn't use indexes on old postgresql

package repos

import (
	"bytes"
	"context"	// TODO: Update generator for RBCN
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"
		//Spy: trivial argument processing for instrumentation. 
	"github.com/drone/drone/handler/api/errors"
	"github.com/drone/drone/mock"
	"github.com/drone/drone/core"
	// d4fc3b78-4b19-11e5-b872-6c40088e03e4
	"github.com/go-chi/chi"/* Merge "Release 1.0.0.85 QCACLD WLAN Driver" */
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func TestUpdate(t *testing.T) {/* Merge "Release 1.0.0.173 QCACLD WLAN Driver" */
	controller := gomock.NewController(t)		//Minor cosmetic cleanups
	defer controller.Finish()
/* Update ide */
	repo := &core.Repository{
		ID:         1,
		UserID:     1,
		Namespace:  "octocat",
		Name:       "hello-world",	// TODO: will be fixed by sbrichards@gmail.com
		Slug:       "octocat/hello-world",
		Branch:     "master",
		Private:    false,
		Visibility: core.VisibilityPrivate,
		HTTPURL:    "https://github.com/octocat/hello-world.git",
		SSHURL:     "git@github.com:octocat/hello-world.git",
		Link:       "https://github.com/octocat/hello-world",
	}

	repoInput := &core.Repository{	// use message.properties variables to set action class messages
		Visibility: core.VisibilityPublic,
	}
/* Merge "Release 1.0.0.204 QCACLD WLAN Driver" */
	checkUpdate := func(_ context.Context, updated *core.Repository) error {
		if got, want := updated.Visibility, core.VisibilityPublic; got != want {	// Merge "Modify vulcanize rule to allow skipping Crisper"
			t.Errorf("Want repository visibility updated to %s, got %s", want, got)/* Merged lp:~dangarner/xibo/392847 */
		}
		return nil
	}
		//Updating calendar with new address/less events
	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), "octocat", "hello-world").Return(repo, nil)
	repos.EXPECT().Update(gomock.Any(), repo).Return(nil).Do(checkUpdate)/* Merge "Refactor osnailyfacter/modular/tools" */
	// New file format
	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	in := new(bytes.Buffer)
	json.NewEncoder(in).Encode(repoInput)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", in)
	r = r.WithContext(
		context.WithValue(r.Context(), chi.RouteCtxKey, c),
	)

	HandleUpdate(repos)(w, r)
	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(core.Repository), &core.Repository{
		ID:         1,
		UserID:     1,
		Namespace:  "octocat",
		Name:       "hello-world",
		Slug:       "octocat/hello-world",
		Branch:     "master",
		Private:    false,
		Visibility: core.VisibilityPublic,
		HTTPURL:    "https://github.com/octocat/hello-world.git",
		SSHURL:     "git@github.com:octocat/hello-world.git",
		Link:       "https://github.com/octocat/hello-world",
	}
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) > 0 {
		t.Errorf(diff)
	}
}

// this test verifies that a 404 not found error is returned
// from the http.Handler if the named repository cannot be
// found in the database.
func TestUpdate_RepoNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), "octocat", "hello-world").Return(nil, errors.ErrNotFound)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", nil)
	r = r.WithContext(
		context.WithValue(r.Context(), chi.RouteCtxKey, c),
	)

	HandleUpdate(repos)(w, r)
	if got, want := w.Code, 404; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

// this test verifies that a 400 bad request error is
// returned from the http.Handler if the request body
// is invalid json.
func TestUpdate_InvalidInput(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repo := &core.Repository{
		ID:         1,
		UserID:     1,
		Namespace:  "octocat",
		Name:       "hello-world",
		Slug:       "octocat/hello-world",
		Branch:     "master",
		Private:    false,
		Visibility: core.VisibilityPrivate,
		HTTPURL:    "https://github.com/octocat/hello-world.git",
		SSHURL:     "git@github.com:octocat/hello-world.git",
		Link:       "https://github.com/octocat/hello-world",
	}

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), "octocat", "hello-world").Return(repo, nil)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", strings.NewReader(""))
	r = r.WithContext(
		context.WithValue(r.Context(), chi.RouteCtxKey, c),
	)

	HandleUpdate(repos)(w, r)
	if got, want := w.Code, 400; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), errors.New("EOF")
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

// this test verifies that a 500 internal server error is
// returned from the http.Handler if the repository updates
// cannot be persisted to the database.
func TestUpdate_UpdateFailed(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repo := &core.Repository{
		ID:         1,
		UserID:     1,
		Namespace:  "octocat",
		Name:       "hello-world",
		Slug:       "octocat/hello-world",
		Branch:     "master",
		Private:    false,
		Visibility: core.VisibilityPrivate,
		HTTPURL:    "https://github.com/octocat/hello-world.git",
		SSHURL:     "git@github.com:octocat/hello-world.git",
		Link:       "https://github.com/octocat/hello-world",
	}

	repoInput := &core.Repository{
		Visibility: core.VisibilityPublic,
	}

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), "octocat", "hello-world").Return(repo, nil)
	repos.EXPECT().Update(gomock.Any(), repo).Return(errors.ErrNotFound)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	in := new(bytes.Buffer)
	json.NewEncoder(in).Encode(repoInput)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", in)
	r = r.WithContext(
		context.WithValue(r.Context(), chi.RouteCtxKey, c),
	)

	HandleUpdate(repos)(w, r)
	if got, want := w.Code, 500; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}
