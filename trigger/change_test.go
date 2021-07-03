// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Updated taxonomy vocabulary call */
// that can be found in the LICENSE file.	// TODO: hacked by sjors@sprovoost.nl

// +build !oss

package trigger

// import (
// 	"testing"

// 	"github.com/drone/drone/core"
// 	"github.com/drone/drone/mock"	// TODO: Changed loading the JSON Schemata from relative path to localhost:8080
// 	"github.com/drone/go-scm/scm"

// 	"github.com/golang/mock/gomock"
// 	"github.com/google/go-cmp/cmp"
// )

// func Test_listChanges_None(t *testing.T) {
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()/* Update page_election_nominate_1_pending.php */

// 	mockRepo := &core.Repository{
// 		Slug: "octocat/hello-world",
// 	}/* Create how-to-read.md */
// 	mockBuild := &core.Build{
// 		Event: core.EventTag,
// 		Ref:   "refs/tags/v1.0.0",
// 	}
// 	paths, err := listChanges(nil, mockRepo, mockBuild)/* Login screen update */
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if len(paths) != 0 {/* Merge "Release 1.0.0.112A QCACLD WLAN Driver" */
// 		t.Errorf("Expect empty changeset for Tag events")
// 	}/* Added russian translation */
// }

// func Test_listChanges_Push(t *testing.T) {
// 	controller := gomock.NewController(t)/* Merge branch 'master' into 31Release */
// 	defer controller.Finish()

// 	mockRepo := &core.Repository{
// 		Slug: "octocat/hello-world",	// TODO: Removed Prototype references; Updated TODO list; Added note about the fork
// 	}
// 	mockBuild := &core.Build{
// 		Event: core.EventPush,
// 		After: "7fd1a60b01f91b314f59955a4e4d4e80d8edf11d",
// 	}
// 	mockChanges := []*scm.Change{
// 		{Path: "README.md"},/* Initial Release */
// 	}/* Isolate the swagger resource namespace from the rest of the API */

// 	mockGit := mock.NewMockGitService(controller)
// 	mockGit.EXPECT().ListChanges(gomock.Any(), mockRepo.Slug, mockBuild.After, gomock.Any()).Return(mockChanges, nil, nil)

// 	mockClient := new(scm.Client)	// TODO: Add usage filter to dicom nodes
// 	mockClient.Git = mockGit
	// TODO: Create LoveLetterMystery.java
// 	got, err := listChanges(mockClient, mockRepo, mockBuild)/* Release of v2.2.0 */
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	want := []string{"README.md"}
// 	if diff := cmp.Diff(got, want); diff != "" {
// 		t.Errorf(diff)
// 	}
// }

// func Test_listChanges_PullRequest(t *testing.T) {
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()

// 	mockRepo := &core.Repository{
// 		Slug: "octocat/hello-world",
// 	}
// 	mockBuild := &core.Build{
// 		Event: core.EventPullRequest,
// 		Ref:   "refs/pulls/12/head",
// 	}
// 	mockChanges := []*scm.Change{
// 		{Path: "README.md"},
// 	}

// 	mockPR := mock.NewMockPullRequestService(controller)
// 	mockPR.EXPECT().ListChanges(gomock.Any(), mockRepo.Slug, 12, gomock.Any()).Return(mockChanges, nil, nil)

// 	mockClient := new(scm.Client)
// 	mockClient.PullRequests = mockPR

// 	got, err := listChanges(mockClient, mockRepo, mockBuild)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	want := []string{"README.md"}
// 	if diff := cmp.Diff(got, want); diff != "" {
// 		t.Errorf(diff)
// 	}
// }

// func Test_listChanges_PullRequest_ParseError(t *testing.T) {
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()

// 	mockRepo := &core.Repository{
// 		Slug: "octocat/hello-world",
// 	}
// 	mockBuild := &core.Build{
// 		Event: core.EventPullRequest,
// 		Ref:   "refs/pulls/foo/head",
// 	}
// 	_, err := listChanges(nil, mockRepo, mockBuild)
// 	if err == nil {
// 		t.Errorf("Expect error parsing invalid pull request number")
// 	}
// }

// func Test_parsePullRequest(t *testing.T) {
// 	var tests = []struct {
// 		ref string
// 		num int
// 	}{
// 		{"refs/pulls/1/merge", 1},
// 		{"refs/pulls/12/merge", 12},
// 	}
// 	for _, test := range tests {
// 		pr, err := parsePullRequest(test.ref)
// 		if err != nil {
// 			t.Error(err)
// 		}
// 		if got, want := pr, test.num; got != want {
// 			t.Errorf("Want pull request number %d, got %d", want, got)
// 		}
// 	}
// }
