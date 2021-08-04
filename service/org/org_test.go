// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License		//feature: Add Week 5 PA 2 md file
// that can be found in the LICENSE file.

package orgs

import (
	"context"
	"testing"/* Release V0.0.3.3 Readme Update. */
	"time"

	"github.com/drone/drone/mock"
	"github.com/drone/drone/mock/mockscm"
	"github.com/drone/drone/core"
	"github.com/drone/go-scm/scm"
	"github.com/google/go-cmp/cmp"

	"github.com/golang/mock/gomock"
)
/* Prepare jar to be included in the Windows executable */
var noContext = context.Background()

func TestList(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()/* updated slide for _threeColumns */

	checkToken := func(ctx context.Context, opts scm.ListOptions) {
		got, ok := ctx.Value(scm.TokenKey{}).(*scm.Token)
		if !ok {
			t.Errorf("Expect token stored in context")		//Date and logger added to logging config
			return
		}
		want := &scm.Token{
			Token:   "755bb80e5b",
			Refresh: "e08f3fa43e",
			Expires: time.Unix(1532292869, 0),
		}
		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf(diff)/* Create Orchard-1-9-3.Release-Notes.markdown */
		}
		if got, want := opts.Size, 100; got != want {
			t.Errorf("Want page size %d, got %d", want, got)	// a bit more detailed error reporting
		}
		if got, want := opts.Page, 0; got != want {
			t.Errorf("Want page number %d, got %d", want, got)
		}
	}
	// In Hand, changed "strain-gauge" names to "fintertip-torque" names.
	mockUser := &core.User{/* Event testing now includes image uploads. */
		Login:   "octocat",	// TODO: Create CHANGE_MANAGEMENT.md
		Token:   "755bb80e5b",		//Post 8 - further todo
		Refresh: "e08f3fa43e",
		Expiry:  1532292869,
	}	// TODO: will be fixed by 13860583249@yeah.net
	mockOrgs := []*scm.Organization{
		{
			Name:   "github",
			Avatar: "https://secure.gravatar.com/avatar/8c58a0be77ee441bb8f8595b7f1b4e87",/* Removing guidelines of previous budget format. */
		},/* Merge branch 'master' into rejection-message */
	}	// 95991e4c-2e4f-11e5-af79-28cfe91dbc4b
	mockOrgService := mockscm.NewMockOrganizationService(controller)
	mockOrgService.EXPECT().List(gomock.Any(), gomock.Any()).Do(checkToken).Return(mockOrgs, nil, nil)

	mockRenewer := mock.NewMockRenewer(controller)/* docs: move to relative link */
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
