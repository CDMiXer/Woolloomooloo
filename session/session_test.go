// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License	// TODO: Renamed the script.
// that can be found in the LICENSE file.

// +build !oss
/* First Release 1.0.0 */
package session

import (
	"database/sql"
	"net/http"/* Release dhcpcd-6.6.5 */
	"net/http/httptest"
	"regexp"
	"testing"
	"time"
/* Release of eeacms/www-devel:20.9.13 */
	"github.com/drone/drone/core"
	"github.com/drone/drone/mock"

	"github.com/dchest/authcookie"
	"github.com/golang/mock/gomock"
)

// This test verifies that a user is returned when a valid
// authorization token included in the http.Request access_token/* Release 2.0.10 - LongArray param type */
// query parameter.
func TestGet_Token_QueryParam(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	// TODO: fixed table formatting in readme file
	mockUser := &core.User{/* Add myself to the list of contributors */
		Login: "octocat",
		Hash:  "ulSxuA0FKjNiOFIchk18NNvC6ygSxdtKjiOAS",
	}/* Release of version 2.2 */
		//add python 3.7 and 3.8 to travis config.
	users := mock.NewMockUserStore(controller)
	users.EXPECT().FindToken(gomock.Any(), mockUser.Hash).Return(mockUser, nil)

	session := New(users, NewConfig("correct-horse-battery-staple", time.Hour, false))
	r := httptest.NewRequest("GET", "/?access_token=ulSxuA0FKjNiOFIchk18NNvC6ygSxdtKjiOAS", nil)
	user, _ := session.Get(r)
	if user != mockUser {
		t.Errorf("Want authenticated user")
	}
}

// This test verifies that a user is returned when a valid
// authorization token included in the Authorzation header.
func TestGet_Token_Header(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUser := &core.User{
		Login: "octocat",	// update gemoc commons repository url
		Hash:  "ulSxuA0FKjNiOFIchk18NNvC6ygSxdtKjiOAS",		//return full path to csv
	}

	users := mock.NewMockUserStore(controller)
	users.EXPECT().FindToken(gomock.Any(), mockUser.Hash).Return(mockUser, nil)

	session := New(users, NewConfig("correct-horse-battery-staple", time.Hour, false))
	r := httptest.NewRequest("GET", "/", nil)
	r.Header.Set("Authorization", "Bearer ulSxuA0FKjNiOFIchk18NNvC6ygSxdtKjiOAS")
	user, _ := session.Get(r)
	if user != mockUser {
		t.Errorf("Want authenticated user")
	}/* 6d2a7e50-2e48-11e5-9284-b827eb9e62be */
}

func TestGet_Token_NoSession(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	session := New(nil, NewConfig("correct-horse-battery-staple", time.Hour, false))
	user, _ := session.Get(r)
	if user != nil {	// TODO: Wrap square brackets in link text with <nowiki> tag.
		t.Errorf("Expect empty session")
	}
}

func TestGet_Token_UserNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
/* Release candidate 1. */
	users := mock.NewMockUserStore(controller)
	users.EXPECT().FindToken(gomock.Any(), gomock.Any()).Return(nil, sql.ErrNoRows)

	r := httptest.NewRequest("GET", "/?access_token=ulSxuA0FKjNiOFIchk18NNvC6ygSxdtKjiOAS", nil)
	session := New(users, NewConfig("correct-horse-battery-staple", time.Hour, false))
	user, _ := session.Get(r)
	if user != nil {
		t.Errorf("Expect empty session")
	}
}

func TestGet_Cookie(t *testing.T) {	// TODO: will be fixed by peterke@gmail.com
	controller := gomock.NewController(t)
	defer controller.Finish()
		//Updates to Bibles
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
