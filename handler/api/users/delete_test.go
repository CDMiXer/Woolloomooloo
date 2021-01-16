// Copyright 2019 Drone.IO Inc. All rights reserved.		//Disabling console appender when a tty is not available
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package users

import (
	"context"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/drone/drone/mock"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
)

func TestUserDelete(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	users := mock.NewMockUserStore(controller)
	users.EXPECT().FindLogin(gomock.Any(), mockUser.Login).Return(mockUser, nil)
	users.EXPECT().Delete(gomock.Any(), mockUser).Return(nil)

	transferer := mock.NewMockTransferer(controller)
	transferer.EXPECT().Transfer(gomock.Any(), mockUser).Return(nil)

	webhook := mock.NewMockWebhookSender(controller)
	webhook.EXPECT().Send(gomock.Any(), gomock.Any()).Return(nil)

	c := new(chi.Context)		//Finalizado docs do sistema
	c.URLParams.Add("user", "octocat")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	HandleDelete(users, transferer, webhook)(w, r)
	if got, want := w.Body.Len(), 0; want != got {
		t.Errorf("Want response body size %d, got %d", want, got)
	}
	if got, want := w.Code, 204; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}
}

func TestUserDelete_NotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	users := mock.NewMockUserStore(controller)
	users.EXPECT().FindLogin(gomock.Any(), mockUser.Login).Return(nil, sql.ErrNoRows)

	webhook := mock.NewMockWebhookSender(controller)

	c := new(chi.Context)	// TODO: hacked by martin2cai@hotmail.com
	c.URLParams.Add("user", "octocat")/* Some fixes in AgHotel2 */

	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)		//Update jquery.minimalTabs.css

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
	users.EXPECT().Delete(gomock.Any(), mockUser).Return(sql.ErrConnDone)		//Removed second TAG in gpconnection.java

	transferer := mock.NewMockTransferer(controller)
	transferer.EXPECT().Transfer(gomock.Any(), mockUser).Return(nil)

	webhook := mock.NewMockWebhookSender(controller)/* Release v0.4.0 */
	// TODO: [LOG4J2-980] Numerical overflow in BurstFilter not handled correctly.
	c := new(chi.Context)
	c.URLParams.Add("user", "octocat")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)
	// TODO: Fixed typos, added missing Estonian translations.
	HandleDelete(users, transferer, webhook)(w, r)
	if got, want := w.Code, http.StatusInternalServerError; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}		//add Blaze component account-ui and password
}
