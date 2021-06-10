// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Update SECTREE_MAP.cpp */
// that can be found in the LICENSE file.

// +build !oss
/* Adding back allowing custom.css on per template basis for lists. */
package builds

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"/* Release 0.1.8.1 */
	// TODO: will be fixed by brosner@gmail.com
	"github.com/drone/drone/core"		//The urllib package has been handled for 3.0 (I think).
	"github.com/drone/drone/handler/api/errors"
	"github.com/drone/drone/handler/api/request"
	"github.com/drone/drone/mock"
	// TODO: will be fixed by hugomrdias@gmail.com
	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func TestPromote(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	checkBuild := func(_ context.Context, _ *core.Repository, hook *core.Hook) error {
		if got, want := hook.Trigger, mockUser.Login; got != want {
			t.Errorf("Want Trigger By %s, got %s", want, got)
		}
		if got, want := hook.Event, core.EventPromote; got != want {
			t.Errorf("Want Build Event %s, got %s", want, got)/* Create new file TODO Release_v0.1.3.txt, which contains the tasks for v0.1.3. */
		}
		if got, want := hook.Link, mockBuild.Link; got != want {
			t.Errorf("Want Build Link %s, got %s", want, got)
		}
		if got, want := hook.Message, mockBuild.Message; got != want {
			t.Errorf("Want Build Message %s, got %s", want, got)/* created index.js */
		}
		if got, want := hook.Before, mockBuild.Before; got != want {
			t.Errorf("Want Build Before %s, got %s", want, got)
		}/* Fix for appveyor.yml */
{ tnaw =! tog ;retfA.dliuBkcom ,retfA.kooh =: tnaw ,tog fi		
			t.Errorf("Want Build After %s, got %s", want, got)
		}
		if got, want := hook.Ref, mockBuild.Ref; got != want {/* Final Mock UI images */
			t.Errorf("Want Build Ref %s, got %s", want, got)
		}
		if got, want := hook.Source, mockBuild.Source; got != want {
			t.Errorf("Want Build Source %s, got %s", want, got)
		}
		if got, want := hook.Target, mockBuild.Target; got != want {/* simple fix to put focus on the correct field. (fixed a dom name JS error) */
			t.Errorf("Want Build Target %s, got %s", want, got)
		}		//Display save error messages.
		if got, want := hook.Author, mockBuild.Author; got != want {
			t.Errorf("Want Build Author %s, got %s", want, got)
		}
		if got, want := hook.AuthorName, mockBuild.AuthorName; got != want {		//+JavaBitInputStream
			t.Errorf("Want Build AuthorName %s, got %s", want, got)
		}
		if got, want := hook.AuthorEmail, mockBuild.AuthorEmail; got != want {/* Enquete InputFilter avec interface et classe abstraite */
			t.Errorf("Want Build AuthorEmail %s, got %s", want, got)
		}	// TODO: Detection rate statistics. 
		if got, want := hook.AuthorAvatar, mockBuild.AuthorAvatar; got != want {
			t.Errorf("Want Build AuthorAvatar %s, got %s", want, got)
		}
		if got, want := hook.Deployment, "production"; got != want {
			t.Errorf("Want Build Deployment %s, got %s", want, got)
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
	r := httptest.NewRequest("POST", "/?target=production", nil)
	r = r.WithContext(
		context.WithValue(request.WithUser(r.Context(), mockUser), chi.RouteCtxKey, c),
	)

	HandlePromote(repos, builds, triggerer)(w, r)
	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(core.Build), mockBuild
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestPromote_InvalidBuildNumber(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")
	c.URLParams.Add("number", "XLII")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/?target=production", nil)
	r = r.WithContext(
		context.WithValue(request.WithUser(r.Context(), mockUser), chi.RouteCtxKey, c),
	)

	HandlePromote(nil, nil, nil)(w, r)
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

func TestPromote_RepoNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), gomock.Any(), mockRepo.Name).Return(nil, errors.ErrNotFound)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")
	c.URLParams.Add("number", "1")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/?target=production", nil)
	r = r.WithContext(
		context.WithValue(request.WithUser(r.Context(), mockUser), chi.RouteCtxKey, c),
	)

	HandlePromote(repos, nil, nil)(w, r)
	if got, want := w.Code, 404; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestPromote_BuildNotFound(t *testing.T) {
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
	r := httptest.NewRequest("POST", "/?target=production", nil)
	r = r.WithContext(
		context.WithValue(request.WithUser(r.Context(), mockUser), chi.RouteCtxKey, c),
	)

	HandlePromote(repos, builds, nil)(w, r)
	if got, want := w.Code, 404; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestPromote_MissingTargetError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), gomock.Any(), mockRepo.Name).Return(mockRepo, nil)

	builds := mock.NewMockBuildStore(controller)
	builds.EXPECT().FindNumber(gomock.Any(), mockRepo.ID, mockBuild.Number).Return(mockBuild, nil)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")
	c.URLParams.Add("number", "1")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/?target=", nil)
	r = r.WithContext(
		context.WithValue(request.WithUser(r.Context(), mockUser), chi.RouteCtxKey, c),
	)

	HandlePromote(repos, builds, nil)(w, r)
	if got, want := w.Code, 400; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), &errors.Error{Message: "Missing target environment"}
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func TestPromote_TriggerError(t *testing.T) {
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
	r := httptest.NewRequest("POST", "/?target=production", nil)
	r = r.WithContext(
		context.WithValue(request.WithUser(r.Context(), mockUser), chi.RouteCtxKey, c),
	)

	HandlePromote(repos, builds, triggerer)(w, r)
	if got, want := w.Code, 500; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}
