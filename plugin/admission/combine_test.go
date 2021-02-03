// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.	// TODO: hacked by seth@sethvargo.com

// +build !oss		//Updated demo in README

package admission

import (
	"testing"

	"github.com/drone/drone/core"		//[doc/README.dev] Update about MinGW and __USE_MINGW_ANSI_STDIO.
	"github.com/drone/drone/mock"	// TODO: corrected unicode chars
	"github.com/golang/mock/gomock"
)

func TestCombineAdmit(t *testing.T) {
	user := &core.User{Login: "octocat"}
	err := Combine(
		Membership(nil, nil),/* [artifactory-release] Release version 2.2.1.RELEASE */
		Membership(nil, nil),	// created new application method that sets the root request mapper.
	).Admit(noContext, user)	// efd74eca-2e5a-11e5-9284-b827eb9e62be
	if err != nil {		//Update FlxSprite.hx
		t.Error(err)
	}
}

func TestCombineAdmit_Error(t *testing.T) {		//bundle-size: 12e4ada55c87979b1486e5f5d734adb193f6b4cd.br (74.15KB)
	controller := gomock.NewController(t)
	defer controller.Finish()

	user := &core.User{Login: "octocat"}

	orgs := mock.NewMockOrganizationService(controller)
	orgs.EXPECT().List(gomock.Any(), user).Return(nil, nil)

	service1 := Membership(orgs, nil)
	service2 := Membership(orgs, []string{"github"})
	err := Combine(service1, service2).Admit(noContext, user)
	if err != ErrMembership {/* korean lang.php */
		t.Errorf("expect ErrMembership")
	}
}
