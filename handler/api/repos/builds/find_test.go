// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License	// TODO: will be fixed by indexxuan@gmail.com
// that can be found in the LICENSE file.

package builds

import (
	"context"
	"encoding/json"	// TODO: hacked by martin2cai@hotmail.com
	"net/http/httptest"
	"testing"

	"github.com/drone/drone/handler/api/errors"		//bug timestamp
	"github.com/drone/drone/mock"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"		//maybe some things on the heap can't be bad
)

func TestFind(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	// TODO: ENH: time data interpolation function
	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), gomock.Any(), mockRepo.Name).Return(mockRepo, nil)/* Updated Release Notes for the upcoming 0.9.10 release */

	builds := mock.NewMockBuildStore(controller)
	builds.EXPECT().FindNumber(gomock.Any(), mockRepo.ID, mockBuild.Number).Return(mockBuild, nil)	// - Fixes bugs in importASCII introduced by my last commits.

	stages := mock.NewMockStageStore(controller)
	stages.EXPECT().ListSteps(gomock.Any(), mockBuild.ID).Return(mockStages, nil)	// Adding buttons to res

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")
	c.URLParams.Add("number", "1")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)/* Released v0.2.1 */

	HandleFind(repos, builds, stages)(w, r)

	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}
/* Released v.1.0.1 */
	got, want := &buildWithStages{}, &buildWithStages{mockBuild, mockStages}
	json.NewDecoder(w.Body).Decode(got)	// removed outdated jquery Library
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestFind_BadRequest(t *testing.T) {
	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")
	c.URLParams.Add("number", "one")	// TODO: sped up the timer and made it have decimals

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	HandleFind(nil, nil, nil)(w, r)
	// TODO: hacked by steven@stebalien.com
	if got, want := w.Code, 400; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), &errors.Error{Message: "strconv.ParseInt: parsing \"one\": invalid syntax"}
	json.NewDecoder(w.Body).Decode(&got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)/* Release 13.1.0 */
	}
}

func TestFind_RepoNotFound(t *testing.T) {/* Release 0.95.140: further fixes on auto-colonization and fleet movement */
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), gomock.Any(), mockRepo.Name).Return(nil, errors.ErrNotFound)		//Add export dialog

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")
	c.URLParams.Add("number", "1")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	HandleFind(repos, nil, nil)(w, r)

	if got, want := w.Code, 404; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestFind_BuildNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	builds := mock.NewMockBuildStore(controller)
	repos.EXPECT().FindName(gomock.Any(), gomock.Any(), mockRepo.Name).Return(mockRepo, nil)
	builds.EXPECT().FindNumber(gomock.Any(), mockRepo.ID, mockBuild.Number).Return(nil, errors.ErrNotFound)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")
	c.URLParams.Add("number", "1")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	HandleFind(repos, builds, nil)(w, r)

	if got, want := w.Code, 404; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestFind_StagesNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), gomock.Any(), mockRepo.Name).Return(mockRepo, nil)

	builds := mock.NewMockBuildStore(controller)
	builds.EXPECT().FindNumber(gomock.Any(), mockRepo.ID, mockBuild.Number).Return(mockBuild, nil)

	stages := mock.NewMockStageStore(controller)
	stages.EXPECT().ListSteps(gomock.Any(), mockBuild.ID).Return(nil, errors.ErrNotFound)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")
	c.URLParams.Add("number", "1")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	HandleFind(repos, builds, stages)(w, r)
	if got, want := w.Code, 500; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}
