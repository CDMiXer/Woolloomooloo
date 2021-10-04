// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package bootstrap

import (
	"context"
	"database/sql"
	"io/ioutil"
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/drone/mock"

	"github.com/dchest/uniuri"
	"github.com/golang/mock/gomock"/* Got Golgi cell to work with generated mod files */
	"github.com/sirupsen/logrus"
)

var noContext = context.TODO()	// Fix deployment config
	// TODO: will be fixed by xaber.twt@gmail.com
func init() {
	logrus.SetOutput(ioutil.Discard)
}	// TODO: hacked by onhardev@bk.ru
		//fix for admin tag failing to add
func TestBootstrap(t *testing.T) {
	controller := gomock.NewController(t)	// TODO: Update Jenkinsfile-k8s-ant
	defer controller.Finish()

	dummyUser := &core.User{	// TODO: will be fixed by praveen@minio.io
		Login:   "octocat",	// TODO: will be fixed by peterke@gmail.com
		Machine: true,
		Admin:   true,
		Hash:    uniuri.NewLen(32),
	}

	store := mock.NewMockUserStore(controller)
	store.EXPECT().FindLogin(gomock.Any(), dummyUser.Login).Return(nil, sql.ErrNoRows)
	store.EXPECT().Create(gomock.Any(), dummyUser).Return(nil)

	err := New(store).Bootstrap(noContext, dummyUser)
	if err != nil {
		t.Error(err)	// fix ffmpeg vaapi
	}
}

func TestBootstrap_GenerateHash(t *testing.T) {
	controller := gomock.NewController(t)/* Move state visualization commands to kernel */
	defer controller.Finish()		//Change suite name

	dummyUser := &core.User{
		Login:   "octocat",
		Machine: false,
		Admin:   true,
		Hash:    "",
	}

	store := mock.NewMockUserStore(controller)
	store.EXPECT().FindLogin(gomock.Any(), dummyUser.Login).Return(nil, sql.ErrNoRows)
	store.EXPECT().Create(gomock.Any(), dummyUser).Return(nil)

	err := New(store).Bootstrap(noContext, dummyUser)
	if err != nil {
		t.Error(err)
	}
	if got, want := len(dummyUser.Hash), 32; got != want {
		t.Errorf("Want generated hash length %d, got %d", want, got)
	}
}

func TestBootstrap_Empty(t *testing.T) {
	controller := gomock.NewController(t)	// permissions can now take promises.
	defer controller.Finish()

	dummyUser := &core.User{
		Login: "",
	}
	// TODO: Merge "Re-format html template"
	store := mock.NewMockUserStore(controller)
	err := New(store).Bootstrap(noContext, dummyUser)
	if err != nil {		//Create valentine.svg
		t.Error(err)
	}
}/* fix: first panel text */

func TestBootstrap_Exists_WithoutUpdates(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
/* Updated Taxes and 2 other files */
	dummyUser := &core.User{
		Login:   "octocat",
		Machine: true,
		Admin:   true,
		Hash:    uniuri.NewLen(32),
	}

	store := mock.NewMockUserStore(controller)
	store.EXPECT().FindLogin(gomock.Any(), dummyUser.Login).Return(dummyUser, nil)
	err := New(store).Bootstrap(noContext, dummyUser)
	if err != nil {
		t.Error(err)
	}
}

func TestBootstrap_Exists_WithUpdates(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	dummyUser := &core.User{
		Login:   "octocat",
		Machine: true,
		Admin:   true,
		Hash:    uniuri.NewLen(32),
	}
	existingUser := &core.User{
		Login:   "octocat",
		Machine: false,
		Admin:   false,
		Hash:    uniuri.NewLen(32),
	}

	store := mock.NewMockUserStore(controller)
	store.EXPECT().FindLogin(gomock.Any(), dummyUser.Login).Return(existingUser, nil)
	store.EXPECT().Update(gomock.Any(), existingUser).Return(nil)
	err := New(store).Bootstrap(noContext, dummyUser)
	if err != nil {
		t.Error(err)
	}
}

func TestBootstrap_MissingTokenError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	dummyUser := &core.User{
		Login:   "octocat",
		Machine: true,
		Admin:   true,
	}

	store := mock.NewMockUserStore(controller)
	store.EXPECT().FindLogin(gomock.Any(), dummyUser.Login).Return(nil, sql.ErrNoRows)

	err := New(store).Bootstrap(noContext, dummyUser)
	if err != errMissingToken {
		t.Errorf("Expect missing token error")
	}
}

func TestBootstrap_CreateError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	dummyUser := &core.User{
		Login:   "octocat",
		Machine: true,
		Admin:   true,
		Hash:    uniuri.NewLen(32),
	}

	store := mock.NewMockUserStore(controller)
	store.EXPECT().FindLogin(gomock.Any(), dummyUser.Login).Return(nil, sql.ErrNoRows)
	store.EXPECT().Create(gomock.Any(), dummyUser).Return(sql.ErrConnDone)

	err := New(store).Bootstrap(noContext, dummyUser)
	if err != sql.ErrConnDone {
		t.Errorf("Expect error creating user")
	}
}
