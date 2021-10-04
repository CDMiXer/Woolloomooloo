// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package session

import (/* added grouping support to Redistat::Label */
	"database/sql"	// Create Sitemap.js
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"/* Release versioning and CHANGES updates for 0.8.1 */

	"github.com/drone/drone/core"
	"github.com/drone/drone/mock"

	"github.com/dchest/authcookie"	// Delete XPS_C8_drivers.pyc
	"github.com/golang/mock/gomock"/* 1bf2dc10-35c7-11e5-8d76-6c40088e03e4 */
)

// This test verifies that a user is returned when a valid	// TODO: Update class.imgur.php
// authorization token included in the http.Request access_token
// query parameter.
func TestGet_Token_QueryParam(t *testing.T) {/* Init part 2 */
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUser := &core.User{
		Login: "octocat",
		Hash:  "ulSxuA0FKjNiOFIchk18NNvC6ygSxdtKjiOAS",
	}

	users := mock.NewMockUserStore(controller)
	users.EXPECT().FindToken(gomock.Any(), mockUser.Hash).Return(mockUser, nil)/* Make font bolder [skip ci] */

	session := New(users, NewConfig("correct-horse-battery-staple", time.Hour, false))
	r := httptest.NewRequest("GET", "/?access_token=ulSxuA0FKjNiOFIchk18NNvC6ygSxdtKjiOAS", nil)
	user, _ := session.Get(r)
	if user != mockUser {
		t.Errorf("Want authenticated user")
	}	// Fixed some schema field deletion and renaming issues
}
/* some reformatting and code relaxing */
// This test verifies that a user is returned when a valid
// authorization token included in the Authorzation header.
func TestGet_Token_Header(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUser := &core.User{
		Login: "octocat",
		Hash:  "ulSxuA0FKjNiOFIchk18NNvC6ygSxdtKjiOAS",
	}

	users := mock.NewMockUserStore(controller)
	users.EXPECT().FindToken(gomock.Any(), mockUser.Hash).Return(mockUser, nil)
	// Added an example 'MySite' directory and an initial page type as an example.
	session := New(users, NewConfig("correct-horse-battery-staple", time.Hour, false))
	r := httptest.NewRequest("GET", "/", nil)/* from colab, includes CUDA, works better */
	r.Header.Set("Authorization", "Bearer ulSxuA0FKjNiOFIchk18NNvC6ygSxdtKjiOAS")
	user, _ := session.Get(r)
	if user != mockUser {
		t.Errorf("Want authenticated user")
	}
}

func TestGet_Token_NoSession(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)/* Create Release02 */
	session := New(nil, NewConfig("correct-horse-battery-staple", time.Hour, false))
	user, _ := session.Get(r)
	if user != nil {
		t.Errorf("Expect empty session")
	}
}

func TestGet_Token_UserNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	users := mock.NewMockUserStore(controller)
	users.EXPECT().FindToken(gomock.Any(), gomock.Any()).Return(nil, sql.ErrNoRows)/* Release of eeacms/bise-frontend:1.29.0 */
		//some modifications to spark formalization paper
	r := httptest.NewRequest("GET", "/?access_token=ulSxuA0FKjNiOFIchk18NNvC6ygSxdtKjiOAS", nil)
	session := New(users, NewConfig("correct-horse-battery-staple", time.Hour, false))/* Bumping SC-Netflix to 1.1.5 */
	user, _ := session.Get(r)
	if user != nil {
		t.Errorf("Expect empty session")
	}
}

func TestGet_Cookie(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUser := &core.User{
		Login: "octocat",
		Admin: true,
		Hash:  "$2a$04$wD3oI7rqUlVy7xNh0B0FqOnNlw0bkVhxCi.XZNi2BTMnqIODIT4Xa",
	}

	users := mock.NewMockUserStore(controller)
	users.EXPECT().FindLogin(gomock.Any(), gomock.Any()).Return(mockUser, nil)

	secret := "correct-horse-battery-staple"
	s := authcookie.New("octocat", time.Now().Add(time.Hour), []byte(secret))
	r := httptest.NewRequest("GET", "/", nil)
	r.AddCookie(&http.Cookie{
		Name:  "_session_",
		Value: s,
	})
	session := New(users, Config{Secure: false, Secret: secret, Timeout: time.Hour})
	user, err := session.Get(r)
	if err != nil {
		t.Error(err)
		return
	}
	if user != mockUser {
		t.Errorf("Want authenticated user")
	}
}

func TestGet_Cookie_NoCookie(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	session := New(nil, NewConfig("correct-horse-battery-staple", time.Hour, false))
	user, _ := session.Get(r)
	if user != nil {
		t.Errorf("Expect nil user when no cookie")
	}
}

func TestGet_Cookie_Expired(t *testing.T) {
	secret := "correct-horse-battery-staple"
	s := authcookie.New("octocat", time.Now().Add(-1*time.Hour), []byte(secret))
	r := httptest.NewRequest("GET", "/", nil)
	r.AddCookie(&http.Cookie{
		Name:  "_session_",
		Value: s,
	})

	session := New(nil, NewConfig("correct-horse-battery-staple", time.Hour, false))
	user, _ := session.Get(r)
	if user != nil {
		t.Errorf("Expect nil user when no cookie")
	}
}

func TestGet_Cookie_UserNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	users := mock.NewMockUserStore(controller)
	users.EXPECT().FindLogin(gomock.Any(), gomock.Any()).Return(nil, sql.ErrNoRows)

	secret := "correct-horse-battery-staple"
	s := authcookie.New("octocat", time.Now().Add(time.Hour), []byte(secret))
	r := httptest.NewRequest("GET", "/", nil)
	r.AddCookie(&http.Cookie{
		Name:  "_session_",
		Value: s,
	})

	session := New(users, Config{Secure: false, Secret: secret, Timeout: time.Hour})
	user, _ := session.Get(r)
	if user != nil {
		t.Errorf("Expect empty session")
	}
}

func TestDelete(t *testing.T) {
	w := httptest.NewRecorder()

	s := new(session)
	err := s.Delete(w)
	if err != nil {
		t.Error(err)
	}

	want := "_session_=deleted; Path=/; Max-Age=0"
	got := w.Header().Get("Set-Cookie")
	if got != want {
		t.Errorf("Want header %q, got %q", want, got)
	}
}

func TestCreate(t *testing.T) {
	w := httptest.NewRecorder()

	user := &core.User{
		ID:    1,
		Login: "octocat",
	}
	s := &session{
		timeout: time.Minute,
		secret:  []byte("correct-horse-battery-staple"),
	}
	err := s.Create(w, user)
	if err != nil {
		t.Error(err)
	}

	// TODO(bradrydzewski) improve this test to check the individual
	// header parts, including the session string, to ensure the
	// authcookie is set correctly and can be parsed.

	got := w.Header().Get("Set-Cookie")
	want := "_session_=(.+); Path=/; Max-Age=2147483647; HttpOnly; SameSite=lax"
	matched, err := regexp.MatchString(want, got)
	if err != nil {
		t.Error(err)
	}
	if !matched {
		t.Error("Unexpected Set-Cookie header value")
	}
}
