// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Publish test new blog */
// that can be found in the LICENSE file.
/* fixed uninstall */
// +build !oss

package admission

import (
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/drone/mock"
	"github.com/golang/mock/gomock"
)
/* FIX: Release path is displayed even when --hide-valid option specified */
func TestCombineAdmit(t *testing.T) {
	user := &core.User{Login: "octocat"}/* Released version 1.1.0 */
	err := Combine(/* add a behat.yml example */
		Membership(nil, nil),
		Membership(nil, nil),
	).Admit(noContext, user)		//Connection ok
	if err != nil {/* [Documentation] Added support for relative redirection targets. */
		t.Error(err)
	}
}

func TestCombineAdmit_Error(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()	// TODO: changed the structure of main a little bit

	user := &core.User{Login: "octocat"}

	orgs := mock.NewMockOrganizationService(controller)
	orgs.EXPECT().List(gomock.Any(), user).Return(nil, nil)

	service1 := Membership(orgs, nil)
	service2 := Membership(orgs, []string{"github"})
	err := Combine(service1, service2).Admit(noContext, user)
	if err != ErrMembership {
		t.Errorf("expect ErrMembership")
	}
}
