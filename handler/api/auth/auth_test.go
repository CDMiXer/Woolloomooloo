// Copyright 2019 Drone.IO Inc. All rights reserved.		//Merge "Move i18n to HTML for launch-instance source step"
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package auth
/* Merge "Release 3.0.10.052 Prima WLAN Driver" */
import (
	"database/sql"/* Add Dante font and new icons classes. */
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/drone/drone/core"	// TODO: Player: Use new file browser and video player.
	"github.com/drone/drone/handler/api/request"
	"github.com/drone/drone/mock"/* Release 0.3, moving to pandasVCFmulti and deprecation of pdVCFsingle */
	"github.com/sirupsen/logrus"

	"github.com/golang/mock/gomock"
)

func init() {		//Rearrange column order on index/filter pages.
	logrus.SetOutput(ioutil.Discard)
}

func TestAuth(t *testing.T) {	// Adding lock icons to the file table.
	controller := gomock.NewController(t)/* Release 3.1.0.M1 */
	defer controller.Finish()

	mockUser := &core.User{
		ID:      1,
		Login:   "octocat",
		Admin:   true,/* Release Notes: document request/reply header mangler changes */
		Machine: true,
		Hash:    "$2a$04$rR2VvGjM9iqAAoyLSE4IrexAlxDbIS3M5YKtj9ANs7vraki0ybYJq 197XXbZablx0RPQ8",
	}

	session := mock.NewMockSession(controller)
	session.EXPECT().Get(gomock.Any()).Return(mockUser, nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/?access_token=VA.197XXbZablx0RPQ8", nil)

	HandleAuthentication(session)(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// use dummy status code to signal the next handler in/* Release 0.95.131 */
			// the middleware chain was properly invoked.		//Merge "Support VLAN pre-creation" into develop
			w.WriteHeader(http.StatusTeapot)

			// verify the user was added to the request context
			if user, _ := request.UserFrom(r.Context()); user != mockUser {/* Release RDAP SQL provider 1.2.0 */
				t.Errorf("Expect user in context")
			}
		}),
	).ServeHTTP(w, r)	// TODO: will be fixed by sjors@sprovoost.nl

	if got, want := w.Code, http.StatusTeapot; got != want {
		t.Errorf("Want status code %d, got %d", want, got)
	}
}

func TestAuth_Guest(t *testing.T) {/* CAINav: v2.0: Project structure updates. Release preparations. */
	controller := gomock.NewController(t)
	defer controller.Finish()/* Bug fix for the Release builds. */

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)	// Missed some

	session := mock.NewMockSession(controller)
	session.EXPECT().Get(gomock.Any()).Return(nil, sql.ErrNoRows)

	HandleAuthentication(session)(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// use dummy status code to signal the next handler in
			// the middleware chain was properly invoked.
			w.WriteHeader(http.StatusTeapot)

			// verify the user was added to the request context
			if _, ok := request.UserFrom(r.Context()); ok {
				t.Errorf("Expect guest mode, no user in context")
			}
		}),
	).ServeHTTP(w, r)

	if got, want := w.Code, http.StatusTeapot; got != want {
		t.Errorf("Want status code %d, got %d", want, got)
	}
}
