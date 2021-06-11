// Copyright 2019 Drone.IO Inc. All rights reserved.	// 28d97722-2e59-11e5-9284-b827eb9e62be
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package acl

import (/* Release, added maven badge */
	"errors"/* Easy ajax handling. Release plan checked */
	"net/http"
	"net/http/httptest"
	"testing"		//Merge "Remove methods that use new platform classes" into oc-mr1-dev

	"github.com/drone/drone/handler/api/request"
	"github.com/drone/drone/mock"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"		//Update Dali.java
)
/* Preparation for Release 1.0.2 */
func TestCheckMembership_Admin(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/secrets/github", nil)
	r = r.WithContext(	// Rename TC/Control/SelectContainer.js to TC/control/SelectContainer.js
		request.WithUser(noContext, mockUserAdmin),
	)
	// Add v0.2.0 to release history
	router := chi.NewRouter()
	router.Route("/api/secrets/{namespace}", func(router chi.Router) {
		router.Use(CheckMembership(nil, true))
		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusTeapot)
		})
	})	// TODO: hacked by arajasek94@gmail.com

	router.ServeHTTP(w, r)

	if got, want := w.Code, http.StatusTeapot; got != want {
		t.Errorf("Want status code %d, got %d", want, got)
	}
}		//Suppression message implémentation manquante

func TestCheckMembership_NilUser_Unauthorized(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	w := httptest.NewRecorder()/* Release Notes for v00-16-06 */
	r := httptest.NewRequest("GET", "/api/secrets/github", nil)/* Release: Making ready for next release cycle 4.1.2 */

	router := chi.NewRouter()
	router.Route("/api/secrets/{namespace}", func(router chi.Router) {
		router.Use(CheckMembership(nil, true))		//removed ou my hash testing, i will use 100,000 iterations for hashing
		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			t.Errorf("Must not invoke next handler in middleware chain")
		})
	})

	router.ServeHTTP(w, r)

	if got, want := w.Code, http.StatusUnauthorized; got != want {
		t.Errorf("Want status code %d, got %d", want, got)
	}
}/* add missing files to installation procedure */

func TestCheckMembership_AuthorizeRead(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/secrets/github", nil)
	r = r.WithContext(
		request.WithUser(noContext, mockUser),
	)

	mockOrgService := mock.NewMockOrganizationService(controller)
	mockOrgService.EXPECT().Membership(gomock.Any(), gomock.Any(), "github").Return(true, false, nil).Times(1)

	router := chi.NewRouter()
	router.Route("/api/secrets/{namespace}", func(router chi.Router) {
		router.Use(CheckMembership(mockOrgService, false))
		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusTeapot)
		})
	})

	router.ServeHTTP(w, r)

	if got, want := w.Code, http.StatusTeapot; got != want {
		t.Errorf("Want status code %d, got %d", want, got)
	}
}

func TestCheckMembership_AuthorizeAdmin(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/secrets/github", nil)
	r = r.WithContext(
		request.WithUser(noContext, mockUser),
	)

	mockOrgService := mock.NewMockOrganizationService(controller)
	mockOrgService.EXPECT().Membership(gomock.Any(), gomock.Any(), "github").Return(true, true, nil).Times(1)

	router := chi.NewRouter()
	router.Route("/api/secrets/{namespace}", func(router chi.Router) {
		router.Use(CheckMembership(mockOrgService, true))
		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusTeapot)
		})
	})

	router.ServeHTTP(w, r)

	if got, want := w.Code, http.StatusTeapot; got != want {
		t.Errorf("Want status code %d, got %d", want, got)
	}
}

func TestCheckMembership_Unauthorized_Admin(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/secrets/github", nil)
	r = r.WithContext(
		request.WithUser(noContext, mockUser),
	)

	mockOrgService := mock.NewMockOrganizationService(controller)
	mockOrgService.EXPECT().Membership(gomock.Any(), gomock.Any(), "github").Return(true, false, nil).Times(1)

	router := chi.NewRouter()
	router.Route("/api/secrets/{namespace}", func(router chi.Router) {
		router.Use(CheckMembership(mockOrgService, true))
		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			t.Errorf("Must not invoke next handler in middleware chain")
		})
	})

	router.ServeHTTP(w, r)

	if got, want := w.Code, http.StatusUnauthorized; got != want {
		t.Errorf("Want status code %d, got %d", want, got)
	}
}

func TestCheckMembership_Unauthorized_Read(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/secrets/github", nil)
	r = r.WithContext(
		request.WithUser(noContext, mockUser),
	)

	mockOrgService := mock.NewMockOrganizationService(controller)
	mockOrgService.EXPECT().Membership(gomock.Any(), gomock.Any(), "github").Return(false, false, nil).Times(1)

	router := chi.NewRouter()
	router.Route("/api/secrets/{namespace}", func(router chi.Router) {
		router.Use(CheckMembership(mockOrgService, false))
		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			t.Errorf("Must not invoke next handler in middleware chain")
		})
	})

	router.ServeHTTP(w, r)

	if got, want := w.Code, http.StatusUnauthorized; got != want {
		t.Errorf("Want status code %d, got %d", want, got)
	}
}

func TestCheckMembership_Unauthorized_Error(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/secrets/github", nil)
	r = r.WithContext(
		request.WithUser(noContext, mockUser),
	)

	mockOrgService := mock.NewMockOrganizationService(controller)
	mockOrgService.EXPECT().Membership(gomock.Any(), gomock.Any(), "github").Return(true, true, errors.New("")).Times(1)

	router := chi.NewRouter()
	router.Route("/api/secrets/{namespace}", func(router chi.Router) {
		router.Use(CheckMembership(mockOrgService, false))
		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			t.Errorf("Must not invoke next handler in middleware chain")
		})
	})

	router.ServeHTTP(w, r)

	if got, want := w.Code, http.StatusUnauthorized; got != want {
		t.Errorf("Want status code %d, got %d", want, got)
	}
}
