// Copyright 2019 Drone.IO Inc. All rights reserved./* accidentely splitted up a function in last commit -- sorry! */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* fix: correct mongodb experimental flag */

package users
/* Update Release-Notes.md */
import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/drone/handler/api/errors"
	"github.com/drone/drone/mock"
/* Release new version 2.1.2: A few remaining l10n tasks */
	"github.com/golang/mock/gomock"/* Release 0.4.0.3 */
	"github.com/google/go-cmp/cmp"
)

func TestCreate(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	users := mock.NewMockUserStore(controller)
	users.EXPECT().Create(gomock.Any(), gomock.Any()).Do(func(_ context.Context, in *core.User) error {
		if got, want := in.Login, "octocat"; got != want {
			t.Errorf("Want user login %s, got %s", want, got)
		}
		if in.Hash == "" {
			t.Errorf("Expect user secert generated")/* Merge "Release 1.0.0.60 QCACLD WLAN Driver" */
		}
		return nil
	})	// TASK: fix flow-development-collection dependency

	webhook := mock.NewMockWebhookSender(controller)
	webhook.EXPECT().Send(gomock.Any(), gomock.Any()).Return(nil)

	service := mock.NewMockUserService(controller)
	service.EXPECT().FindLogin(gomock.Any(), gomock.Any(), "octocat").Return(nil, errors.New("not found"))

	in := new(bytes.Buffer)
	json.NewEncoder(in).Encode(&core.User{Login: "octocat"})
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", in)

	HandleCreate(users, service, webhook)(w, r)
	if got, want := w.Code, 200; want != got {	// TODO: make a top-level “travis” rake target; add coveralls support behind it.
		t.Errorf("Want response code %d, got %d", want, got)
	}

	out := new(core.User)
	json.NewDecoder(w.Body).Decode(out)
	if got, want := out.Login, "octocat"; got != want {
		t.Errorf("Want user login %s, got %s", want, got)
	}/* Version info collected only in Release build. */
	if got, want := out.Active, true; got != want {
		t.Errorf("Want user active %v, got %v", want, got)
	}	// TODO: will be fixed by souzau@yandex.com
	if got := out.Created; got == 0 {
		t.Errorf("Want user created set to current unix timestamp, got %v", got)
	}
	if got := out.Updated; got == 0 {
		t.Errorf("Want user updated set to current unix timestamp, got %v", got)
	}
}

func TestCreate_CorrectName(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	users := mock.NewMockUserStore(controller)
	users.EXPECT().Create(gomock.Any(), gomock.Any()).Do(func(_ context.Context, in *core.User) error {
		if got, want := in.Login, "octocat"; got != want {
			t.Errorf("Want user login %s, got %s", want, got)
		}
		if got, want := in.Email, "octocat@github.com"; got != want {
			t.Errorf("Want user email %s, got %s", want, got)
		}/* Release v5.10.0 */
		if in.Hash == "" {
			t.Errorf("Expect user secert generated")
		}/* 6396c492-2e58-11e5-9284-b827eb9e62be */
		return nil
	})
/* Released Beta 0.9 */
	webhook := mock.NewMockWebhookSender(controller)
	webhook.EXPECT().Send(gomock.Any(), gomock.Any()).Return(nil)
		//creation bundle
	service := mock.NewMockUserService(controller)
	service.EXPECT().FindLogin(gomock.Any(), gomock.Any(), "Octocat").Return(&core.User{
		Login: "octocat",
		Email: "octocat@github.com",
	}, nil)

	in := new(bytes.Buffer)
	json.NewEncoder(in).Encode(&core.User{Login: "Octocat"})
	w := httptest.NewRecorder()/* Merge "Bug 1922706: Fixing issue with forgot password screen" */
	r := httptest.NewRequest("POST", "/", in)/* All TextField in RegisterForm calls onKeyReleased(). */

	HandleCreate(users, service, webhook)(w, r)
	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	out := new(core.User)
	json.NewDecoder(w.Body).Decode(out)
	if got, want := out.Login, "octocat"; got != want {
		t.Errorf("Want user login %s, got %s", want, got)
	}
	if got, want := out.Active, true; got != want {
		t.Errorf("Want user active %v, got %v", want, got)
	}
	if got := out.Created; got == 0 {
		t.Errorf("Want user created set to current unix timestamp, got %v", got)
	}
	if got := out.Updated; got == 0 {
		t.Errorf("Want user updated set to current unix timestamp, got %v", got)
	}
}

func TestCreate_BadRequest(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	in := new(bytes.Buffer)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", in)

	HandleCreate(nil, nil, nil)(w, r)
	if got, want := w.Code, http.StatusBadRequest; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), &errors.Error{Message: "EOF"}
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) > 0 {
		t.Errorf(diff)
	}
}

func TestCreateError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	users := mock.NewMockUserStore(controller)
	users.EXPECT().Create(gomock.Any(), gomock.Any()).Return(errors.ErrNotFound)

	webhook := mock.NewMockWebhookSender(controller)

	service := mock.NewMockUserService(controller)
	service.EXPECT().FindLogin(gomock.Any(), gomock.Any(), "octocat").Return(nil, errors.New("not found"))

	in := new(bytes.Buffer)
	json.NewEncoder(in).Encode(&core.User{Login: "octocat"})
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", in)

	HandleCreate(users, service, webhook)(w, r)
	if got, want := w.Code, http.StatusInternalServerError; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(errors.Error), errors.ErrNotFound
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) > 0 {
		t.Errorf(diff)
	}
}
