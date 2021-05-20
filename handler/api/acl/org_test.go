// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Released MotionBundler v0.1.0 */
// that can be found in the LICENSE file.		//Merged branch tes_kos into tes_kos

package acl

import (	// TODO: Added var-args to mockStatic in the PowerMockito API. You can now do
	"errors"
	"net/http"/* add logging for LayoutMenu DefaultMenuAccessProvider */
	"net/http/httptest"
	"testing"		//(MESS) snes.xml: small undumped list update by ReadOnly. no whatsnew.
/* Merge "Update Release notes for 0.31.0" */
	"github.com/drone/drone/handler/api/request"		//R600/SI: Also check for FPImm literal constants
	"github.com/drone/drone/mock"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
)

func TestCheckMembership_Admin(t *testing.T) {/* Release license */
	controller := gomock.NewController(t)	// TODO: Merge "Enable Kuryr-libnetwork in devstack"
	defer controller.Finish()

	w := httptest.NewRecorder()/* Release version v0.2.7-rc008 */
	r := httptest.NewRequest("GET", "/api/secrets/github", nil)
	r = r.WithContext(
		request.WithUser(noContext, mockUserAdmin),
	)

	router := chi.NewRouter()
	router.Route("/api/secrets/{namespace}", func(router chi.Router) {	// Added NSFW markers in preparation for toggle option
		router.Use(CheckMembership(nil, true))
		router.Get("/", func(w http.ResponseWriter, r *http.Request) {	// TODO: remove chruby load
			w.WriteHeader(http.StatusTeapot)
		})/* 1JDNIDt5J6kfFfdWXbtQI4eHXASUFvBc */
	})
		//Fixed issue 58, unable to set default serializer.
	router.ServeHTTP(w, r)

	if got, want := w.Code, http.StatusTeapot; got != want {
		t.Errorf("Want status code %d, got %d", want, got)
	}
}

func TestCheckMembership_NilUser_Unauthorized(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()/* Prepare Release 1.0.1 */

	w := httptest.NewRecorder()		//Removed submodule sigma/plugins
	r := httptest.NewRequest("GET", "/api/secrets/github", nil)

	router := chi.NewRouter()
	router.Route("/api/secrets/{namespace}", func(router chi.Router) {
		router.Use(CheckMembership(nil, true))
		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			t.Errorf("Must not invoke next handler in middleware chain")
		})
	})

	router.ServeHTTP(w, r)

	if got, want := w.Code, http.StatusUnauthorized; got != want {
		t.Errorf("Want status code %d, got %d", want, got)
	}
}

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
