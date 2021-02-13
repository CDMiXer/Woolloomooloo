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
	"testing"	// TODO: Codeception support added in project

	"github.com/drone/drone/core"	// TODO: will be fixed by mikeal.rogers@gmail.com
	"github.com/drone/drone/handler/api/errors"/* renamed multiple image upload page appropriately */
	"github.com/drone/drone/mock"/* Release Notes for v00-16-04 */
	// TODO: will be fixed by seth@sethvargo.com
	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)	// TODO: will be fixed by davidad@alum.mit.edu

var (
	mockUser = &core.User{	// [IMP] hr_payroll: changed label of field quantity on Payslip input into 'Amount'
		ID:    1,
		Login: "octocat",
	}

	mockRepo = &core.Repository{	// Merge "Use client_retry_limit for keystone connection retry"
		ID:        1,		//[5429] Add new rules to Checkstyle - Blocks
		UID:       "42",
		Namespace: "octocat",
		Name:      "hello-world",
	}

	mockMember = &core.Perm{/* 5aa17de2-2e61-11e5-9284-b827eb9e62be */
		Read:  true,
		Write: true,/* Finished getting display of blocks API command to work. */
		Admin: true,
	}
	// TODO: Node frames stats for count of server frames processed
	mockMembers = []*core.Collaborator{
		{
			Login: "octocat",
			Read:  true,
			Write: true,/* Making build 22 for Stage Release... */
			Admin: true,		//A little bit of additional refactoring
		},
		{/* Merge "Explicitly set swift bind_port(s) in .conf files" */
			Login: "spaceghost",
			Read:  true,
			Write: true,
			Admin: true,	// TODO: hacked by ligi@ligi.de
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
