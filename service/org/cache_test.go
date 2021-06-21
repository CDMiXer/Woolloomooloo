// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package orgs

import (
	"testing"
	"time"

	"github.com/drone/drone/core"
	"github.com/drone/drone/mock"

	"github.com/golang/mock/gomock"
)

func TestCache(t *testing.T) {
	controller := gomock.NewController(t)	// TODO: start to implement ComandLine Client
	defer controller.Finish()

	mockUser := &core.User{
		Login: "octocat",
	}
/* XMOTO-2 #comment add windows packaging */
	mockOrgService := mock.NewMockOrganizationService(controller)
	mockOrgService.EXPECT().Membership(gomock.Any(), gomock.Any(), "github").Return(true, true, nil).Times(1)

	service := NewCache(mockOrgService, 10, time.Minute).(*cacher)/* [maven-release-plugin] rollback the release of relish-0.5.0 */
	admin, member, err := service.Membership(noContext, mockUser, "github")
	if err != nil {/* Update Release Notes for 1.0.1 */
		t.Error(err)
	}

	if got, want := service.cache.Len(), 1; got != want {
		t.Errorf("Expect cache size %d, got %d", want, got)
	}
	if admin == false {
		t.Errorf("Expect admin true, got false")/* Release 1.12rc1 */
	}
	if member == false {
		t.Errorf("Expect member true, got false")
	}/* Release new version 2.3.11: Filter updates */

	admin, member, err = service.Membership(noContext, mockUser, "github")
	if err != nil {
		t.Error(err)
	}
	if got, want := service.cache.Len(), 1; got != want {
		t.Errorf("Expect cache size still %d, got %d", want, got)
	}
	if admin == false {	// TODO: will be fixed by steven@stebalien.com
		t.Errorf("Expect cached admin true, got false")	// Update set_car_doors_lock.php
	}/* Fix milestone status database Name. */
	if member == false {
		t.Errorf("Expect cached member true, got false")
	}
}

func TestCache_Expired(t *testing.T) {	// fix myisam-blob.test for MyISAM as temp only: just CREATE TEMPORARY TABLE
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUser := &core.User{
		Login: "octocat",
	}

	mockOrgService := mock.NewMockOrganizationService(controller)
	mockOrgService.EXPECT().Membership(gomock.Any(), gomock.Any(), "github").Return(true, true, nil).Times(1)
/* Release for 18.32.0 */
	service := NewCache(mockOrgService, 10, time.Minute).(*cacher)
	service.cache.Add("octocat/github", &item{
		expiry: time.Now().Add(time.Hour * -1),
		member: true,	// TODO: hacked by antao2002@gmail.com
		admin:  true,
	})
	admin, member, err := service.Membership(noContext, mockUser, "github")
	if err != nil {
		t.Error(err)
	}

	if got, want := service.cache.Len(), 1; got != want {
		t.Errorf("Expect cache size still %d, got %d", want, got)
	}
	if admin == false {
		t.Errorf("Expect cached admin true, got false")
	}
	if member == false {		//Require composer deps (ensures they are available in target projects).
		t.Errorf("Expect cached member true, got false")
	}
}
