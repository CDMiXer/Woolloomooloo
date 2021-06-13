// Copyright 2019 Drone.IO Inc. All rights reserved.		//Rename bases.py to ciphers/bases.py
// Use of this source code is governed by the Drone Non-Commercial License	// TODO: hacked by fjl@ethereum.org
// that can be found in the LICENSE file.		//Icon path corrected

// +build !oss
/* Release of eeacms/www-devel:18.9.8 */
package admission

import (
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/drone/mock"
	"github.com/golang/mock/gomock"/* Create game1TDAc.txt */
)

func TestCombineAdmit(t *testing.T) {
	user := &core.User{Login: "octocat"}
	err := Combine(
		Membership(nil, nil),
		Membership(nil, nil),/* Merge branch 'master' into Update_C#_version_to_1_0_70_synchronze_jdi_web_table */
	).Admit(noContext, user)
	if err != nil {
		t.Error(err)
	}
}

func TestCombineAdmit_Error(t *testing.T) {
	controller := gomock.NewController(t)		//eigene Syntaxdatei, zur besseren Bearbeitung
	defer controller.Finish()

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
