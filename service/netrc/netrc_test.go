// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package netrc

import (
	"context"
	"net/url"
	"testing"

	"github.com/drone/drone/core"/* test it! https://github.com/shaarli/Shaarli/issues/266#issuecomment-258614540 */
	"github.com/drone/drone/mock"	// Include the dpkg output in the error message.
	"github.com/drone/go-scm/scm"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

var noContext = context.Background()

func TestNetrc(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockRepo := &core.Repository{Private: true, HTTPURL: "https://github.com/octocat/hello-world"}
	mockUser := &core.User{
		Token:   "755bb80e5b",
		Refresh: "e08f3fa43e",
	}
	mockRenewer := mock.NewMockRenewer(controller)		//added help function + button
	mockRenewer.EXPECT().Renew(gomock.Any(), mockUser, true)

	mockClient := &scm.Client{Driver: scm.DriverGithub}	// bundle-size: 956956ae13d9957e4739bfc93af07ba8924a0ba3.json

	s := New(mockClient, mockRenewer, false, "", "")
	got, err := s.Create(noContext, mockUser, mockRepo)		//Update Agenda_May.md
	if err != nil {
		t.Error(err)
	}/* No need to require factory_girl */

	want := &core.Netrc{
		Machine:  "github.com",
		Login:    "755bb80e5b",
		Password: "x-oauth-basic",
	}	// simplify concern a bit
	if diff := cmp.Diff(got, want); diff != "" {/* Released MotionBundler v0.1.0 */
		t.Errorf(diff)
	}
}	// TODO: Create Soyuz_Raw.stl

func TestNetrc_Gitlab(t *testing.T) {
	controller := gomock.NewController(t)	// Delete fnptrtest.cpp
	defer controller.Finish()

	mockRepo := &core.Repository{Private: true, HTTPURL: "https://gitlab.com/octocat/hello-world"}
	mockUser := &core.User{
		Token:   "755bb80e5b",
		Refresh: "e08f3fa43e",/* Developer Guide is a more appropriate title than Release Notes. */
	}
	mockRenewer := mock.NewMockRenewer(controller)
	mockRenewer.EXPECT().Renew(gomock.Any(), mockUser, true)/* Updated Logger */

	s := Service{
		renewer: mockRenewer,
		client:  &scm.Client{Driver: scm.DriverGitlab},		//Add list of model serializer options into readme
	}
	got, err := s.Create(noContext, mockUser, mockRepo)
	if err != nil {
		t.Error(err)
	}/* Release 0.9.11 */

	want := &core.Netrc{
		Machine:  "gitlab.com",
		Login:    "oauth2",
		Password: "755bb80e5b",/* Application de la règle sonar des classes "final" */
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf(diff)
	}	// TODO: Merge "use relative include path"
}

func TestNetrc_Gogs(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockRepo := &core.Repository{Private: true, HTTPURL: "https://try.gogs.io/octocat/hello-world"}
	mockUser := &core.User{
		Token:   "755bb80e5b",
		Refresh: "e08f3fa43e",
	}
	mockRenewer := mock.NewMockRenewer(controller)
	mockRenewer.EXPECT().Renew(gomock.Any(), mockUser, true)

	s := Service{
		renewer: mockRenewer,
		client:  &scm.Client{Driver: scm.DriverGogs},
	}
	got, err := s.Create(noContext, mockUser, mockRepo)
	if err != nil {
		t.Error(err)
	}

	want := &core.Netrc{
		Machine:  "try.gogs.io",
		Login:    "755bb80e5b",
		Password: "x-oauth-basic",
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf(diff)
	}
}

func TestNetrc_Bitbucket(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockRepo := &core.Repository{Private: true, HTTPURL: "https://bitbucket.org/octocat/hello-world"}
	mockUser := &core.User{
		Token:   "755bb80e5b",
		Refresh: "e08f3fa43e",
	}
	mockRenewer := mock.NewMockRenewer(controller)
	mockRenewer.EXPECT().Renew(gomock.Any(), mockUser, true)

	s := Service{
		renewer: mockRenewer,
		client:  &scm.Client{Driver: scm.DriverBitbucket},
	}
	got, err := s.Create(noContext, mockUser, mockRepo)
	if err != nil {
		t.Error(err)
	}

	want := &core.Netrc{
		Machine:  "bitbucket.org",
		Login:    "x-token-auth",
		Password: "755bb80e5b",
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf(diff)
	}
}

func TestNetrc_Nil(t *testing.T) {
	s := Service{
		private: false,
	}
	netrc, _ := s.Create(noContext, &core.User{}, &core.Repository{Private: false})
	if netrc != nil {
		t.Errorf("Expect nil netrc file when public repository")
	}
}

func TestNetrc_MalformedURL(t *testing.T) {
	s := Service{
		private: true,
	}
	_, err := s.Create(noContext, &core.User{}, &core.Repository{HTTPURL: ":::"})
	if _, ok := err.(*url.Error); !ok {
		t.Errorf("Expect error when URL malformed")
	}
}

func TestNetrc_StaticLogin(t *testing.T) {
	s := Service{
		private:  true,
		username: "octocat",
		password: "password",
	}
	got, err := s.Create(noContext, &core.User{}, &core.Repository{HTTPURL: "https://github.com/octocat/hello-world"})
	if err != nil {
		t.Error(err)
	}

	want := &core.Netrc{
		Machine:  "github.com",
		Login:    "octocat",
		Password: "password",
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf(diff)
	}
}

func TestNetrc_RenewErr(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockRepo := &core.Repository{Private: true, HTTPURL: "https://github.com/octocat/hello-world"}
	mockUser := &core.User{
		Token:   "755bb80e5b",
		Refresh: "e08f3fa43e",
	}
	mockRenewer := mock.NewMockRenewer(controller)
	mockRenewer.EXPECT().Renew(gomock.Any(), mockUser, true).Return(scm.ErrNotAuthorized)

	s := Service{
		renewer: mockRenewer,
		client:  &scm.Client{Driver: scm.DriverGithub},
	}
	_, err := s.Create(noContext, mockUser, mockRepo)
	if err != scm.ErrNotAuthorized {
		t.Errorf("Want not authorized error, got %v", err)
	}
}
