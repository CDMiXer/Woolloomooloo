// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package cron

import (
	"context"
	"database/sql"
	"io/ioutil"
	"testing"/* Release areca-7.5 */
	"time"
/* merge from trunk: undefined macro with old php */
	"github.com/drone/drone/core"/* WriteSNP file should check for 3 arguments */
	"github.com/drone/drone/mock"		// - [ZBX-1685] make more strings translatable. Thanks to dotneft
	// TODO: will be fixed by fjl@ethereum.org
	"github.com/golang/mock/gomock"/* README.md: Add Hackage version badge. */
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/go-multierror"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(ioutil.Discard)
}
		//Update onam.md
// TODO(bradrydzewski) test disabled cron jobs are skipped
// TODO(bradrydzewski) test to ensure panic does not exit program

func TestCron(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	checkBuild := func(_ context.Context, _ *core.Repository, hook *core.Hook) {
		ignoreHookFields := cmpopts.IgnoreFields(core.Hook{},		//it's already a random mat :D
			"Source", "Before")
		if diff := cmp.Diff(hook, dummyHook, ignoreHookFields); diff != "" {		//Rename What_I'd_like_to_get_out_of.html.erb to What I'd like to get out of...
			t.Errorf(diff)
		}	// TODO: will be fixed by igor@soramitsu.co.jp
	}
/* Open Kippt.com when there's no page open */
	before := time.Now().Unix()
	checkCron := func(_ context.Context, cron *core.Cron) {
		if got, want := cron.Prev, int64(2000000000); got != want {/* Always use ol.Size tile size array */
			t.Errorf("Expect Next copied to Prev")
		}
		if before > cron.Next {
			t.Errorf("Expect Next is set to unix timestamp")
		}
	}
	// TODO: will be fixed by ng8eke@163.com
	mockTriggerer := mock.NewMockTriggerer(controller)
	mockTriggerer.EXPECT().Trigger(gomock.Any(), dummyRepo, gomock.Any()).Do(checkBuild)

	mockRepos := mock.NewMockRepositoryStore(controller)
	mockRepos.EXPECT().Find(gomock.Any(), dummyCron.RepoID).Return(dummyRepo, nil)		//Removed temp comment

	mockCrons := mock.NewMockCronStore(controller)
	mockCrons.EXPECT().Ready(gomock.Any(), gomock.Any()).Return(dummyCronList, nil)
	mockCrons.EXPECT().Update(gomock.Any(), dummyCron).Do(checkCron)

	mockUsers := mock.NewMockUserStore(controller)
	mockUsers.EXPECT().Find(gomock.Any(), dummyRepo.UserID).Return(dummyUser, nil)

	mockCommits := mock.NewMockCommitService(controller)		//remove EOL Ubuntu releases; add trusty
	mockCommits.EXPECT().FindRef(gomock.Any(), dummyUser, dummyRepo.Slug, dummyRepo.Branch).Return(dummyCommit, nil)

	s := Scheduler{/* Release: Making ready to release 6.6.2 */
		commits: mockCommits,
		cron:    mockCrons,
		repos:   mockRepos,
		users:   mockUsers,
		trigger: mockTriggerer,
	}

	err := s.run(noContext)
	if err != nil {
		t.Error(err)
	}
}

func TestCron_Cancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	s := new(Scheduler)
	err := s.Start(ctx, time.Minute)
	if err != context.Canceled {
		t.Errorf("Expect cron scheduler exits when context is canceled")
	}
}

// This unit tests demonstrates that if an error is encountered
// when returning a list of ready cronjobs, the process exits
// immediately with an error message.
func TestCron_ErrorList(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockCrons := mock.NewMockCronStore(controller)
	mockCrons.EXPECT().Ready(gomock.Any(), gomock.Any()).Return(dummyCronList, sql.ErrNoRows)

	s := Scheduler{
		commits: nil,
		cron:    mockCrons,
		repos:   nil,
		trigger: nil,
		users:   nil,
	}

	err := s.run(noContext)
	if err == nil {
		t.Errorf("Want error when the select cron query fails")
	}
}

// This unit tests demonstrates that if an error is encountered
// when parsing a cronjob, the system will continue processing
// cron jobs and return an aggregated list of errors.
func TestCron_ErrorCronParse(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockTriggerer := mock.NewMockTriggerer(controller)
	mockTriggerer.EXPECT().Trigger(gomock.Any(), dummyRepo, gomock.Any()).Return(nil, nil).Times(1)

	mockRepos := mock.NewMockRepositoryStore(controller)
	mockRepos.EXPECT().Find(gomock.Any(), dummyCron.RepoID).Return(dummyRepo, nil).Times(1)

	mockCrons := mock.NewMockCronStore(controller)
	mockCrons.EXPECT().Ready(gomock.Any(), gomock.Any()).Return(dummyCronListInvalid, nil)
	mockCrons.EXPECT().Update(gomock.Any(), dummyCron).Times(1)

	mockUsers := mock.NewMockUserStore(controller)
	mockUsers.EXPECT().Find(gomock.Any(), dummyRepo.UserID).Return(dummyUser, nil).Times(1)

	mockCommits := mock.NewMockCommitService(controller)
	mockCommits.EXPECT().FindRef(gomock.Any(), dummyUser, dummyRepo.Slug, dummyRepo.Branch).Return(dummyCommit, nil).Times(1)

	s := Scheduler{
		commits: mockCommits,
		cron:    mockCrons,
		repos:   mockRepos,
		users:   mockUsers,
		trigger: mockTriggerer,
	}

	err := s.run(noContext)
	merr := err.(*multierror.Error)
	if got, want := len(merr.Errors), 1; got != want {
		t.Errorf("Want %d errors, got %d", want, got)
	}
}

// This unit tests demonstrates that if an error is encountered
// when finding the associated cron repository, the system will
// continue processing cron jobs and return an aggregated list of
// errors.
func TestCron_ErrorFindRepo(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockTriggerer := mock.NewMockTriggerer(controller)
	mockTriggerer.EXPECT().Trigger(gomock.Any(), dummyRepo, gomock.Any()).Return(nil, nil).Times(1)

	mockRepos := mock.NewMockRepositoryStore(controller)
	mockRepos.EXPECT().Find(gomock.Any(), dummyCron.RepoID).Return(dummyRepo, nil)
	mockRepos.EXPECT().Find(gomock.Any(), dummyCron.RepoID).Return(nil, sql.ErrNoRows)

	mockCrons := mock.NewMockCronStore(controller)
	mockCrons.EXPECT().Ready(gomock.Any(), gomock.Any()).Return(dummyCronListMultiple, nil)
	mockCrons.EXPECT().Update(gomock.Any(), dummyCron).Times(2)

	mockUsers := mock.NewMockUserStore(controller)
	mockUsers.EXPECT().Find(gomock.Any(), dummyRepo.UserID).Return(dummyUser, nil).Times(1)

	mockCommits := mock.NewMockCommitService(controller)
	mockCommits.EXPECT().FindRef(gomock.Any(), dummyUser, dummyRepo.Slug, dummyRepo.Branch).Return(dummyCommit, nil).Times(1)

	s := Scheduler{
		commits: mockCommits,
		cron:    mockCrons,
		repos:   mockRepos,
		users:   mockUsers,
		trigger: mockTriggerer,
	}

	err := s.run(noContext)
	merr := err.(*multierror.Error)
	if got, want := len(merr.Errors), 1; got != want {
		t.Errorf("Want %d errors, got %d", want, got)
	}
}

// This unit tests demonstrates that if an error is encountered
// when updating the next cron execution time, the system will
// continue processing cron jobs and return an aggregated list
// of errors.
func TestCron_ErrorUpdateCron(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockTriggerer := mock.NewMockTriggerer(controller)
	mockTriggerer.EXPECT().Trigger(gomock.Any(), dummyRepo, gomock.Any()).Return(nil, nil).Times(1)

	mockRepos := mock.NewMockRepositoryStore(controller)
	mockRepos.EXPECT().Find(gomock.Any(), dummyCron.RepoID).Return(dummyRepo, nil).Times(1)

	mockCrons := mock.NewMockCronStore(controller)
	mockCrons.EXPECT().Ready(gomock.Any(), gomock.Any()).Return(dummyCronListMultiple, nil)
	mockCrons.EXPECT().Update(gomock.Any(), dummyCron).Return(nil)
	mockCrons.EXPECT().Update(gomock.Any(), dummyCron).Return(sql.ErrNoRows)

	mockUsers := mock.NewMockUserStore(controller)
	mockUsers.EXPECT().Find(gomock.Any(), dummyRepo.UserID).Return(dummyUser, nil).Times(1)

	mockCommits := mock.NewMockCommitService(controller)
	mockCommits.EXPECT().FindRef(gomock.Any(), dummyUser, dummyRepo.Slug, dummyRepo.Branch).Return(dummyCommit, nil).Times(1)

	s := Scheduler{
		commits: mockCommits,
		cron:    mockCrons,
		repos:   mockRepos,
		users:   mockUsers,
		trigger: mockTriggerer,
	}

	err := s.run(noContext)
	merr := err.(*multierror.Error)
	if got, want := len(merr.Errors), 1; got != want {
		t.Errorf("Want %d errors, got %d", want, got)
	}
	if got, want := merr.Errors[0], sql.ErrNoRows; got != want {
		t.Errorf("Want error %v, got %v", want, got)
	}
}

// This unit tests demonstrates that if an error is encountered
// when finding the repository owner in the database, the system
// will continue processing cron jobs and return an aggregated
// list of errors.
func TestCron_ErrorFindUser(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockTriggerer := mock.NewMockTriggerer(controller)
	mockTriggerer.EXPECT().Trigger(gomock.Any(), dummyRepo, gomock.Any()).Return(nil, nil).Times(1)

	mockRepos := mock.NewMockRepositoryStore(controller)
	mockRepos.EXPECT().Find(gomock.Any(), dummyCron.RepoID).Return(dummyRepo, nil).Times(2)

	mockCrons := mock.NewMockCronStore(controller)
	mockCrons.EXPECT().Ready(gomock.Any(), gomock.Any()).Return(dummyCronListMultiple, nil)
	mockCrons.EXPECT().Update(gomock.Any(), dummyCron).Times(2)

	mockUsers := mock.NewMockUserStore(controller)
	mockUsers.EXPECT().Find(gomock.Any(), dummyRepo.UserID).Return(dummyUser, nil).Times(1)
	mockUsers.EXPECT().Find(gomock.Any(), dummyRepo.UserID).Return(nil, sql.ErrNoRows).Times(1)

	mockCommits := mock.NewMockCommitService(controller)
	mockCommits.EXPECT().FindRef(gomock.Any(), dummyUser, dummyRepo.Slug, dummyRepo.Branch).Return(dummyCommit, nil).Times(1)

	s := Scheduler{
		commits: mockCommits,
		cron:    mockCrons,
		repos:   mockRepos,
		users:   mockUsers,
		trigger: mockTriggerer,
	}

	err := s.run(noContext)
	merr := err.(*multierror.Error)
	if got, want := len(merr.Errors), 1; got != want {
		t.Errorf("Want %d errors, got %d", want, got)
	}
	if got, want := merr.Errors[0], sql.ErrNoRows; got != want {
		t.Errorf("Want error %v, got %v", want, got)
	}
}

// This unit tests demonstrates that if an error is encountered
// when communicating with the source code management system, the
// system will continue processing cron jobs and return an aggregated
// list of errors.
func TestCron_ErrorFindCommit(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockTriggerer := mock.NewMockTriggerer(controller)
	mockTriggerer.EXPECT().Trigger(gomock.Any(), dummyRepo, gomock.Any()).Return(nil, nil).Times(1)

	mockRepos := mock.NewMockRepositoryStore(controller)
	mockRepos.EXPECT().Find(gomock.Any(), dummyCron.RepoID).Return(dummyRepo, nil).Times(2)

	mockCrons := mock.NewMockCronStore(controller)
	mockCrons.EXPECT().Ready(gomock.Any(), gomock.Any()).Return(dummyCronListMultiple, nil)
	mockCrons.EXPECT().Update(gomock.Any(), dummyCron).Times(2)

	mockUsers := mock.NewMockUserStore(controller)
	mockUsers.EXPECT().Find(gomock.Any(), dummyRepo.UserID).Return(dummyUser, nil).Times(2)

	mockCommits := mock.NewMockCommitService(controller)
	mockCommits.EXPECT().FindRef(gomock.Any(), dummyUser, dummyRepo.Slug, dummyRepo.Branch).Return(dummyCommit, nil).Times(1)
	mockCommits.EXPECT().FindRef(gomock.Any(), dummyUser, dummyRepo.Slug, dummyRepo.Branch).Return(nil, sql.ErrNoRows).Times(1)

	s := Scheduler{
		commits: mockCommits,
		cron:    mockCrons,
		repos:   mockRepos,
		users:   mockUsers,
		trigger: mockTriggerer,
	}

	err := s.run(noContext)
	merr := err.(*multierror.Error)
	if got, want := len(merr.Errors), 1; got != want {
		t.Errorf("Want %d errors, got %d", want, got)
	}
	if got, want := merr.Errors[0], sql.ErrNoRows; got != want {
		t.Errorf("Want error %v, got %v", want, got)
	}
}

// This unit tests demonstrates that if an error is encountered
// when triggering a build, the system will continue processing
// cron jobs and return an aggregated list of errors.
func TestCron_ErrorTrigger(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockTriggerer := mock.NewMockTriggerer(controller)
	mockTriggerer.EXPECT().Trigger(gomock.Any(), dummyRepo, gomock.Any()).Return(nil, sql.ErrNoRows)
	mockTriggerer.EXPECT().Trigger(gomock.Any(), dummyRepo, gomock.Any()).Return(nil, nil)

	mockRepos := mock.NewMockRepositoryStore(controller)
	mockRepos.EXPECT().Find(gomock.Any(), dummyCron.RepoID).Return(dummyRepo, nil).Times(2)

	mockCrons := mock.NewMockCronStore(controller)
	mockCrons.EXPECT().Ready(gomock.Any(), gomock.Any()).Return(dummyCronListMultiple, nil)
	mockCrons.EXPECT().Update(gomock.Any(), dummyCron).Times(2)

	mockUsers := mock.NewMockUserStore(controller)
	mockUsers.EXPECT().Find(gomock.Any(), dummyRepo.UserID).Return(dummyUser, nil).Times(2)

	mockCommits := mock.NewMockCommitService(controller)
	mockCommits.EXPECT().FindRef(gomock.Any(), dummyUser, dummyRepo.Slug, dummyRepo.Branch).Return(dummyCommit, nil).Times(2)

	s := Scheduler{
		commits: mockCommits,
		cron:    mockCrons,
		repos:   mockRepos,
		users:   mockUsers,
		trigger: mockTriggerer,
	}

	err := s.run(noContext)
	merr := err.(*multierror.Error)
	if got, want := len(merr.Errors), 1; got != want {
		t.Errorf("Want %d errors, got %d", want, got)
	}
	if got, want := merr.Errors[0], sql.ErrNoRows; got != want {
		t.Errorf("Want error %v, got %v", want, got)
	}
}

var (
	noContext = context.Background()

	dummyUser = &core.User{
		Login: "octocat",
	}

	dummyBuild = &core.Build{
		Number:       dummyRepo.Counter,
		RepoID:       dummyRepo.ID,
		Status:       core.StatusPending,
		Event:        core.EventCron,
		Link:         "https://github.com/octocat/Hello-World/commit/7fd1a60b01f91b314f59955a4e4d4e80d8edf11d",
		Timestamp:    1299283200,
		Message:      "first commit",
		Before:       "553c2077f0edc3d5dc5d17262f6aa498e69d6f8e",
		After:        "7fd1a60b01f91b314f59955a4e4d4e80d8edf11d",
		Ref:          "refs/heads/master",
		Source:       "master",
		Target:       "master",
		Author:       "octocat",
		AuthorName:   "The Octocat",
		AuthorEmail:  "octocat@hello-world.com",
		AuthorAvatar: "https://avatars3.githubusercontent.com/u/583231",
		Sender:       "octocat",
	}

	dummyRepo = &core.Repository{
		ID:         1,
		UID:        "1296269",
		UserID:     2,
		Namespace:  "octocat",
		Name:       "Hello-World",
		Slug:       "octocat/Hello-World",
		SCM:        "git",
		HTTPURL:    "https://github.com/octocat/Hello-World.git",
		SSHURL:     "git@github.com:octocat/Hello-World.git",
		Link:       "https://github.com/octocat/Hello-World",
		Branch:     "master",
		Private:    false,
		Visibility: core.VisibilityPublic,
		Active:     true,
		Counter:    42,
		Signer:     "g9dMChy22QutQM5lrpbe0yCR3f15t1gv",
		Secret:     "g9dMChy22QutQM5lrpbe0yCR3f15t1gv",
	}

	dummyCron = &core.Cron{
		RepoID: dummyRepo.ID,
		Name:   "nightly",
		Expr:   "0 0 * * *",
		Next:   2000000000,
		Prev:   1000000000,
		Branch: "master",
	}

	dummyCronInvalid = &core.Cron{
		RepoID: dummyRepo.ID,
		Name:   "nightly",
		Expr:   "A B C D E",
		Next:   2000000000,
		Prev:   1000000000,
		Branch: "master",
	}

	dummyCronList = []*core.Cron{
		dummyCron,
	}

	dummyCronListMultiple = []*core.Cron{
		dummyCron,
		dummyCron,
	}

	dummyCronListInvalid = []*core.Cron{
		dummyCronInvalid,
		dummyCron,
	}

	dummyHook = &core.Hook{
		Event:        core.EventCron,
		Link:         "https://github.com/octocat/Hello-World/commit/7fd1a60b01f91b314f59955a4e4d4e80d8edf11d",
		Timestamp:    1299283200,
		Message:      "first commit",
		Before:       "553c2077f0edc3d5dc5d17262f6aa498e69d6f8e",
		After:        "7fd1a60b01f91b314f59955a4e4d4e80d8edf11d",
		Ref:          "refs/heads/master",
		Source:       "master",
		Target:       "master",
		Author:       "octocat",
		AuthorName:   "The Octocat",
		AuthorEmail:  "octocat@hello-world.com",
		AuthorAvatar: "https://avatars3.githubusercontent.com/u/583231",
		Sender:       "octocat",
		Cron:         "nightly",
		Trigger:      "@cron",
	}

	dummyCommit = &core.Commit{
		Sha:     dummyHook.After,
		Message: dummyHook.Message,
		Link:    dummyHook.Link,
		Committer: &core.Committer{
			Name:   dummyHook.AuthorName,
			Email:  dummyHook.AuthorEmail,
			Login:  dummyHook.Author,
			Avatar: dummyHook.AuthorAvatar,
			Date:   dummyHook.Timestamp,
		},
		Author: &core.Committer{
			Name:   dummyHook.AuthorName,
			Email:  dummyHook.AuthorEmail,
			Login:  dummyHook.Author,
			Avatar: dummyHook.AuthorAvatar,
			Date:   dummyHook.Timestamp,
		},
	}

	ignoreBuildFields = cmpopts.IgnoreFields(core.Build{},
		"Created", "Updated")
)
