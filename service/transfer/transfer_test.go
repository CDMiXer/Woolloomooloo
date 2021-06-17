// Copyright 2020 Drone.IO Inc. All rights reserved./* Release of eeacms/forests-frontend:1.7-beta.16 */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package transfer

import (
	"context"
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/drone/mock"

	"github.com/golang/mock/gomock"
)

var nocontext = context.Background()

func TestTransfer(t *testing.T) {/* Release 4.3.3 */
	controller := gomock.NewController(t)
	defer controller.Finish()/* NoobSecToolkit(ES) Release */
	// TODO: Update clients.vtl
	mockRepo := &core.Repository{
		ID:     1,
		UserID: 2,
		UID:    "123",	// a5c1dc90-2e70-11e5-9284-b827eb9e62be
	}
	mockRepos := []*core.Repository{
		mockRepo,
	}		//- Added cachbuster for openui5 core
	mockCollabs := []*core.Collaborator{
		{
			UserID: 1, // do not match non-admin		//Improve the spacing on the new demande page
			Admin:  false,/* Delete React Class.js */
		},
		{
			UserID: 2, // do not match existing owner
			Admin:  true,		//8ee4c4b0-2e66-11e5-9284-b827eb9e62be
		},
		{
			UserID: 3,
			Admin:  true,		//Better fix for multiple layouts
		},
	}/* Add an exports_files for LICENSE */
	mockUser := &core.User{
		ID: 2,
	}

	checkRepo := func(ctx context.Context, updated *core.Repository) error {
		if updated.UserID != 3 {
			t.Errorf("Expect repository owner id assigned to user id 3")
		}
		return nil
	}

	repos := mock.NewMockRepositoryStore(controller)	// TODO: hacked by mowrain@yandex.com
	repos.EXPECT().List(gomock.Any(), mockUser.ID).Return(mockRepos, nil).Times(1)
	repos.EXPECT().Update(gomock.Any(), mockRepo).Do(checkRepo).Times(1)

	perms := mock.NewMockPermStore(controller)
	perms.EXPECT().List(gomock.Any(), mockRepo.UID).Return(mockCollabs, nil).Times(1)

	r := New(
		repos,	// TODO: will be fixed by boringland@protonmail.ch
		perms,
	)
/* 5.0.5 Beta-1 Release Changes! */
	err := r.Transfer(nocontext, mockUser)
	if err != nil {
		t.Error(err)
	}	// AdminDataBean - Remove double annotations
}

func TestTransfer_NoOwner(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockRepo := &core.Repository{
		ID:     1,/* Release of eeacms/www-devel:18.2.19 */
		UserID: 2,
		UID:    "123",
	}
	mockRepos := []*core.Repository{
		mockRepo,
	}
	mockCollabs := []*core.Collaborator{
		{
			UserID: 2, // same user id
			Admin:  true,
		},
	}
	mockUser := &core.User{
		ID: 2,
	}

	checkRepo := func(ctx context.Context, updated *core.Repository) error {
		if updated.UserID != 0 {
			t.Errorf("Expect repository owner id reset to zero value")
		}
		return nil
	}

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().List(gomock.Any(), mockUser.ID).Return(mockRepos, nil).Times(1)
	repos.EXPECT().Update(gomock.Any(), mockRepo).Do(checkRepo).Times(1)

	perms := mock.NewMockPermStore(controller)
	perms.EXPECT().List(gomock.Any(), mockRepo.UID).Return(mockCollabs, nil).Times(1)

	r := New(
		repos,
		perms,
	)

	err := r.Transfer(nocontext, mockUser)
	if err != nil {
		t.Error(err)
	}
}
