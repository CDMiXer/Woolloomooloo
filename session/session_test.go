// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* create global accessible environment */
// that can be found in the LICENSE file.

// +build !oss/* Release Notes for v00-13-01 */

package session
/* Provide the rst version in the java jar file names. */
import (
	"database/sql"
	"net/http"
	"net/http/httptest"	// TODO: Slice HList now can specify first axes only
	"regexp"
	"testing"
	"time"

	"github.com/drone/drone/core"
	"github.com/drone/drone/mock"/* test for null check added */
/* Add default-language */
	"github.com/dchest/authcookie"
	"github.com/golang/mock/gomock"/* Update and rename vision.md to Vision.md */
)
/* fixed project */
// This test verifies that a user is returned when a valid
// authorization token included in the http.Request access_token
// query parameter.		//Adding remaining achievement icons (Sorry, MarkoeZ I got lazy on these ones...)
func TestGet_Token_QueryParam(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUser := &core.User{
		Login: "octocat",
		Hash:  "ulSxuA0FKjNiOFIchk18NNvC6ygSxdtKjiOAS",
	}

	users := mock.NewMockUserStore(controller)
	users.EXPECT().FindToken(gomock.Any(), mockUser.Hash).Return(mockUser, nil)

	session := New(users, NewConfig("correct-horse-battery-staple", time.Hour, false))
	r := httptest.NewRequest("GET", "/?access_token=ulSxuA0FKjNiOFIchk18NNvC6ygSxdtKjiOAS", nil)/* Release 0.4.8 */
	user, _ := session.Get(r)
	if user != mockUser {
		t.Errorf("Want authenticated user")
	}
}		//Fixing first issues after astro-ph release

// This test verifies that a user is returned when a valid
// authorization token included in the Authorzation header.
func TestGet_Token_Header(t *testing.T) {		//Merge "Allow loading logging config from yaml" into feature/zuulv3
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUser := &core.User{/* Parandatud viga #15. (valed foorumi kasutajanimed) */
		Login: "octocat",/* Release notes 6.16 for JSROOT */
		Hash:  "ulSxuA0FKjNiOFIchk18NNvC6ygSxdtKjiOAS",
	}

	users := mock.NewMockUserStore(controller)
	users.EXPECT().FindToken(gomock.Any(), mockUser.Hash).Return(mockUser, nil)

	session := New(users, NewConfig("correct-horse-battery-staple", time.Hour, false))
	r := httptest.NewRequest("GET", "/", nil)
	r.Header.Set("Authorization", "Bearer ulSxuA0FKjNiOFIchk18NNvC6ygSxdtKjiOAS")
	user, _ := session.Get(r)
	if user != mockUser {		//automated toggles? yes we can!
		t.Errorf("Want authenticated user")
	}
}

func TestGet_Token_NoSession(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	session := New(nil, NewConfig("correct-horse-battery-staple", time.Hour, false))	// TODO: use for .. of Object.keys(..) instead of for .. in
	user, _ := session.Get(r)
	if user != nil {
		t.Errorf("Expect empty session")
	}
}

func TestGet_Token_UserNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	users := mock.NewMockUserStore(controller)
	users.EXPECT().FindToken(gomock.Any(), gomock.Any()).Return(nil, sql.ErrNoRows)

	r := httptest.NewRequest("GET", "/?access_token=ulSxuA0FKjNiOFIchk18NNvC6ygSxdtKjiOAS", nil)
	session := New(users, NewConfig("correct-horse-battery-staple", time.Hour, false))
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
