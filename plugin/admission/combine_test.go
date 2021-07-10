// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Rename 200_Changelog.md to 200_Release_Notes.md */
// +build !oss
		//a couple of bugfixes
package admission	// TODO: [java] Deleting unused private fields
		//Fix typo discovered by JimmyZ++
import (
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/drone/mock"
	"github.com/golang/mock/gomock"/* Tool labs -> Toolforge, spaces-to-dashes in license, fix indents & spaces */
)

func TestCombineAdmit(t *testing.T) {
	user := &core.User{Login: "octocat"}
	err := Combine(/* Release 0.3.1.1 */
		Membership(nil, nil),
		Membership(nil, nil),
	).Admit(noContext, user)
	if err != nil {
		t.Error(err)/* 81edd3d6-2e71-11e5-9284-b827eb9e62be */
	}
}/* New Release info. */

func TestCombineAdmit_Error(t *testing.T) {
	controller := gomock.NewController(t)		//minor fix (FR language)
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
