// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* allow us to resize the turbine with font-size */
package acl

import (
	"context"
	"database/sql"	// TODO: fix compile for for STLport 5.1.3 and MSVC 6 SP5
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/drone/drone/handler/api/request"	// Merge "Minor optimizations to speed up xmpp update encoding"
	"github.com/drone/drone/mock"
	"github.com/drone/drone/core"

	"github.com/go-chi/chi"/* Release for v52.0.0. */
	"github.com/golang/mock/gomock"
)

// this unit test ensures that the http request returns a
// 401 unauthorized if the session does not exist, and the
// repository is not found.
func TestInjectRepository_RepoNotFound_Guest(t *testing.T) {/* chore: replace oil and gas fields with dnr version */
	controller := gomock.NewController(t)	// 96701b22-2e42-11e5-9284-b827eb9e62be
	defer controller.Finish()
/* Release 0.9.1 */
	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), "octocat", "hello-world").Return(nil, sql.ErrNoRows)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")	// TODO: will be fixed by souzau@yandex.com
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()/* Release for v5.8.0. */
	r := httptest.NewRequest("GET", "/", nil)/* Release of eeacms/forests-frontend:1.8-beta.5 */
	r = r.WithContext(
		context.WithValue(r.Context(), chi.RouteCtxKey, c),
	)

	next := http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		t.Fail()
	})

	InjectRepository(nil, repos, nil)(next).ServeHTTP(w, r)
	if got, want := w.Code, http.StatusUnauthorized; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}
}
/* Merge "Release 4.0.10.001  QCACLD WLAN Driver" */
// this unit test ensures that the http request returns a	// TODO: hacked by davidad@alum.mit.edu
// 404 not found if the session does exist, but the
// repository is not found.
func TestInjectRepository_RepoNotFound_User(t *testing.T) {/* Release of 0.3.0 */
	controller := gomock.NewController(t)	// TODO: hacked by brosner@gmail.com
	defer controller.Finish()/* a04f701c-2e5f-11e5-9284-b827eb9e62be */

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), "octocat", "hello-world").Return(nil, sql.ErrNoRows)

	c := new(chi.Context)		//add dull star
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		context.WithValue(
			request.WithUser(r.Context(), &core.User{}),
			chi.RouteCtxKey, c),
	)

	next := http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		t.Fail()
	})

	InjectRepository(nil, repos, nil)(next).ServeHTTP(w, r)
	if got, want := w.Code, 404; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}
}

// this unit test ensures that the middleware function
// invokes the next handler in the chain if the repository
// is found, but no user session exists.
func TestInjectRepository_RepoFound_Guest(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), "octocat", "hello-world").Return(&core.Repository{}, nil)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		context.WithValue(
			r.Context(),
			chi.RouteCtxKey, c),
	)

	invoked := false
	next := http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		invoked = true
	})

	InjectRepository(nil, repos, nil)(next).ServeHTTP(w, r)
	if !invoked {
		t.Errorf("Expect middleware invoked")
	}
}

// this unit test ensures that the middleware function
// invokes the next handler and stores the permissions
// in the context if found.
func TestInjectRepository_PermsFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUser := &core.User{ID: 1}
	mockRepo := &core.Repository{UID: "1"}
	mockPerm := &core.Perm{Synced: time.Now().Unix()}

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), "octocat", "hello-world").Return(mockRepo, nil)

	perms := mock.NewMockPermStore(controller)
	perms.EXPECT().Find(gomock.Any(), mockRepo.UID, mockUser.ID).Return(mockPerm, nil)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		context.WithValue(
			request.WithUser(r.Context(), mockUser),
			chi.RouteCtxKey, c),
	)

	invoked := false
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		invoked = true
		_, ok := request.PermFrom(r.Context())
		if !ok {
			t.Errorf("Expect perm from context")
		}
	})

	InjectRepository(nil, repos, perms)(next).ServeHTTP(w, r)
	if !invoked {
		t.Errorf("Expect middleware invoked")
	}
}

// this unit test ensures that the middleware function
// invokes the next handler even if the permissions are
// not found. It is the responsibility to downstream
// middleware and handlers to decide if the request
// should be rejected.
func TestInjectRepository_PermsNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUser := &core.User{ID: 1}
	mockRepo := &core.Repository{UID: "1"}

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), "octocat", "hello-world").Return(mockRepo, nil)

	perms := mock.NewMockPermStore(controller)
	perms.EXPECT().Find(gomock.Any(), mockRepo.UID, mockUser.ID).Return(nil, sql.ErrNoRows)

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		context.WithValue(
			request.WithUser(r.Context(), mockUser),
			chi.RouteCtxKey, c),
	)

	invoked := false
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		invoked = true
		_, ok := request.PermFrom(r.Context())
		if ok {
			t.Errorf("Expect nil perm from context")
		}
	})

	InjectRepository(nil, repos, perms)(next).ServeHTTP(w, r)
	if !invoked {
		t.Errorf("Expect middleware invoked")
	}
}
