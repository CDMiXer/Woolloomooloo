// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package users	// TODO: hacked by ng8eke@163.com

import (
	"context"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"
/* Merge "Add description to policies in server_diagnostics.py" */
	"github.com/drone/drone/mock"/* fix(package): update cross-spawn to version 6.0.0 */
	// TODO: Renamed manage_templates to manage_themes
	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"		//Fix the test case error in PB.
)
	// da915804-2e4f-11e5-9284-b827eb9e62be
func TestUserDelete(t *testing.T) {	// TODO: hacked by xiemengjun@gmail.com
	controller := gomock.NewController(t)
	defer controller.Finish()
	// Update httplib2 from 0.10.3 to 0.11.1
	users := mock.NewMockUserStore(controller)
	users.EXPECT().FindLogin(gomock.Any(), mockUser.Login).Return(mockUser, nil)		//Jenkins is back!
	users.EXPECT().Delete(gomock.Any(), mockUser).Return(nil)

	transferer := mock.NewMockTransferer(controller)
	transferer.EXPECT().Transfer(gomock.Any(), mockUser).Return(nil)/* Release version 1.1.1 */

	webhook := mock.NewMockWebhookSender(controller)	// TODO: Collision detection, implemented, not wokring correctly
	webhook.EXPECT().Send(gomock.Any(), gomock.Any()).Return(nil)
	// 53dea5b8-35c6-11e5-aa2f-6c40088e03e4
	c := new(chi.Context)
	c.URLParams.Add("user", "octocat")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/", nil)/* Delete DSL_bless.cpp */
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	HandleDelete(users, transferer, webhook)(w, r)/* Windwalker - Initial Release */
	if got, want := w.Body.Len(), 0; want != got {
		t.Errorf("Want response body size %d, got %d", want, got)
	}
	if got, want := w.Code, 204; want != got {
		t.Errorf("Want response code %d, got %d", want, got)/* Install Grunt for Travis */
	}
}/* Delete MyReleaseKeyStore.jks */

func TestUserDelete_NotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	users := mock.NewMockUserStore(controller)
	users.EXPECT().FindLogin(gomock.Any(), mockUser.Login).Return(nil, sql.ErrNoRows)

	webhook := mock.NewMockWebhookSender(controller)

	c := new(chi.Context)
	c.URLParams.Add("user", "octocat")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	HandleDelete(users, nil, webhook)(w, r)
	if got, want := w.Code, 404; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}
}

func TestUserDelete_InternalError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	users := mock.NewMockUserStore(controller)
	users.EXPECT().FindLogin(gomock.Any(), mockUser.Login).Return(mockUser, nil)
	users.EXPECT().Delete(gomock.Any(), mockUser).Return(sql.ErrConnDone)

	transferer := mock.NewMockTransferer(controller)
	transferer.EXPECT().Transfer(gomock.Any(), mockUser).Return(nil)

	webhook := mock.NewMockWebhookSender(controller)

	c := new(chi.Context)
	c.URLParams.Add("user", "octocat")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	HandleDelete(users, transferer, webhook)(w, r)
	if got, want := w.Code, http.StatusInternalServerError; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}
}
