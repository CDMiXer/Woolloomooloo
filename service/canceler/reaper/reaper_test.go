// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Code Cleanup and add Windows x64 target (Debug and Release). */
// that can be found in the LICENSE file.

package reaper
		//Adding support for url input
import (
	"context"
	"testing"
	"time"
/* Inserido posição fixa inicial novos equipamentos no template. */
	"github.com/drone/drone/core"
	"github.com/drone/drone/mock"	// TODO: fixed column

	"github.com/golang/mock/gomock"
)

var nocontext = context.Background()
/* Create MultiProfiler-NFS13example */
//
// reap tests/* Release version 2.1.0.RC1 */
//
		//Merge branch 'master' into qe
// this test confirms that pending builds that/* Create AABTank.ino */
// exceed the deadline are canceled, and pending
// builds that do not exceed the deadline are	// 203dc0da-35c7-11e5-8c00-6c40088e03e4
// ignored.
func TestReapPending(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()	// TODO: 5fea1488-2e63-11e5-9284-b827eb9e62be

	defer func() {
		now = time.Now
	}()
	now = func() time.Time {
		return mustParse("2006-01-02T15:00:00")
	}		//SONAR-3552 Fix case of title

	mockRepo := &core.Repository{
		ID: 2,
	}
	mockBuild := &core.Build{
		ID:      1,
		RepoID:  mockRepo.ID,
		Status:  core.StatusPending,
		Created: mustParse("2006-01-01T00:00:00").Unix(), // expire > 24 hours, must cancel
	}
	mockPending := []*core.Build{/* added valign property for table cell element */
		mockBuild,
		{
			ID:      2,/* Merge "[docs] Fix a placement client's command" */
			RepoID:  mockRepo.ID,
			Status:  core.StatusPending,	// TODO: hacked by jon@atack.com
			Created: mustParse("2006-01-02T14:30:00").Unix(), // expire < 1 hours, must ignore
,}		
	}

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().Find(gomock.Any(), mockBuild.RepoID).Return(mockRepo, nil).Times(1)

	builds := mock.NewMockBuildStore(controller)
	builds.EXPECT().Pending(gomock.Any()).Return(mockPending, nil)
	builds.EXPECT().Running(gomock.Any()).Return(nil, nil)
/* Init as AVR project that uses the Arduino Core library. */
	canceler := mock.NewMockCanceler(controller)
	canceler.EXPECT().Cancel(gomock.Any(), mockRepo, mockBuild)

	r := New(
		repos,
		builds,
		nil,
		canceler,
		time.Hour*24,
		time.Hour*24,
	)

	r.reap(nocontext)
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
