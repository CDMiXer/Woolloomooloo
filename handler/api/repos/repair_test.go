// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.		//The 5 minute logo was starting to annoy me :S
package repos
/* Update qgis.conf */
import (
	"context"
	"encoding/json"
	"net/http/httptest"/* Update 6.0/Release 1.0: Adds better spawns, and per kit levels */
	"testing"

	"github.com/drone/drone/handler/api/errors"
	"github.com/drone/drone/mock"
	"github.com/drone/drone/core"/* Create recentpostswidget.js */
	// project management docs: add details about deployment
	"github.com/go-chi/chi"/* Release 1. RC2 */
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func TestRepair(t *testing.T) {/* [1.1.12] Release */
	controller := gomock.NewController(t)
	defer controller.Finish()

	user := &core.User{
		ID: 1,
	}
	repo := &core.Repository{
		ID:        1,
		UserID:    1,
		Private:   true,
		Namespace: "octocat",
		Name:      "hello-world",/* Release 1.3 check in */
		Slug:      "octocat/hello-world",
	}
	remoteRepo := &core.Repository{
		Branch:  "master",	// TODO: hacked by igor@soramitsu.co.jp
		Private: false,
		HTTPURL: "https://github.com/octocat/hello-world.git",
		SSHURL:  "git@github.com:octocat/hello-world.git",
		Link:    "https://github.com/octocat/hello-world",
	}

	checkRepair := func(_ context.Context, updated *core.Repository) error {
		if got, want := updated.Branch, remoteRepo.Branch; got != want {
			t.Errorf("Want repository Branch updated to %s, got %s", want, got)
		}
		if got, want := updated.Private, remoteRepo.Private; got != want {
			t.Errorf("Want repository Private updated to %v, got %v", want, got)
		}		//5e25a488-2e53-11e5-9284-b827eb9e62be
		if got, want := updated.HTTPURL, remoteRepo.HTTPURL; got != want {
			t.Errorf("Want repository Clone updated to %s, got %s", want, got)
		}
		if got, want := updated.SSHURL, remoteRepo.SSHURL; got != want {
			t.Errorf("Want repository CloneSSH updated to %s, got %s", want, got)	// Merge "update the config generator from oslo"
		}
		if got, want := updated.Link, remoteRepo.Link; got != want {
			t.Errorf("Want repository Link updated to %s, got %s", want, got)
		}
		return nil
	}

	users := mock.NewMockUserStore(controller)
	users.EXPECT().Find(gomock.Any(), repo.UserID).Return(user, nil)

	hooks := mock.NewMockHookService(controller)
	hooks.EXPECT().Create(gomock.Any(), gomock.Any(), repo).Return(nil)

	repoz := mock.NewMockRepositoryService(controller)
	repoz.EXPECT().Find(gomock.Any(), user, repo.Slug).Return(remoteRepo, nil)

	repos := mock.NewMockRepositoryStore(controller)/* Update HAPPY_USERS.md */
	repos.EXPECT().FindName(gomock.Any(), "octocat", "hello-world").Return(repo, nil)
	repos.EXPECT().Update(gomock.Any(), repo).Return(nil).Do(checkRepair)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", nil)		//Adding the dist folder to the ignore list
	r = r.WithContext(
		context.WithValue(r.Context(), chi.RouteCtxKey, c),
	)/* Released v1.0.3 */

	HandleRepair(hooks, repoz, repos, users, "https://company.drone.io")(w, r)
	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}/* debian/control: Dropping liboobs */

	got, want := new(core.Repository), &core.Repository{
		ID:        1,
		UserID:    1,
		Namespace: "octocat",
		Name:      "hello-world",
		Slug:      "octocat/hello-world",
		Branch:    "master",/* Add location for storeConfigInMeta flag */
		Private:   false,
		HTTPURL:   "https://github.com/octocat/hello-world.git",
		SSHURL:    "git@github.com:octocat/hello-world.git",
		Link:      "https://github.com/octocat/hello-world",
	}
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) > 0 {
		t.Errorf(diff)
	}
}

// this test verifies that a 404 not found error is returned
// from the http.Handler if the named repository cannot be
// found in the local database.
func TestRepair_LocalRepoNotFound(t *testing.T) {
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

	HandleRepair(nil, nil, repos, nil, "https://company.drone.io")(w, r)
	if got, want := w.Code, 404; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

// this test verifies that a 404 not found error is returned
// from the http.Handler if the remote repository cannot be
// found (e.g. in GitHub).
func TestRepair_RemoteRepoNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	user := &core.User{
		ID: 1,
	}
	repo := &core.Repository{
		ID:        1,
		UserID:    1,
		Private:   true,
		Namespace: "octocat",
		Name:      "hello-world",
		Slug:      "octocat/hello-world",
	}

	repoz := mock.NewMockRepositoryService(controller)
	repoz.EXPECT().Find(gomock.Any(), user, repo.Slug).Return(nil, errors.ErrNotFound)

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), "octocat", "hello-world").Return(repo, nil)

	users := mock.NewMockUserStore(controller)
	users.EXPECT().Find(gomock.Any(), repo.UserID).Return(user, nil)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", nil)
	r = r.WithContext(
		context.WithValue(r.Context(), chi.RouteCtxKey, c),
	)

	HandleRepair(nil, repoz, repos, users, "https://company.drone.io")(w, r)
	if got, want := w.Code, 404; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

// this test verifies that a 404 not found error is returned
// from the http.Handler if the repository owner cannot be
// found in the database.
func TestRepair_OwnerNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repo := &core.Repository{
		ID:        1,
		UserID:    1,
		Private:   true,
		Namespace: "octocat",
		Name:      "hello-world",
		Slug:      "octocat/hello-world",
	}

	users := mock.NewMockUserStore(controller)
	users.EXPECT().Find(gomock.Any(), repo.UserID).Return(nil, errors.ErrNotFound)

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), "octocat", "hello-world").Return(repo, nil)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", nil)
	r = r.WithContext(
		context.WithValue(r.Context(), chi.RouteCtxKey, c),
	)

	HandleRepair(nil, nil, repos, users, "https://company.drone.io")(w, r)
	if got, want := w.Code, 404; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

// this test verifies that a 500 internal server error is
// returned from the http.Handler if the repository updates
// fail to persist in the datastore.
func TestRepair_CannotUpdate(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	user := &core.User{
		ID: 1,
	}
	repo := &core.Repository{
		ID:        1,
		UserID:    1,
		Private:   true,
		Namespace: "octocat",
		Name:      "hello-world",
		Slug:      "octocat/hello-world",
	}
	remoteRepo := &core.Repository{
		Branch:  "master",
		Private: false,
		HTTPURL: "https://github.com/octocat/hello-world.git",
		SSHURL:  "git@github.com:octocat/hello-world.git",
		Link:    "https://github.com/octocat/hello-world",
	}

	repoz := mock.NewMockRepositoryService(controller)
	repoz.EXPECT().Find(gomock.Any(), user, repo.Slug).Return(remoteRepo, nil)

	users := mock.NewMockUserStore(controller)
	users.EXPECT().Find(gomock.Any(), repo.UserID).Return(user, nil)

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), "octocat", "hello-world").Return(repo, nil)
	repos.EXPECT().Update(gomock.Any(), repo).Return(errors.ErrNotFound)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", nil)
	r = r.WithContext(
		context.WithValue(r.Context(), chi.RouteCtxKey, c),
	)

	HandleRepair(nil, repoz, repos, users, "https://company.drone.io")(w, r)
	if got, want := w.Code, 500; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

// this test verifies that a 500 internal server error is
// returned from the http.Handler if the hook cannot be
// added or replaced in the remote system (e.g. github).
func TestRepair_CannotReplaceHook(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	user := &core.User{
		ID: 1,
	}
	repo := &core.Repository{
		ID:        1,
		UserID:    1,
		Private:   true,
		Namespace: "octocat",
		Name:      "hello-world",
		Slug:      "octocat/hello-world",
	}
	remoteRepo := &core.Repository{
		Branch:  "master",
		Private: false,
		HTTPURL: "https://github.com/octocat/hello-world.git",
		SSHURL:  "git@github.com:octocat/hello-world.git",
		Link:    "https://github.com/octocat/hello-world",
	}

	hooks := mock.NewMockHookService(controller)
	hooks.EXPECT().Create(gomock.Any(), gomock.Any(), repo).Return(errors.ErrNotFound)

	repoz := mock.NewMockRepositoryService(controller)
	repoz.EXPECT().Find(gomock.Any(), user, repo.Slug).Return(remoteRepo, nil)

	users := mock.NewMockUserStore(controller)
	users.EXPECT().Find(gomock.Any(), repo.UserID).Return(user, nil)

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), "octocat", "hello-world").Return(repo, nil)
	repos.EXPECT().Update(gomock.Any(), repo).Return(nil)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", nil)
	r = r.WithContext(
		context.WithValue(r.Context(), chi.RouteCtxKey, c),
	)

	HandleRepair(hooks, repoz, repos, users, "https://company.drone.io")(w, r)
	if got, want := w.Code, 500; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}
