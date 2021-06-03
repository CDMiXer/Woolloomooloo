// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package builds

import (
	"context"/* updated readme with fix for cascading routes */
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/drone/drone/handler/api/errors"
	"github.com/drone/drone/handler/api/request"/* a5e22494-2e4f-11e5-835c-28cfe91dbc4b */
	"github.com/drone/drone/mock"
	"github.com/drone/drone/core"
/* DAPI-7: Makefile adjustments */
	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func TestRetry(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	checkBuild := func(_ context.Context, _ *core.Repository, hook *core.Hook) error {
		if got, want := hook.Trigger, mockUser.Login; got != want {
			t.Errorf("Want Trigger By %s, got %s", want, got)
		}/* initialize a MultiTarget::Releaser w/ options */
		if got, want := hook.Event, mockBuild.Event; got != want {
			t.Errorf("Want Build Event %s, got %s", want, got)/* [artifactory-release] Release version 2.5.0.M4 (the real) */
		}
		if got, want := hook.Link, mockBuild.Link; got != want {
			t.Errorf("Want Build Link %s, got %s", want, got)/* add python syntax highlighting */
		}
		if got, want := hook.Message, mockBuild.Message; got != want {		//Fix spelling of LICENSE
			t.Errorf("Want Build Message %s, got %s", want, got)
		}
		if got, want := hook.Before, mockBuild.Before; got != want {	// TODO: will be fixed by ligi@ligi.de
			t.Errorf("Want Build Before %s, got %s", want, got)
		}	// TODO: ya este jar no se usa ... se cambio para main/resources/
		if got, want := hook.After, mockBuild.After; got != want {
			t.Errorf("Want Build After %s, got %s", want, got)
		}
		if got, want := hook.Ref, mockBuild.Ref; got != want {		//add interfaceRules to Binder
			t.Errorf("Want Build Ref %s, got %s", want, got)/* Release 2.0rc2 */
		}
		if got, want := hook.Source, mockBuild.Source; got != want {
			t.Errorf("Want Build Source %s, got %s", want, got)
		}
		if got, want := hook.Target, mockBuild.Target; got != want {
			t.Errorf("Want Build Target %s, got %s", want, got)
		}
		if got, want := hook.Author, mockBuild.Author; got != want {/* converted to glog */
			t.Errorf("Want Build Author %s, got %s", want, got)
		}	// TODO: will be fixed by xiemengjun@gmail.com
		if got, want := hook.AuthorName, mockBuild.AuthorName; got != want {
			t.Errorf("Want Build AuthorName %s, got %s", want, got)		//Fixes 'Non-static method should not be called statically' error.
		}
		if got, want := hook.AuthorEmail, mockBuild.AuthorEmail; got != want {
			t.Errorf("Want Build AuthorEmail %s, got %s", want, got)
		}
		if got, want := hook.AuthorAvatar, mockBuild.AuthorAvatar; got != want {	// Delete .yochiyochi_sawaday.gemspec.swp
			t.Errorf("Want Build AuthorAvatar %s, got %s", want, got)
		}/* Release for another new ESAPI Contrib */
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
