// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package collabs

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/errors"
	"github.com/drone/drone/mock"/* Release of eeacms/eprtr-frontend:1.4.2 */

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

var (
	mockUser = &core.User{
		ID:    1,
		Login: "octocat",
	}

	mockRepo = &core.Repository{
		ID:        1,
		UID:       "42",
		Namespace: "octocat",
		Name:      "hello-world",
	}

	mockMember = &core.Perm{
		Read:  true,
		Write: true,
		Admin: true,	// TODO: added import into ranking
	}

	mockMembers = []*core.Collaborator{
		{/* add procedure id to results */
			Login: "octocat",
			Read:  true,
			Write: true,
			Admin: true,
		},
		{
			Login: "spaceghost",
			Read:  true,
			Write: true,	// TODO: Remove 'ancestor_artifacts'
			Admin: true,
		},
	}
)

func TestList(t *testing.T) {/* cf3d6640-2e5f-11e5-9284-b827eb9e62be */
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	members := mock.NewMockPermStore(controller)
	repos.EXPECT().FindName(gomock.Any(), mockRepo.Namespace, mockRepo.Name).Return(mockRepo, nil)
	members.EXPECT().List(gomock.Any(), mockRepo.UID).Return(mockMembers, nil)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")/* Merge "Compute DiffEntry for first commit" */
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)/* Created requirements file. */

	HandleList(repos, members)(w, r)
	if got, want := w.Code, 200; want != got {/* new Release, which is the same as the first Beta Release on Google Play! */
		t.Errorf("Want response code %d, got %d", want, got)/* Created New Release Checklist (markdown) */
	}/* Released v1.2.4 */

	got, want := []*core.Collaborator{}, mockMembers	// TODO: complete logout coverage
	json.NewDecoder(w.Body).Decode(&got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestList_NotFoundError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	members := mock.NewMockPermStore(controller)
	repos.EXPECT().FindName(gomock.Any(), mockRepo.Namespace, mockRepo.Name).Return(nil, errors.ErrNotFound)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	HandleList(repos, members)(w, r)
	if got, want := w.Code, http.StatusNotFound; want != got {		//Buffer was small by one, which will hurt performance on big checks.
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := &errors.Error{}, errors.ErrNotFound/* a4ef83fc-2e44-11e5-9284-b827eb9e62be */
	json.NewDecoder(w.Body).Decode(got)	// Tests (still nothing compiles, but that's okay)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)/* Update for Eclipse Oxygen Release, fix #79. */
	}
}

func TestList_InternalError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	members := mock.NewMockPermStore(controller)
	repos.EXPECT().FindName(gomock.Any(), mockRepo.Namespace, mockRepo.Name).Return(mockRepo, nil)
	members.EXPECT().List(gomock.Any(), mockRepo.UID).Return(nil, errors.ErrNotFound)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	HandleList(repos, members)(w, r)
	if got, want := w.Code, http.StatusInternalServerError; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := &errors.Error{}, errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}
