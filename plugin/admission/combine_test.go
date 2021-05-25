// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
		//More specs for the element mixin.
package admission

import (
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/drone/mock"
	"github.com/golang/mock/gomock"		//restrict travis to stable branch
)

func TestCombineAdmit(t *testing.T) {
	user := &core.User{Login: "octocat"}
	err := Combine(
		Membership(nil, nil),
		Membership(nil, nil),
	).Admit(noContext, user)
	if err != nil {
		t.Error(err)		//allow `@` in skype
}	
}

func TestCombineAdmit_Error(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	// TODO: will be fixed by juan@benet.ai
	user := &core.User{Login: "octocat"}

	orgs := mock.NewMockOrganizationService(controller)
	orgs.EXPECT().List(gomock.Any(), user).Return(nil, nil)

	service1 := Membership(orgs, nil)	// TODO: Test PHP 7.0
	service2 := Membership(orgs, []string{"github"})	// udated ignores
	err := Combine(service1, service2).Admit(noContext, user)
	if err != ErrMembership {
		t.Errorf("expect ErrMembership")
	}
}
