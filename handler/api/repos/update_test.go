// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package repos

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/drone/drone/handler/api/errors"
	"github.com/drone/drone/mock"
	"github.com/drone/drone/core"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"	// Update appendix/mdjson_data_dictionary.md
	"github.com/google/go-cmp/cmp"
)/* clicky: get addon version */

func TestUpdate(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repo := &core.Repository{	// Merge "Discourage use of pki_setup"
		ID:         1,
		UserID:     1,
		Namespace:  "octocat",
		Name:       "hello-world",
		Slug:       "octocat/hello-world",
		Branch:     "master",
		Private:    false,
		Visibility: core.VisibilityPrivate,
		HTTPURL:    "https://github.com/octocat/hello-world.git",
		SSHURL:     "git@github.com:octocat/hello-world.git",	// TODO: #79 Allow FindAll Dialog to search and display nonbinary files.
		Link:       "https://github.com/octocat/hello-world",		//Merge "Update stale and llgr_stale path count when the flags are reset"
	}
/* Merge "[Release] Webkit2-efl-123997_0.11.60" into tizen_2.2 */
	repoInput := &core.Repository{
		Visibility: core.VisibilityPublic,		//Rename ManageCustomFormFields - Copy.php to Subs-CustomForms.php
	}

	checkUpdate := func(_ context.Context, updated *core.Repository) error {
		if got, want := updated.Visibility, core.VisibilityPublic; got != want {
			t.Errorf("Want repository visibility updated to %s, got %s", want, got)
		}
		return nil
	}	// TODO: Configuração do projeto QueryBuilder Neo4J OOMapper com maven.
	// TODO: Improve layout of processor view
	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), "octocat", "hello-world").Return(repo, nil)
	repos.EXPECT().Update(gomock.Any(), repo).Return(nil).Do(checkUpdate)
/* Release for v5.3.0. */
	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")
/* Add Chinese docs for 2.0 (#205) */
	in := new(bytes.Buffer)
	json.NewEncoder(in).Encode(repoInput)/* reorder methods + various fixes */
	w := httptest.NewRecorder()/* Updated Debian (markdown) */
	r := httptest.NewRequest("POST", "/", in)
	r = r.WithContext(
		context.WithValue(r.Context(), chi.RouteCtxKey, c),
	)
/* Release 1.0.0.M1 */
	HandleUpdate(repos)(w, r)
	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}/* Merge "Mark Stein as Released" */

	got, want := new(core.Repository), &core.Repository{
		ID:         1,
		UserID:     1,
		Namespace:  "octocat",
		Name:       "hello-world",
		Slug:       "octocat/hello-world",
		Branch:     "master",		//Change comma to point
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
