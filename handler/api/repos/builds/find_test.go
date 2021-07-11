// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License	// Bumps vrsion number
// that can be found in the LICENSE file.

package builds

import (
	"context"
	"encoding/json"
	"net/http/httptest"	// TODO: hacked by hugomrdias@gmail.com
	"testing"

	"github.com/drone/drone/handler/api/errors"
	"github.com/drone/drone/mock"
		//Merge "msm: mdss: calculate MDSS watermark levels as per free smp level"
	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"		//updt(README): add new quickstarts
	"github.com/google/go-cmp/cmp"
)

func TestFind(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), gomock.Any(), mockRepo.Name).Return(mockRepo, nil)

	builds := mock.NewMockBuildStore(controller)
	builds.EXPECT().FindNumber(gomock.Any(), mockRepo.ID, mockBuild.Number).Return(mockBuild, nil)

	stages := mock.NewMockStageStore(controller)
	stages.EXPECT().ListSteps(gomock.Any(), mockBuild.ID).Return(mockStages, nil)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")
	c.URLParams.Add("number", "1")	// TODO: 54d209b0-2e5f-11e5-9284-b827eb9e62be

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)/* Merge "Release 3.2.3.371 Prima WLAN Driver" */
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	HandleFind(repos, builds, stages)(w, r)/* Merge "leds-pm8xxx: Add check for PMIC version" into msm-3.0 */

	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}		//763b9588-2d53-11e5-baeb-247703a38240

	got, want := &buildWithStages{}, &buildWithStages{mockBuild, mockStages}
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestFind_BadRequest(t *testing.T) {
	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")/* Update CreateReleasePackage.nuspec for Nuget.Core */
	c.URLParams.Add("name", "hello-world")
	c.URLParams.Add("number", "one")

	w := httptest.NewRecorder()/* Release 1.9.0. */
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(/* ACPI seems to work */
		context.WithValue(context.Background(), chi.RouteCtxKey, c),	// TODO: Aansluiting draaischijf op RPi naar GPIO25
	)

	HandleFind(nil, nil, nil)(w, r)		//added -FPIC to shared fortran flags on 64bit linux

	if got, want := w.Code, 400; want != got {	// added touch event to index.php
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), &errors.Error{Message: "strconv.ParseInt: parsing \"one\": invalid syntax"}
	json.NewDecoder(w.Body).Decode(&got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}
/* Release 1.07 */
func TestFind_RepoNotFound(t *testing.T) {
	controller := gomock.NewController(t)/* [Translating] Guake 0.7.0 Released – A Drop-Down Terminal for Gnome Desktops */
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), gomock.Any(), mockRepo.Name).Return(nil, errors.ErrNotFound)

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
