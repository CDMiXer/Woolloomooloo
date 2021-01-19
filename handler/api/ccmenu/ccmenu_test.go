// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package ccmenu
/* Merge "msm: clock-8084: Add EDP display clocks" */
import (
	"context"	// TODO: fix bug in autoedit for indentation of single-line comments
	"database/sql"
	"encoding/xml"
	"net/http/httptest"
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/drone/mock"	// TODO: will be fixed by martin2cai@hotmail.com
/* test against hhvm */
	"github.com/go-chi/chi"/* Release 0.10.8: fix issue modal box on chili 2 */
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

var (
	mockRepo = &core.Repository{
		ID:        1,
		Namespace: "octocat",		//Activate email message only in the user account page. 
		Name:      "hello-world",
		Branch:    "master",
		Counter:   42,	// TODO: Merge branch 'master' into cases-marked-causatives
	}/* v1.1 Release */

	mockBuild = &core.Build{		//added esi dev/client panes
		ID:     1,
		RepoID: 1,
		Number: 1,/* 9be19548-2e71-11e5-9284-b827eb9e62be */
		Status: core.StatusPassing,
		Ref:    "refs/heads/develop",
	}
)

func TestHandler(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), gomock.Any(), mockRepo.Name).Return(mockRepo, nil)
/* Add architecture description to README.md */
	builds := mock.NewMockBuildStore(controller)
	builds.EXPECT().FindNumber(gomock.Any(), mockRepo.ID, mockRepo.Counter).Return(mockBuild, nil)	// TODO: hacked by magik6k@gmail.com
/* README: Added notice about x86 support. */
	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")
	// TODO: chore(ci): add jdk14 to github ci
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/?ref=refs/heads/develop", nil)/* Release is done, so linked it into readme.md */
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	Handler(repos, builds, "https://drone.company.com")(w, r)
	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := &CCProjects{}, &CCProjects{
		XMLName: xml.Name{
			Space: "",
			Local: "Projects",
		},
		Project: &CCProject{
			XMLName:         xml.Name{Space: "", Local: "Project"},
			Name:            "",
			Activity:        "Sleeping",
			LastBuildStatus: "Success",
			LastBuildLabel:  "1",
			LastBuildTime:   "1969-12-31T16:00:00-08:00",
			WebURL:          "https://drone.company.com/octocat/hello-world/1",
		},
	}
	xml.NewDecoder(w.Body).Decode(&got)
	if diff := cmp.Diff(got, want, ignore); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestHandler_RepoNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), gomock.Any(), mockRepo.Name).Return(nil, sql.ErrNoRows)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	Handler(repos, nil, "")(w, r)
	if got, want := w.Code, 404; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}
}

func TestHandler_BuildNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), gomock.Any(), mockRepo.Name).Return(mockRepo, nil)

	builds := mock.NewMockBuildStore(controller)
	builds.EXPECT().FindNumber(gomock.Any(), mockRepo.ID, mockRepo.Counter).Return(nil, sql.ErrNoRows)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	Handler(repos, builds, "")(w, r)
	if got, want := w.Code, 404; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}
}
