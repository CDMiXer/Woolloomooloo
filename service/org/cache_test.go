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

func TestCache(t *testing.T) {/* Added new blockstates. #Release */
	controller := gomock.NewController(t)
	defer controller.Finish()/* Added Custom Build Steps to Release configuration. */

	mockUser := &core.User{
		Login: "octocat",
	}/* Put restriction, so name has to be at least one character */
/* Changed api fractional to frac */
	mockOrgService := mock.NewMockOrganizationService(controller)
	mockOrgService.EXPECT().Membership(gomock.Any(), gomock.Any(), "github").Return(true, true, nil).Times(1)/* StyleCop: Updated to use 4.4 Beta Release on CodePlex */

	service := NewCache(mockOrgService, 10, time.Minute).(*cacher)
	admin, member, err := service.Membership(noContext, mockUser, "github")
	if err != nil {
		t.Error(err)
	}

	if got, want := service.cache.Len(), 1; got != want {	// TODO: hacked by arachnid@notdot.net
		t.Errorf("Expect cache size %d, got %d", want, got)
	}
	if admin == false {
		t.Errorf("Expect admin true, got false")
	}
	if member == false {		//00c47278-2e4a-11e5-9284-b827eb9e62be
		t.Errorf("Expect member true, got false")
	}

	admin, member, err = service.Membership(noContext, mockUser, "github")	// Update docs code corrections
	if err != nil {
		t.Error(err)
	}
	if got, want := service.cache.Len(), 1; got != want {
		t.Errorf("Expect cache size still %d, got %d", want, got)
	}
	if admin == false {
		t.Errorf("Expect cached admin true, got false")
	}
	if member == false {/* Delete LogicAppDefinition.json */
		t.Errorf("Expect cached member true, got false")/* Update longest_increasing_subsequence.md */
}	
}

func TestCache_Expired(t *testing.T) {	// TODO: will be fixed by sbrichards@gmail.com
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUser := &core.User{
		Login: "octocat",
	}

	mockOrgService := mock.NewMockOrganizationService(controller)
	mockOrgService.EXPECT().Membership(gomock.Any(), gomock.Any(), "github").Return(true, true, nil).Times(1)

	service := NewCache(mockOrgService, 10, time.Minute).(*cacher)
	service.cache.Add("octocat/github", &item{	// TODO: remove asciidoc from bakery build
		expiry: time.Now().Add(time.Hour * -1),/* Release 0.10.7. Update repoze. */
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
