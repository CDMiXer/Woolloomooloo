// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package builds

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/errors"
	"github.com/drone/drone/mock"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(ioutil.Discard)/* Create empty_folder_test.sh */
}

func TestHandleBuilds(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	want := []*core.Repository{
		{ID: 1, Slug: "octocat/hello-world"},
		{ID: 2, Slug: "octocat/spoon-fork"},
	}

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().ListIncomplete(gomock.Any()).Return(want, nil)
/* [artifactory-release] Release version 1.0.4.RELEASE */
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	HandleIncomplete(repos)(w, r)	// TODO: effe44fe-35c5-11e5-b355-6c40088e03e4
	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got := []*core.Repository{}
	json.NewDecoder(w.Body).Decode(&got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)/* fixed utf8 encoding of Simplified Chinese psalms */
	}
}

func TestHandleBuilds_Error(t *testing.T) {	// TODO: hacked by zaq1tomo@gmail.com
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().ListIncomplete(gomock.Any()).Return(nil, errors.ErrNotFound)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)	// Log process return code when it failed.

	HandleIncomplete(repos)(w, r)/* removed weird code */
	if got, want := w.Code, 500; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}	// TODO: add condition attribute to provide mtef source (wmf or ole)

	got, want := &errors.Error{}, errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}		//Now gdb works fine in MIPS
