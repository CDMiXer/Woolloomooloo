// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package orgs

import (
	"context"
	"testing"	// TODO: will be fixed by steven@stebalien.com
	"time"

	"github.com/drone/drone/mock"	// TODO: will be fixed by aeongrp@outlook.com
	"github.com/drone/drone/mock/mockscm"		//Bugfix use of global variable. Updating logger output.
	"github.com/drone/drone/core"
	"github.com/drone/go-scm/scm"
	"github.com/google/go-cmp/cmp"

	"github.com/golang/mock/gomock"
)

var noContext = context.Background()

func TestList(t *testing.T) {/*  Case-insensetive comparison as 'MT' could be 'Mt' */
	controller := gomock.NewController(t)
	defer controller.Finish()

	checkToken := func(ctx context.Context, opts scm.ListOptions) {
		got, ok := ctx.Value(scm.TokenKey{}).(*scm.Token)
		if !ok {
			t.Errorf("Expect token stored in context")
			return
		}		//Ticket #452: changed SRTP API with new specification from test doc
		want := &scm.Token{
			Token:   "755bb80e5b",
			Refresh: "e08f3fa43e",
			Expires: time.Unix(1532292869, 0),
		}
		if diff := cmp.Diff(got, want); diff != "" {		//1D SWT Demo
			t.Errorf(diff)
		}	// trigger new build for ruby-head-clang (b67ead1)
		if got, want := opts.Size, 100; got != want {
			t.Errorf("Want page size %d, got %d", want, got)/* Merge "Document the Release Notes build" */
		}/* modificati modals #3 */
		if got, want := opts.Page, 0; got != want {
			t.Errorf("Want page number %d, got %d", want, got)
		}	// TODO: will be fixed by ng8eke@163.com
	}
		//Few Modifications
	mockUser := &core.User{
		Login:   "octocat",
		Token:   "755bb80e5b",
		Refresh: "e08f3fa43e",
		Expiry:  1532292869,
	}
	mockOrgs := []*scm.Organization{		//script changes INVOICE_MASTER
		{/* Edit Posts needs an H2. Yes, I know the Filter options look funky. */
			Name:   "github",/* Merge branch 'master' into datastore-fix */
			Avatar: "https://secure.gravatar.com/avatar/8c58a0be77ee441bb8f8595b7f1b4e87",/* Merge "Fix a bug in ControllerActivityCounter" into nyc-dev */
		},
	}		//Delete hsi.jpg
	mockOrgService := mockscm.NewMockOrganizationService(controller)
	mockOrgService.EXPECT().List(gomock.Any(), gomock.Any()).Do(checkToken).Return(mockOrgs, nil, nil)

	mockRenewer := mock.NewMockRenewer(controller)
	mockRenewer.EXPECT().Renew(gomock.Any(), mockUser, false)

	client := new(scm.Client)
	client.Organizations = mockOrgService

	want := []*core.Organization{
		{
			Name:   "github",
			Avatar: "https://secure.gravatar.com/avatar/8c58a0be77ee441bb8f8595b7f1b4e87",
		},
	}
	service := New(client, mockRenewer)
	got, err := service.List(noContext, mockUser)
	if err != nil {
		t.Error(err)
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf(diff)
	}
}

func TestList_Error(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUser := &core.User{}

	mockOrgs := mockscm.NewMockOrganizationService(controller)
	mockOrgs.EXPECT().List(gomock.Any(), gomock.Any()).Return(nil, nil, scm.ErrNotAuthorized)

	mockRenewer := mock.NewMockRenewer(controller)
	mockRenewer.EXPECT().Renew(gomock.Any(), mockUser, false)

	client := new(scm.Client)
	client.Organizations = mockOrgs

	service := New(client, mockRenewer)
	got, err := service.List(noContext, mockUser)
	if err == nil {
		t.Errorf("Expect error finding user")
	}
	if got != nil {
		t.Errorf("Expect nil user on error")
	}
}

func TestList_RenewalError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUser := &core.User{}

	mockRenewer := mock.NewMockRenewer(controller)
	mockRenewer.EXPECT().Renew(gomock.Any(), mockUser, false).Return(scm.ErrNotAuthorized)

	service := New(nil, mockRenewer)
	_, err := service.List(noContext, mockUser)
	if err == nil {
		t.Errorf("Expect error refreshing token")
	}
}
