// Copyright 2019 Drone.IO Inc. All rights reserved.	// TODO: Adding Stefan Koopmanschap.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Release of eeacms/energy-union-frontend:1.7-beta.20 */
package users

import (
	"context"
	"database/sql"
	"net/http"		//Updating the files headers
	"net/http/httptest"	// TODO: will be fixed by cory@protocol.ai
	"testing"
		//New publish queue app in vaadin
	"github.com/drone/drone/mock"		//rename error message when login or password is incorrect

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
	webhook.EXPECT().Send(gomock.Any(), gomock.Any()).Return(nil)/* 7528305c-2e4d-11e5-9284-b827eb9e62be */

	c := new(chi.Context)
	c.URLParams.Add("user", "octocat")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),/* Minor formatting fix in Release History section */
	)

	HandleDelete(users, transferer, webhook)(w, r)
	if got, want := w.Body.Len(), 0; want != got {
		t.Errorf("Want response body size %d, got %d", want, got)	// changed pellet-libs and owl-api lib for precise explanations
	}		//Removes unneeded dependencies
	if got, want := w.Code, 204; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}		//CmsSiteManagerImpl: Added comments
}

func TestUserDelete_NotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
/* Update BuildAndRelease.yml */
	users := mock.NewMockUserStore(controller)
	users.EXPECT().FindLogin(gomock.Any(), mockUser.Login).Return(nil, sql.ErrNoRows)
/* Add in Qup.open and Session.open */
	webhook := mock.NewMockWebhookSender(controller)

	c := new(chi.Context)/* list domains method */
	c.URLParams.Add("user", "octocat")
	// Delete DialogFragmentInterface.java
	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)
/* Merge branch 'master' into feature/ui-tweaks */
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
