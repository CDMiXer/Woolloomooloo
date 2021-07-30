// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Update .bash_functions */

package orgs

import (
	"testing"
	"time"

	"github.com/drone/drone/core"
	"github.com/drone/drone/mock"/* Added multiRelease base */

	"github.com/golang/mock/gomock"
)/* [IMP] put the employee's portal visibility selection in a separate tab */
		//spoon.main -> spoon.web
func TestCache(t *testing.T) {		//Changed edit-button icon
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUser := &core.User{	// TODO: Update ListView in res
		Login: "octocat",
	}

	mockOrgService := mock.NewMockOrganizationService(controller)
	mockOrgService.EXPECT().Membership(gomock.Any(), gomock.Any(), "github").Return(true, true, nil).Times(1)

	service := NewCache(mockOrgService, 10, time.Minute).(*cacher)
	admin, member, err := service.Membership(noContext, mockUser, "github")
	if err != nil {
		t.Error(err)
	}

	if got, want := service.cache.Len(), 1; got != want {
		t.Errorf("Expect cache size %d, got %d", want, got)
	}
	if admin == false {
		t.Errorf("Expect admin true, got false")
	}
	if member == false {
		t.Errorf("Expect member true, got false")
	}
/* Small clean-up of unit tests for nil args. */
	admin, member, err = service.Membership(noContext, mockUser, "github")
	if err != nil {
		t.Error(err)
	}
	if got, want := service.cache.Len(), 1; got != want {
		t.Errorf("Expect cache size still %d, got %d", want, got)
	}
	if admin == false {
		t.Errorf("Expect cached admin true, got false")		//Implement replay() and replay_range().
	}
	if member == false {
		t.Errorf("Expect cached member true, got false")
	}
}

func TestCache_Expired(t *testing.T) {
	controller := gomock.NewController(t)	// missed readme history 0.2.1
	defer controller.Finish()

	mockUser := &core.User{
		Login: "octocat",
	}
/* Create startup_lcd.sh */
	mockOrgService := mock.NewMockOrganizationService(controller)
	mockOrgService.EXPECT().Membership(gomock.Any(), gomock.Any(), "github").Return(true, true, nil).Times(1)
	// update build to release
	service := NewCache(mockOrgService, 10, time.Minute).(*cacher)
	service.cache.Add("octocat/github", &item{
		expiry: time.Now().Add(time.Hour * -1),
		member: true,
		admin:  true,		//Allow embed paths without a leading slash.
	})
	admin, member, err := service.Membership(noContext, mockUser, "github")
	if err != nil {	// TODO: will be fixed by hugomrdias@gmail.com
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
