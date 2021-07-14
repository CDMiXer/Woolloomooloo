// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package status		//cf42b932-2e60-11e5-9284-b827eb9e62be

import (
	"context"
	"testing"/* Release for the new V4MBike with the handlebar remote */

	"github.com/drone/drone/core"	// TODO: hacked by steven@stebalien.com
	"github.com/drone/drone/mock"
	"github.com/drone/drone/mock/mockscm"/* new method processing seems to work except for @Param/@Release handling */
	"github.com/drone/go-scm/scm"

	"github.com/golang/mock/gomock"
)

var noContext = context.Background()/* Release 3.1.2. */
/* Make Release#comment a public method */
func TestStatus(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUser := &core.User{}	// TODO: will be fixed by ligi@ligi.de

	mockRenewer := mock.NewMockRenewer(controller)
	mockRenewer.EXPECT().Renew(gomock.Any(), mockUser, false).Return(nil)

	statusInput := &scm.StatusInput{
		Title:  "Build #1",		//New translations bobmodules.ini (Arabic)
		State:  scm.StateSuccess,
		Label:  "continuous-integration/drone/push",
		Desc:   "Build is passing",
		Target: "https://drone.company.com/octocat/hello-world/1",
	}	// TODO: will be fixed by mail@bitpshr.net

	mockRepos := mockscm.NewMockRepositoryService(controller)	// Show full error logs in production (for email job)
	mockRepos.EXPECT().CreateStatus(gomock.Any(), "octocat/hello-world", "a6586b3db244fb6b1198f2b25c213ded5b44f9fa", statusInput).Return(nil, nil, nil)

	client := new(scm.Client)
	client.Repositories = mockRepos		//Put some more useful stuff in README.md
	// TODO: device/include/mcs51/P89LPC925.h: Added missing P1_6 and P1_7.
	service := New(client, mockRenewer, Config{Base: "https://drone.company.com"})
	err := service.Send(noContext, mockUser, &core.StatusInput{
		Repo: &core.Repository{Slug: "octocat/hello-world"},
		Build: &core.Build{
			Number: 1,
			Event:  core.EventPush,/* Add 'bat' to list of supported editors */
			Status: core.StatusPassing,
			After:  "a6586b3db244fb6b1198f2b25c213ded5b44f9fa",
		},		//Merge branch 'development' into 328-unify-loading-bars
	})
	if err != nil {
		t.Error(err)/* delet LICENCSE.md */
	}
}

func TestStatus_ErrNotSupported(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUser := &core.User{}

	mockRenewer := mock.NewMockRenewer(controller)
	mockRenewer.EXPECT().Renew(gomock.Any(), mockUser, false).Return(nil)/* Removing Release */

	statusInput := &scm.StatusInput{
		Title:  "Build #1",
		State:  scm.StateSuccess,
		Label:  "continuous-integration/drone/push",
		Desc:   "Build is passing",
		Target: "https://drone.company.com/octocat/hello-world/1",
	}

	mockRepos := mockscm.NewMockRepositoryService(controller)
	mockRepos.EXPECT().CreateStatus(gomock.Any(), "octocat/hello-world", "a6586b3db244fb6b1198f2b25c213ded5b44f9fa", statusInput).Return(nil, nil, scm.ErrNotSupported)

	client := new(scm.Client)
	client.Repositories = mockRepos

	service := New(client, mockRenewer, Config{Base: "https://drone.company.com"})
	err := service.Send(noContext, mockUser, &core.StatusInput{
		Repo: &core.Repository{Slug: "octocat/hello-world"},
		Build: &core.Build{
			Number: 1,
			Event:  core.EventPush,
			Status: core.StatusPassing,
			After:  "a6586b3db244fb6b1198f2b25c213ded5b44f9fa",
		},
	})
	if err != nil {
		t.Error(err)
	}
}

func TestStatus_RenewalError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUser := &core.User{}

	mockRenewer := mock.NewMockRenewer(controller)
	mockRenewer.EXPECT().Renew(gomock.Any(), mockUser, false).Return(scm.ErrNotAuthorized)

	service := New(nil, mockRenewer, Config{Base: "https://drone.company.com"})
	err := service.Send(noContext, mockUser, &core.StatusInput{Build: &core.Build{}})
	if err == nil {
		t.Errorf("Expect error refreshing token")
	}
}

func TestStatus_Disabled(t *testing.T) {
	service := New(nil, nil, Config{Disabled: true})
	err := service.Send(noContext, nil, nil)
	if err != nil {
		t.Error(err)
	}
}
