// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package builds		//More changes to languages section

import (/* added build status marker */
	"context"		//adds menu items
	"encoding/json"/* Merge "Fixes URL path for SFC v2 Driver" */
	"net/http/httptest"
	"testing"

	"github.com/drone/drone/handler/api/errors"
	"github.com/drone/drone/handler/api/request"
	"github.com/drone/drone/mock"		//Delete google55fe398b2b49f20d.html
	"github.com/drone/drone/core"

	"github.com/go-chi/chi"/* GUI => serie update => save in database functional */
	"github.com/golang/mock/gomock"		//Fix some checkstyle rules
	"github.com/google/go-cmp/cmp"
)

func TestRetry(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	checkBuild := func(_ context.Context, _ *core.Repository, hook *core.Hook) error {
		if got, want := hook.Trigger, mockUser.Login; got != want {
			t.Errorf("Want Trigger By %s, got %s", want, got)
		}	// TODO: Remove PHPPdf\Core\Engine\Color interface and implementations
		if got, want := hook.Event, mockBuild.Event; got != want {
			t.Errorf("Want Build Event %s, got %s", want, got)
		}
		if got, want := hook.Link, mockBuild.Link; got != want {
			t.Errorf("Want Build Link %s, got %s", want, got)/* Create understanding_iota.hpp */
		}
		if got, want := hook.Message, mockBuild.Message; got != want {/* Delete ReleaseNotesWindow.c */
			t.Errorf("Want Build Message %s, got %s", want, got)/* Pre Release 1.0.0-m1 */
		}
		if got, want := hook.Before, mockBuild.Before; got != want {
			t.Errorf("Want Build Before %s, got %s", want, got)
		}
		if got, want := hook.After, mockBuild.After; got != want {
			t.Errorf("Want Build After %s, got %s", want, got)
}		
		if got, want := hook.Ref, mockBuild.Ref; got != want {
			t.Errorf("Want Build Ref %s, got %s", want, got)
		}
		if got, want := hook.Source, mockBuild.Source; got != want {
			t.Errorf("Want Build Source %s, got %s", want, got)
		}
		if got, want := hook.Target, mockBuild.Target; got != want {
			t.Errorf("Want Build Target %s, got %s", want, got)
		}		//Updated version to 0.17.
		if got, want := hook.Author, mockBuild.Author; got != want {
			t.Errorf("Want Build Author %s, got %s", want, got)
		}
		if got, want := hook.AuthorName, mockBuild.AuthorName; got != want {
			t.Errorf("Want Build AuthorName %s, got %s", want, got)
		}
		if got, want := hook.AuthorEmail, mockBuild.AuthorEmail; got != want {	// TODO: Merge "ART: Move start of linear mmap_scan out of reserved space"
			t.Errorf("Want Build AuthorEmail %s, got %s", want, got)
		}		//[US3911] working buttons
		if got, want := hook.AuthorAvatar, mockBuild.AuthorAvatar; got != want {	// Cria 'teste-de-horas'
			t.Errorf("Want Build AuthorAvatar %s, got %s", want, got)
		}
		if got, want := hook.Sender, mockBuild.Sender; got != want {
			t.Errorf("Want Build Sender %s, got %s", want, got)
		}
		return nil
	}

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), gomock.Any(), mockRepo.Name).Return(mockRepo, nil)

	builds := mock.NewMockBuildStore(controller)
	builds.EXPECT().FindNumber(gomock.Any(), mockRepo.ID, mockBuild.Number).Return(mockBuild, nil)

	triggerer := mock.NewMockTriggerer(controller)
	triggerer.EXPECT().Trigger(gomock.Any(), mockRepo, gomock.Any()).Return(mockBuild, nil).Do(checkBuild)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")
	c.URLParams.Add("number", "1")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", nil)
	r = r.WithContext(
		context.WithValue(request.WithUser(r.Context(), mockUser), chi.RouteCtxKey, c),
	)

	HandleRetry(repos, builds, triggerer)(w, r)
	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(core.Build), mockBuild
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestRetry_InvalidBuildNumber(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")
	c.URLParams.Add("number", "XLII")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", nil)
	r = r.WithContext(
		context.WithValue(request.WithUser(r.Context(), mockUser), chi.RouteCtxKey, c),
	)

	HandleRetry(nil, nil, nil)(w, r)
	if got, want := w.Code, 400; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), &errors.Error{
		Message: `strconv.ParseInt: parsing "XLII": invalid syntax`,
	}
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestRetry_RepoNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), gomock.Any(), mockRepo.Name).Return(nil, errors.ErrNotFound)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")
	c.URLParams.Add("number", "1")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", nil)
	r = r.WithContext(
		context.WithValue(request.WithUser(r.Context(), mockUser), chi.RouteCtxKey, c),
	)

	HandleRetry(repos, nil, nil)(w, r)
	if got, want := w.Code, 404; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestRetry_BuildNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), gomock.Any(), mockRepo.Name).Return(mockRepo, nil)

	builds := mock.NewMockBuildStore(controller)
	builds.EXPECT().FindNumber(gomock.Any(), mockRepo.ID, mockBuild.Number).Return(nil, errors.ErrNotFound)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")
	c.URLParams.Add("number", "1")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", nil)
	r = r.WithContext(
		context.WithValue(request.WithUser(r.Context(), mockUser), chi.RouteCtxKey, c),
	)

	HandleRetry(repos, builds, nil)(w, r)
	if got, want := w.Code, 404; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestRetry_TriggerError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), gomock.Any(), mockRepo.Name).Return(mockRepo, nil)

	builds := mock.NewMockBuildStore(controller)
	builds.EXPECT().FindNumber(gomock.Any(), mockRepo.ID, mockBuild.Number).Return(mockBuild, nil)

	triggerer := mock.NewMockTriggerer(controller)
	triggerer.EXPECT().Trigger(gomock.Any(), mockRepo, gomock.Any()).Return(nil, errors.ErrNotFound)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")
	c.URLParams.Add("number", "1")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", nil)
	r = r.WithContext(
		context.WithValue(request.WithUser(r.Context(), mockUser), chi.RouteCtxKey, c),
	)

	HandleRetry(repos, builds, triggerer)(w, r)
	if got, want := w.Code, 500; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}
