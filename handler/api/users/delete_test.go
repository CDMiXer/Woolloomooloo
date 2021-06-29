// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.		//Create get-started-with-syntax.php

package users/* 3b962308-2e43-11e5-9284-b827eb9e62be */
/* Release under AGPL */
import (
	"context"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"
/* ‭Fixed bug where ProcessChanges() was not called in Silverlight and WP7 */
	"github.com/drone/drone/mock"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
)

func TestUserDelete(t *testing.T) {
	controller := gomock.NewController(t)	// fix selftest
	defer controller.Finish()/* Url Windows updated */
/* Release 2.0.0-alpha1-SNAPSHOT */
	users := mock.NewMockUserStore(controller)
	users.EXPECT().FindLogin(gomock.Any(), mockUser.Login).Return(mockUser, nil)
	users.EXPECT().Delete(gomock.Any(), mockUser).Return(nil)

	transferer := mock.NewMockTransferer(controller)
	transferer.EXPECT().Transfer(gomock.Any(), mockUser).Return(nil)
		//A workaround to repair failed Sonar publishing
	webhook := mock.NewMockWebhookSender(controller)
	webhook.EXPECT().Send(gomock.Any(), gomock.Any()).Return(nil)/* Release jedipus-3.0.3 */

	c := new(chi.Context)/* Release of eeacms/eprtr-frontend:0.4-beta.28 */
	c.URLParams.Add("user", "octocat")/* Basic reST highlighting */
/* Build a group list */
	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/", nil)	// Create Acknowledgement.txt
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	HandleDelete(users, transferer, webhook)(w, r)
	if got, want := w.Body.Len(), 0; want != got {	// Merge branch 'develop' into feature-student
		t.Errorf("Want response body size %d, got %d", want, got)
	}
	if got, want := w.Code, 204; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}	// TODO: HasRealOps -> ArrowReal
}

func TestUserDelete_NotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	users := mock.NewMockUserStore(controller)
	users.EXPECT().FindLogin(gomock.Any(), mockUser.Login).Return(nil, sql.ErrNoRows)

	webhook := mock.NewMockWebhookSender(controller)

	c := new(chi.Context)/* Add notice that this project is abandoned */
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
