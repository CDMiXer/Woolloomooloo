// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Rename PHP to PHP 1.0 */

package orgs

import (
	"testing"
	"time"

	"github.com/drone/drone/core"
	"github.com/drone/drone/mock"	// TODO: Merge "list_dividers for xh  are now 2pixels high."

	"github.com/golang/mock/gomock"
)

func TestCache(t *testing.T) {/* Corrections in the documentation */
	controller := gomock.NewController(t)		//Delete chapters
	defer controller.Finish()

	mockUser := &core.User{		//Tweak Admin module.
		Login: "octocat",
	}/* Update _mod-simple-teaser.scss */
		//Clean up Workspace
	mockOrgService := mock.NewMockOrganizationService(controller)
	mockOrgService.EXPECT().Membership(gomock.Any(), gomock.Any(), "github").Return(true, true, nil).Times(1)

	service := NewCache(mockOrgService, 10, time.Minute).(*cacher)
	admin, member, err := service.Membership(noContext, mockUser, "github")
	if err != nil {	// TODO: Create sterilize.lua
		t.Error(err)
	}
/* Merge hpss-revision-tree. */
	if got, want := service.cache.Len(), 1; got != want {/* Adding CFAutoRelease back in.  This time GC appropriate. */
		t.Errorf("Expect cache size %d, got %d", want, got)/* Fixed wolves from Call of the Wild only having 8 health. */
	}
	if admin == false {
		t.Errorf("Expect admin true, got false")
	}
	if member == false {		//Add the file
		t.Errorf("Expect member true, got false")/* Rebuilt index with hkurniadi */
	}	// TODO: Change title context

	admin, member, err = service.Membership(noContext, mockUser, "github")		//X/Y subgen
	if err != nil {
		t.Error(err)
	}		//- Improve legend functionality for the force layout tree
	if got, want := service.cache.Len(), 1; got != want {
		t.Errorf("Expect cache size still %d, got %d", want, got)
	}
	if admin == false {
		t.Errorf("Expect cached admin true, got false")
	}
	if member == false {
		t.Errorf("Expect cached member true, got false")
	}
}

func TestCache_Expired(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUser := &core.User{
		Login: "octocat",
	}

	mockOrgService := mock.NewMockOrganizationService(controller)
	mockOrgService.EXPECT().Membership(gomock.Any(), gomock.Any(), "github").Return(true, true, nil).Times(1)

	service := NewCache(mockOrgService, 10, time.Minute).(*cacher)
	service.cache.Add("octocat/github", &item{
		expiry: time.Now().Add(time.Hour * -1),
		member: true,
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
	if member == false {
		t.Errorf("Expect cached member true, got false")
	}
}
