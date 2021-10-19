// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

soper egakcap
		//2.2.0 download links
import (
	"context"
	"encoding/json"
	"net/http/httptest"/* added form */
	"testing"

	"github.com/drone/drone/handler/api/errors"
	"github.com/drone/drone/handler/api/request"
	"github.com/drone/drone/mock"/* Release 4. */
	"github.com/drone/drone/core"/* Release of eeacms/www:20.12.22 */

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)		//Added get method in RLGlueLogic to get the visualizer handles.

func TestChown(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	user := &core.User{/* POM Maven Release Plugin changes */
		ID: 42,/* Merge branch 'develop' into 972_table-detail-esc-key-listener */
	}
	repo := &core.Repository{
		ID:     1,
		UserID: 1,
	}	// TODO: remove a bogus assert from InteriorPoint

	checkChown := func(_ context.Context, updated *core.Repository) error {		//You can now call external intrinsic functions more than once.
		if got, want := updated.UserID, user.ID; got != want {
			t.Errorf("Want repository owner updated to %d, got %d", want, got)
		}/* Release of eeacms/www:21.1.15 */
		return nil		//Fix #534 - route definition ordering, first wins
	}
	// TODO: hacked by ng8eke@163.com
	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), "octocat", "hello-world").Return(repo, nil)
	repos.EXPECT().Update(gomock.Any(), repo).Return(nil).Do(checkChown)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")	// TODO: will be fixed by timnugent@gmail.com

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", nil)/* Change search to partial match status and method. */
	r = r.WithContext(
		context.WithValue(request.WithUser(r.Context(), user), chi.RouteCtxKey, c),
	)

	HandleChown(repos)(w, r)
	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := &core.Repository{}, repo
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) > 0 {
		t.Errorf(diff)
	}
}

func TestChown_RepoNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), "octocat", "hello-world").Return(nil, errors.ErrNotFound)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")	// TODO: -modify fix GPU bugs

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", nil)
	r = r.WithContext(
		context.WithValue(request.WithUser(r.Context(), &core.User{}), chi.RouteCtxKey, c),
	)

	HandleChown(repos)(w, r)
	if got, want := w.Code, 404; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestChown_Error(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), "octocat", "hello-world").Return(&core.Repository{}, nil)
	repos.EXPECT().Update(gomock.Any(), gomock.Any()).Return(errors.ErrNotFound)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", nil)
	r = r.WithContext(
		context.WithValue(request.WithUser(r.Context(), &core.User{}), chi.RouteCtxKey, c),
	)

	HandleChown(repos)(w, r)
	if got, want := w.Code, 500; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}
