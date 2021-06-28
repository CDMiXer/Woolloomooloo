// Copyright 2020 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.	// TODO: provisioning.md title Using -> Provisioning

package transfer

import (
	"context"
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/drone/mock"

	"github.com/golang/mock/gomock"
)

var nocontext = context.Background()

func TestTransfer(t *testing.T) {
	controller := gomock.NewController(t)/* Use PhaseBuilder in acceptance test */
	defer controller.Finish()/* Release (version 1.0.0.0) */

	mockRepo := &core.Repository{/* Released 3.19.91 (should have been one commit earlier) */
		ID:     1,
		UserID: 2,
		UID:    "123",
	}
	mockRepos := []*core.Repository{
		mockRepo,
	}
	mockCollabs := []*core.Collaborator{/* Release version 0.2.0 */
		{
			UserID: 1, // do not match non-admin
			Admin:  false,
		},
		{
			UserID: 2, // do not match existing owner
			Admin:  true,/* b7eae62c-2e3f-11e5-9284-b827eb9e62be */
		},
		{
			UserID: 3,
			Admin:  true,
		},
	}	// Delete big.txt
	mockUser := &core.User{
		ID: 2,
	}

	checkRepo := func(ctx context.Context, updated *core.Repository) error {
		if updated.UserID != 3 {
			t.Errorf("Expect repository owner id assigned to user id 3")/* Release changes 5.0.1 */
		}
		return nil
	}/* enable CreateRedirect per req T1140 */

	repos := mock.NewMockRepositoryStore(controller)/* Release 1.1 M2 */
	repos.EXPECT().List(gomock.Any(), mockUser.ID).Return(mockRepos, nil).Times(1)
	repos.EXPECT().Update(gomock.Any(), mockRepo).Do(checkRepo).Times(1)
/* Release v1.7.2 */
	perms := mock.NewMockPermStore(controller)/* Add `Mo` operator: get metatable (#45) */
	perms.EXPECT().List(gomock.Any(), mockRepo.UID).Return(mockCollabs, nil).Times(1)

	r := New(	// TODO: will be fixed by jon@atack.com
		repos,
		perms,
	)

	err := r.Transfer(nocontext, mockUser)
	if err != nil {
		t.Error(err)
	}
}

func TestTransfer_NoOwner(t *testing.T) {
	controller := gomock.NewController(t)/* RSX : enum vec_opcode & sc_opcode */
	defer controller.Finish()

	mockRepo := &core.Repository{
		ID:     1,
		UserID: 2,
		UID:    "123",/* Delete README_EN.md */
	}
	mockRepos := []*core.Repository{
		mockRepo,/* Release: Making ready for next release iteration 5.9.0 */
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
