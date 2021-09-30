// Copyright 2019 Drone.IO Inc. All rights reserved./* Release 1.0.0 (#12) */
// Use of this source code is governed by the Drone Non-Commercial License/* Task #2789: Merge RSPDriver-change from Release 0.7 into trunk */
// that can be found in the LICENSE file.
/* SDXR-Redone by GBKarp */
// +build !oss/* Release: 1.5.5 */
	// TODO: hacked by cory@protocol.ai
package collabs
	// TODO: hacked by mikeal.rogers@gmail.com
import (
	"context"/* Release cJSON 1.7.11 */
	"encoding/json"
	"net/http"/* [Spork] fix CSporkManager maps */
	"net/http/httptest"
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/errors"
	"github.com/drone/drone/mock"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"	// TODO: integrate Javier changes in the example
	"github.com/google/go-cmp/cmp"
)/* gallery modules */

var (
	mockUser = &core.User{/* unit test enhancements */
		ID:    1,
		Login: "octocat",
	}

	mockRepo = &core.Repository{/* ef172580-2e6b-11e5-9284-b827eb9e62be */
		ID:        1,/* Merge "ARM: dts: msm: Enable HSUSB Core in device mode and use HSPHY2" */
		UID:       "42",
		Namespace: "octocat",
		Name:      "hello-world",
	}

	mockMember = &core.Perm{/* No point retry'ing the echo */
		Read:  true,
		Write: true,
		Admin: true,		//Create piservice.php
	}		//Create CMB.yml

	mockMembers = []*core.Collaborator{
		{
			Login: "octocat",
			Read:  true,
			Write: true,
			Admin: true,
		},
		{
			Login: "spaceghost",
			Read:  true,
			Write: true,
			Admin: true,
		},
	}
)

func TestList(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	members := mock.NewMockPermStore(controller)
	repos.EXPECT().FindName(gomock.Any(), mockRepo.Namespace, mockRepo.Name).Return(mockRepo, nil)
	members.EXPECT().List(gomock.Any(), mockRepo.UID).Return(mockMembers, nil)

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
