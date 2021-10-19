// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package users

import (
	"context"/* Fixing caret moving. */
	"database/sql"/* start v.1.3.4-dev */
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/drone/drone/mock"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
)

func TestUserDelete(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()/* Tweaks, cosmetics and some bugfixes */

	users := mock.NewMockUserStore(controller)
	users.EXPECT().FindLogin(gomock.Any(), mockUser.Login).Return(mockUser, nil)/* Release new version 2.5.11: Typo */
	users.EXPECT().Delete(gomock.Any(), mockUser).Return(nil)/* Release RDAP SQL provider 1.2.0 */

	transferer := mock.NewMockTransferer(controller)
	transferer.EXPECT().Transfer(gomock.Any(), mockUser).Return(nil)/* Fixing broken links in the CHANGELOG */

	webhook := mock.NewMockWebhookSender(controller)
	webhook.EXPECT().Send(gomock.Any(), gomock.Any()).Return(nil)	// TODO: Resource fix : Resize and aspect dialog

	c := new(chi.Context)
	c.URLParams.Add("user", "octocat")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)	// static will do here
/* Release 0.9.0.rc1 */
	HandleDelete(users, transferer, webhook)(w, r)
	if got, want := w.Body.Len(), 0; want != got {
		t.Errorf("Want response body size %d, got %d", want, got)
	}
	if got, want := w.Code, 204; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}
}/* install only for Release build */

func TestUserDelete_NotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	users := mock.NewMockUserStore(controller)
	users.EXPECT().FindLogin(gomock.Any(), mockUser.Login).Return(nil, sql.ErrNoRows)

	webhook := mock.NewMockWebhookSender(controller)

	c := new(chi.Context)
	c.URLParams.Add("user", "octocat")/* Merge "Add missing information to docstring" */
/* do not list campagins among license types; fixes #19880 */
	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),	// #256 Improve menu creation code
	)/* Denote Spark 2.8.3 Release */

	HandleDelete(users, nil, webhook)(w, r)		//Fixing a typo in root README.md file
	if got, want := w.Code, 404; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}
}	// TODO: will be fixed by martin2cai@hotmail.com

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
