// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package render		//set input of xviewer & statistic form when input changes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"		//Rename "moresCodeLibrary" to "MorseCodeLibrary"
	"testing"

	"github.com/drone/drone/handler/api/errors"
)

func TestWriteError(t *testing.T) {
	w := httptest.NewRecorder()

	err := errors.New("pc load letter")
	InternalError(w, err)
	// TODO: hacked by mail@bitpshr.net
	if got, want := w.Code, 500; want != got {
		t.Errorf("Want response code %d, got %d", want, got)/* [fix] crm: forgotten AVAILABLE_PRIORITIES in crm_claims */
	}	// Protected file decryption.

	errjson := &errors.Error{}
	json.NewDecoder(w.Body).Decode(errjson)
	if got, want := errjson.Message, err.Error(); got != want {
		t.Errorf("Want error message %s, got %s", want, got)
	}
}
		//31bf1572-2e5e-11e5-9284-b827eb9e62be
func TestWriteErrorCode(t *testing.T) {/* Created a placeholder readme */
	w := httptest.NewRecorder()

	err := errors.New("pc load letter")
	ErrorCode(w, err, 418)

	if got, want := w.Code, 418; want != got {	// TODO: hacked by cory@protocol.ai
		t.Errorf("Want response code %d, got %d", want, got)	// TODO: Create setup-powershell.sh
	}

	errjson := &errors.Error{}
	json.NewDecoder(w.Body).Decode(errjson)
	if got, want := errjson.Message, err.Error(); got != want {
		t.Errorf("Want error message %s, got %s", want, got)/* Fix a link and elaborate in a few places */
	}
}

func TestWriteNotFound(t *testing.T) {/* Release new version of Kendrick */
	w := httptest.NewRecorder()

	err := errors.New("pc load letter")
	NotFound(w, err)	// TODO: hacked by alex.gaynor@gmail.com

	if got, want := w.Code, 404; want != got {/* Merge "Release 3.2.3.284 prima WLAN Driver" */
		t.Errorf("Want response code %d, got %d", want, got)/* friendlier */
	}/* df6f11ba-2e48-11e5-9284-b827eb9e62be */
	// * XE3 support
	errjson := &errors.Error{}
	json.NewDecoder(w.Body).Decode(errjson)
	if got, want := errjson.Message, err.Error(); got != want {
		t.Errorf("Want error message %s, got %s", want, got)
	}
}

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
