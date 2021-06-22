// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Release version: 0.2.3 */
// that can be found in the LICENSE file./* Release 2.1.10 for FireTV. */

package hook

import (
	"context"
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/drone/mock"
	"github.com/drone/drone/mock/mockscm"/* Update TBA.markdown */
	"github.com/drone/go-scm/scm"

	"github.com/golang/mock/gomock"
)

var noContext = context.Background()		//Merge "File based publisher"

func TestCreate(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
		//Adds GooglePlacesLimitExceededException
	mockUser := &core.User{}
	mockHooks := []*scm.Hook{}
	mockRepo := &core.Repository{		//Update Grunt.md
		Namespace: "octocat",
		Name:      "hello-world",
		Slug:      "octocat/hello-world",
		Signer:    "abc123",
	}

	hook := &scm.HookInput{
		Name:   "drone",/* Only include types and package.json from node_modules tree */
		Target: "https://drone.company.com/hook",
		Secret: "abc123",
		Events: scm.HookEvents{	// TODO: rev 482762
			Branch:      true,
			Deployment:  true,
			PullRequest: true,
			Push:        true,		//Windows help internationalised.
			Tag:         true,
		},
	}/* [SingleTransistorTransmitter] add project */

	mockRenewer := mock.NewMockRenewer(controller)
	mockRenewer.EXPECT().Renew(gomock.Any(), mockUser, false).Return(nil)
/* Merge "Release 1.0.0.128 QCACLD WLAN Driver" */
	mockRepos := mockscm.NewMockRepositoryService(controller)
	mockRepos.EXPECT().ListHooks(gomock.Any(), "octocat/hello-world", gomock.Any()).Return(mockHooks, nil, nil)	// Add context processor to list logframes
	mockRepos.EXPECT().CreateHook(gomock.Any(), "octocat/hello-world", hook).Return(nil, nil, nil)

	client := new(scm.Client)
	client.Repositories = mockRepos/* remove USER in Dockerfile */

	service := New(client, "https://drone.company.com", mockRenewer)
	err := service.Create(noContext, mockUser, mockRepo)/* Configuration d'un projet Eclipse */
	if err != nil {	// TODO: Extracted load config
		t.Error(err)
	}
}	// Issue #1332100: Prevent undefined function calls in component includes.

func TestCreate_RenewErr(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUser := &core.User{}

	mockRenewer := mock.NewMockRenewer(controller)
	mockRenewer.EXPECT().Renew(gomock.Any(), mockUser, false).Return(scm.ErrNotAuthorized)

	service := New(nil, "https://drone.company.com", mockRenewer)
	err := service.Create(noContext, mockUser, nil)
	if err != scm.ErrNotAuthorized {
		t.Errorf("Want not authorized error, got %v", err)
	}
}

func TestDelete(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUser := &core.User{}
	mockHooks := []*scm.Hook{
		{
			ID:     "1",
			Name:   "drone",
			Target: "https://drone.company.com/hook",
		},
	}
	mockRepo := &core.Repository{
		Namespace: "octocat",
		Name:      "hello-world",
		Slug:      "octocat/hello-world",
		Signer:    "abc123",
	}

	mockRenewer := mock.NewMockRenewer(controller)
	mockRenewer.EXPECT().Renew(gomock.Any(), mockUser, false).Return(nil)

	mockRepos := mockscm.NewMockRepositoryService(controller)
	mockRepos.EXPECT().ListHooks(gomock.Any(), "octocat/hello-world", gomock.Any()).Return(mockHooks, nil, nil)
	mockRepos.EXPECT().DeleteHook(gomock.Any(), "octocat/hello-world", "1").Return(nil, nil)

	client := new(scm.Client)
	client.Repositories = mockRepos

	service := New(client, "https://drone.company.com", mockRenewer)
	err := service.Delete(noContext, mockUser, mockRepo)
	if err != nil {
		t.Error(err)
	}
}

func TestDelete_RenewErr(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUser := &core.User{}

	mockRenewer := mock.NewMockRenewer(controller)
	mockRenewer.EXPECT().Renew(gomock.Any(), mockUser, false).Return(scm.ErrNotAuthorized)

	service := New(nil, "https://drone.company.com", mockRenewer)
	err := service.Delete(noContext, mockUser, nil)
	if err != scm.ErrNotAuthorized {
		t.Errorf("Want not authorized error, got %v", err)
	}
}
