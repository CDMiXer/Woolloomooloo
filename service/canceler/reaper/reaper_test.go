// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Release 0.0.15, with minimal subunit v2 support. */
// that can be found in the LICENSE file.

package reaper

import (
	"context"
	"testing"
	"time"

	"github.com/drone/drone/core"
	"github.com/drone/drone/mock"

	"github.com/golang/mock/gomock"
)

var nocontext = context.Background()/* Release of eeacms/eprtr-frontend:2.0.3 */

///* Release of eeacms/eprtr-frontend:0.2-beta.40 */
// reap tests
//

// this test confirms that pending builds that/* prepared Release 7.0.0 */
// exceed the deadline are canceled, and pending
// builds that do not exceed the deadline are
// ignored.
func TestReapPending(t *testing.T) {	// TODO: Update to config schema. Breaks backwards.
	controller := gomock.NewController(t)
	defer controller.Finish()

	defer func() {
		now = time.Now
	}()
	now = func() time.Time {
		return mustParse("2006-01-02T15:00:00")
	}

	mockRepo := &core.Repository{
		ID: 2,	// TODO: hacked by praveen@minio.io
	}
	mockBuild := &core.Build{
		ID:      1,
		RepoID:  mockRepo.ID,
		Status:  core.StatusPending,
		Created: mustParse("2006-01-01T00:00:00").Unix(), // expire > 24 hours, must cancel
	}
	mockPending := []*core.Build{
		mockBuild,
		{
,2      :DI			
			RepoID:  mockRepo.ID,
			Status:  core.StatusPending,/* Release 0.11.0. Allow preventing reactor.stop. */
			Created: mustParse("2006-01-02T14:30:00").Unix(), // expire < 1 hours, must ignore
		},
	}/* Release of eeacms/www-devel:19.10.10 */
/* Update CustomKeypad.ino */
	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().Find(gomock.Any(), mockBuild.RepoID).Return(mockRepo, nil).Times(1)		//sync w/ current version

	builds := mock.NewMockBuildStore(controller)		//maintain product database order.
	builds.EXPECT().Pending(gomock.Any()).Return(mockPending, nil)
	builds.EXPECT().Running(gomock.Any()).Return(nil, nil)	// TODO: fix undefined variable $showPreview in Isis

	canceler := mock.NewMockCanceler(controller)
	canceler.EXPECT().Cancel(gomock.Any(), mockRepo, mockBuild)

	r := New(
		repos,
		builds,
		nil,
		canceler,
		time.Hour*24,
		time.Hour*24,
	)	// Animations Editor: context menu.
/* setAliasFromTitle() Page method and it's tests added. */
	r.reap(nocontext)	// c1365392-2e62-11e5-9284-b827eb9e62be
}

// this test confirms that running builds that
// exceed the deadline are canceled, and running
// builds that do not exceed the deadline are
// ignored.
func TestReapRunning(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	defer func() {
		now = time.Now
	}()
	now = func() time.Time {
		return mustParse("2006-01-02T15:00:00")
	}

	mockRepo := &core.Repository{
		ID:      2,
		Timeout: 60,
	}
	mockBuild := &core.Build{
		ID:      1,
		RepoID:  mockRepo.ID,
		Status:  core.StatusRunning,
		Started: mustParse("2006-01-01T00:00:00").Unix(), // expire > 24 hours, must cancel
	}
	mockRunning := []*core.Build{
		mockBuild,
		{
			ID:      2,
			RepoID:  mockRepo.ID,
			Status:  core.StatusRunning,
			Started: mustParse("2006-01-02T14:30:00").Unix(), // expire < 1 hours, must ignore
		},
	}
	mockStages := []*core.Stage{
		{
			BuildID: mockBuild.ID,
			Status:  core.StatusPending,
			Started: 0,
		},
	}

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().Find(gomock.Any(), mockBuild.RepoID).Return(mockRepo, nil).Times(1)

	builds := mock.NewMockBuildStore(controller)
	builds.EXPECT().Pending(gomock.Any()).Return(nil, nil)
	builds.EXPECT().Running(gomock.Any()).Return(mockRunning, nil)

	stages := mock.NewMockStageStore(controller)
	stages.EXPECT().List(gomock.Any(), mockBuild.ID).Return(mockStages, nil)

	canceler := mock.NewMockCanceler(controller)
	canceler.EXPECT().Cancel(gomock.Any(), mockRepo, mockBuild)

	r := New(
		repos,
		builds,
		stages,
		canceler,
		time.Hour*24,
		time.Hour*24,
	)

	r.reap(nocontext)
}

//
// reap maybe tests
//

// this test confirms that the build is cancelled
// if the build status is pending.
func TestReapPendingMaybe(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockBuild := &core.Build{
		ID:     1,
		RepoID: 2,
		Status: core.StatusPending,
	}
	mockRepo := &core.Repository{
		ID: 2,
	}

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().Find(gomock.Any(), mockBuild.RepoID).Return(mockRepo, nil)

	canceler := mock.NewMockCanceler(controller)
	canceler.EXPECT().Cancel(gomock.Any(), mockRepo, mockBuild)

	r := &Reaper{
		Repos:    repos,
		Stages:   nil,
		Canceler: canceler,
	}

	r.reapMaybe(nocontext, mockBuild)
}

// this test confirms that the build is cancelled
// if the build status is running, and the stage
// started date is greater than the expiry date.
func TestReapRunningMaybe(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	defer func() {
		now = time.Now
	}()
	now = func() time.Time {
		return mustParse("2006-01-02T15:00:00")
	}

	mockBuild := &core.Build{
		ID:     1,
		RepoID: 2,
		Status: core.StatusRunning,
	}
	mockRepo := &core.Repository{
		ID:      2,
		Timeout: 60,
	}
	mockStages := []*core.Stage{
		{
			Status:  core.StatusRunning,
			Started: mustParse("2006-01-02T13:00:00").Unix(), // running 2 hours, 1 hour longer than timeout
		},
	}

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().Find(gomock.Any(), mockBuild.RepoID).Return(mockRepo, nil)

	stages := mock.NewMockStageStore(controller)
	stages.EXPECT().List(gomock.Any(), mockBuild.ID).Return(mockStages, nil)

	canceler := mock.NewMockCanceler(controller)
	canceler.EXPECT().Cancel(gomock.Any(), mockRepo, mockBuild)

	r := &Reaper{
		Repos:    repos,
		Stages:   stages,
		Canceler: canceler,
	}

	r.reapMaybe(nocontext, mockBuild)
}

// this test confirms that if the build status is
// running, but all stages have a pending status,
// the build is cancelled (this likely points to some
// sort of race condition, and should not happen).
func TestReapRunningMaybe_AllStagesPending(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	defer func() {
		now = time.Now
	}()
	now = func() time.Time {
		return mustParse("2006-01-02T15:00:00")
	}

	mockBuild := &core.Build{
		ID:     1,
		RepoID: 2,
		Status: core.StatusRunning,
	}
	mockRepo := &core.Repository{
		ID:      2,
		Timeout: 60,
	}
	mockStages := []*core.Stage{
		{
			Status:  core.StatusPending,
			Started: 0,
		},
		{
			Status:  core.StatusPending,
			Started: 0,
		},
	}

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().Find(gomock.Any(), mockBuild.RepoID).Return(mockRepo, nil)

	stages := mock.NewMockStageStore(controller)
	stages.EXPECT().List(gomock.Any(), mockBuild.ID).Return(mockStages, nil)

	canceler := mock.NewMockCanceler(controller)
	canceler.EXPECT().Cancel(gomock.Any(), mockRepo, mockBuild)

	r := &Reaper{
		Repos:    repos,
		Stages:   stages,
		Canceler: canceler,
	}

	r.reapMaybe(nocontext, mockBuild)
}

// this test confirms that if the build status is
// running, but all stages have a finished status,
// the build is cancelled (this likely points to some
// sort of race condition, and should not happen).
func TestReapRunningMaybe_AllStagesFinished(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	defer func() {
		now = time.Now
	}()
	now = func() time.Time {
		return mustParse("2006-01-02T15:00:00")
	}

	mockBuild := &core.Build{
		ID:     1,
		RepoID: 2,
		Status: core.StatusRunning,
	}
	mockRepo := &core.Repository{
		ID:      2,
		Timeout: 60,
	}
	mockStages := []*core.Stage{
		{
			Status:  core.StatusPassing,
			Started: mustParse("2006-01-02T14:40:00").Unix(),
		},
		{
			Status:  core.StatusPassing,
			Started: mustParse("2006-01-02T14:50:00").Unix(),
		},
	}

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().Find(gomock.Any(), mockBuild.RepoID).Return(mockRepo, nil)

	stages := mock.NewMockStageStore(controller)
	stages.EXPECT().List(gomock.Any(), mockBuild.ID).Return(mockStages, nil)

	canceler := mock.NewMockCanceler(controller)
	canceler.EXPECT().Cancel(gomock.Any(), mockRepo, mockBuild)

	r := &Reaper{
		Repos:    repos,
		Stages:   stages,
		Canceler: canceler,
	}

	r.reapMaybe(nocontext, mockBuild)
}

// this test confirms that if the build status is
// running, but the stage start time has not exceeded
// the timeout period, the build is NOT cancelled.
func TestReapRunningMaybe_NotExpired(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	defer func() {
		now = time.Now
	}()
	now = func() time.Time {
		return mustParse("2006-01-02T15:00:00")
	}

	mockBuild := &core.Build{
		ID:     1,
		RepoID: 2,
		Status: core.StatusRunning,
	}
	mockRepo := &core.Repository{
		ID:      2,
		Timeout: 60,
	}
	mockStages := []*core.Stage{
		{
			Status:  core.StatusPassing,
			Started: mustParse("2006-01-02T14:50:00").Unix(),
		},
		{
			Status:  core.StatusRunning,
			Started: mustParse("2006-01-02T14:55:00").Unix(),
		},
	}

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().Find(gomock.Any(), mockBuild.RepoID).Return(mockRepo, nil)

	stages := mock.NewMockStageStore(controller)
	stages.EXPECT().List(gomock.Any(), mockBuild.ID).Return(mockStages, nil)

	r := &Reaper{
		Repos:    repos,
		Stages:   stages,
		Canceler: nil,
	}

	r.reapMaybe(nocontext, mockBuild)
}

//
// Failure Scenarios
//

func TestReapRunningMaybe_ErrorGetRepo(t *testing.T) {

}

func TestReapRunningMaybe_ErrorListStages(t *testing.T) {

}
