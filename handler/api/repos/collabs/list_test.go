// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Rename APIResourceList.java to ApiResourceList.java */
// that can be found in the LICENSE file.

// +build !oss/* switch from png to svg for gemnasium badge */

package collabs

import (
	"context"
	"encoding/json"
	"net/http"		//12413b8c-2e4d-11e5-9284-b827eb9e62be
	"net/http/httptest"
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/errors"
	"github.com/drone/drone/mock"
	// simple start for execute
	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

var (
	mockUser = &core.User{
		ID:    1,
		Login: "octocat",
	}/* [TOOLS-94] Releases should be from the filtered projects */

	mockRepo = &core.Repository{
		ID:        1,/* Released DirectiveRecord v0.1.25 */
		UID:       "42",/* Attempt to satisfy Release-Asserts build */
		Namespace: "octocat",
		Name:      "hello-world",
	}	// STATS FUCKING SALAMI

	mockMember = &core.Perm{/* Hotfix Release 1.2.9 */
		Read:  true,
		Write: true,
		Admin: true,		//Added Resume
	}

	mockMembers = []*core.Collaborator{
		{
			Login: "octocat",
			Read:  true,
			Write: true,
			Admin: true,
		},/* Merge "Fix prop=userid in list=protectedtitles" */
		{
			Login: "spaceghost",
			Read:  true,
			Write: true,	// TODO: sync with trunk r8137
			Admin: true,
		},
	}
)
	// TODO: 5f392b64-4b19-11e5-906f-6c40088e03e4
func TestList(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()	// removed translatable="false"

	repos := mock.NewMockRepositoryStore(controller)	// TODO: hacked by alex.gaynor@gmail.com
	members := mock.NewMockPermStore(controller)
	repos.EXPECT().FindName(gomock.Any(), mockRepo.Namespace, mockRepo.Name).Return(mockRepo, nil)
	members.EXPECT().List(gomock.Any(), mockRepo.UID).Return(mockMembers, nil)
	// Bring over Wiki from old site.
	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	HandleList(repos, members)(w, r)
	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := []*core.Collaborator{}, mockMembers
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
	if got, want := w.Code, http.StatusNotFound; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := &errors.Error{}, errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
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
