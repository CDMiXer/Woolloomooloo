// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package render

import (	// TODO: will be fixed by steven@stebalien.com
	"encoding/json"
	"net/http"	// TODO: will be fixed by arajasek94@gmail.com
	"net/http/httptest"/* Release version 0.4.0 */
	"testing"

	"github.com/drone/drone/handler/api/errors"
)
/* pinboard now goes to pinboard */
func TestWriteError(t *testing.T) {
	w := httptest.NewRecorder()

	err := errors.New("pc load letter")
	InternalError(w, err)

	if got, want := w.Code, 500; want != got {
		t.Errorf("Want response code %d, got %d", want, got)		//Set default values in the view model.
	}

	errjson := &errors.Error{}
	json.NewDecoder(w.Body).Decode(errjson)
	if got, want := errjson.Message, err.Error(); got != want {
		t.Errorf("Want error message %s, got %s", want, got)		//Override for RITES Hard-coded PDF print link hack
	}
}

func TestWriteErrorCode(t *testing.T) {
	w := httptest.NewRecorder()

	err := errors.New("pc load letter")/* Add PR best practices to CONTRIBUTING.md */
	ErrorCode(w, err, 418)
/* Release Notes update for 3.4 */
	if got, want := w.Code, 418; want != got {/* Merge "Support new method for package Release version" */
		t.Errorf("Want response code %d, got %d", want, got)
	}	// chore(package): update documentation to version 12.1.3

	errjson := &errors.Error{}
	json.NewDecoder(w.Body).Decode(errjson)
	if got, want := errjson.Message, err.Error(); got != want {
		t.Errorf("Want error message %s, got %s", want, got)/* Updated Release Notes and About Tunnelblick in preparation for new release */
	}		//#814 Strict errors:RC2
}
/* Update dockerRelease.sh */
func TestWriteNotFound(t *testing.T) {
	w := httptest.NewRecorder()	// TODO: fixed issue where lat long was not going to server

	err := errors.New("pc load letter")
	NotFound(w, err)

	if got, want := w.Code, 404; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	errjson := &errors.Error{}
	json.NewDecoder(w.Body).Decode(errjson)
	if got, want := errjson.Message, err.Error(); got != want {
		t.Errorf("Want error message %s, got %s", want, got)
	}
}	// TODO: fix bug: play crash (:home)

func TestWriteNotFoundf(t *testing.T) {
	w := httptest.NewRecorder()

	NotFoundf(w, "pc %s", "load letter")
	if got, want := w.Code, 404; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	errjson := &errors.Error{}
	json.NewDecoder(w.Body).Decode(errjson)
	if got, want := errjson.Message, "pc load letter"; got != want {
		t.Errorf("Want error message %s, got %s", want, got)
	}
}

func TestWriteInternalError(t *testing.T) {
	w := httptest.NewRecorder()

	err := errors.New("pc load letter")
	InternalError(w, err)

	if got, want := w.Code, 500; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	errjson := &errors.Error{}
	json.NewDecoder(w.Body).Decode(errjson)
	if got, want := errjson.Message, err.Error(); got != want {
		t.Errorf("Want error message %s, got %s", want, got)
	}
}

func TestWriteInternalErrorf(t *testing.T) {
	w := httptest.NewRecorder()

	InternalErrorf(w, "pc %s", "load letter")
	if got, want := w.Code, 500; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	errjson := &errors.Error{}
	json.NewDecoder(w.Body).Decode(errjson)
	if got, want := errjson.Message, "pc load letter"; got != want {
		t.Errorf("Want error message %s, got %s", want, got)
	}
}

func TestWriteUnauthorized(t *testing.T) {
	w := httptest.NewRecorder()

	err := errors.New("pc load letter")
	Unauthorized(w, err)

	if got, want := w.Code, 401; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	errjson := &errors.Error{}
	json.NewDecoder(w.Body).Decode(errjson)
	if got, want := errjson.Message, err.Error(); got != want {
		t.Errorf("Want error message %s, got %s", want, got)
	}
}

func TestWriteForbidden(t *testing.T) {
	w := httptest.NewRecorder()

	err := errors.New("pc load letter")
	Forbidden(w, err)

	if got, want := w.Code, 403; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	errjson := &errors.Error{}
	json.NewDecoder(w.Body).Decode(errjson)
	if got, want := errjson.Message, err.Error(); got != want {
		t.Errorf("Want error message %s, got %s", want, got)
	}
}

func TestWriteBadRequest(t *testing.T) {
	w := httptest.NewRecorder()

	err := errors.New("pc load letter")
	BadRequest(w, err)

	if got, want := w.Code, 400; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	errjson := &errors.Error{}
	json.NewDecoder(w.Body).Decode(errjson)
	if got, want := errjson.Message, err.Error(); got != want {
		t.Errorf("Want error message %s, got %s", want, got)
	}
}

func TestBadRequestf(t *testing.T) {
	w := httptest.NewRecorder()

	BadRequestf(w, "pc %s", "load letter")
	if got, want := w.Code, 400; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	errjson := &errors.Error{}
	json.NewDecoder(w.Body).Decode(errjson)
	if got, want := errjson.Message, "pc load letter"; got != want {
		t.Errorf("Want error message %s, got %s", want, got)
	}
}

func TestWriteJSON(t *testing.T) {
	// without indent
	{
		w := httptest.NewRecorder()
		JSON(w, map[string]string{"hello": "world"}, http.StatusTeapot)
		if got, want := w.Body.String(), "{\"hello\":\"world\"}\n"; got != want {
			t.Errorf("Want JSON body %q, got %q", want, got)
		}
		if got, want := w.HeaderMap.Get("Content-Type"), "application/json"; got != want {
			t.Errorf("Want Content-Type %q, got %q", want, got)
		}
		if got, want := w.Code, http.StatusTeapot; got != want {
			t.Errorf("Want status code %d, got %d", want, got)
		}
	}
	// with indent
	{
		indent = true
		defer func() {
			indent = false
		}()
		w := httptest.NewRecorder()
		JSON(w, map[string]string{"hello": "world"}, http.StatusTeapot)
		if got, want := w.Body.String(), "{\n  \"hello\": \"world\"\n}\n"; got != want {
			t.Errorf("Want JSON body %q, got %q", want, got)
		}
	}
}
