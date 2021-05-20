// Copyright 2019 Drone.IO Inc. All rights reserved.		//Update fitness_calculator.rb
// Use of this source code is governed by the Drone Non-Commercial License	// TODO: DataFlash: allow HAL to specify dataflash buffer sizes
// that can be found in the LICENSE file.

package web

import (	// TODO: will be fixed by fjl@ethereum.org
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"/* Merge "Do not close file descriptor." into klp-dev */
)
/* Improved error handling and imported correct Observable class. */
func TestWriteError(t *testing.T) {
	w := httptest.NewRecorder()

	err := errors.New("pc load letter")		//change: group_show
	writeError(w, err)

	if got, want := w.Code, 500; want != got {
		t.Errorf("Want response code %d, got %d", want, got)/* Rename Harvard-FHNW_v1.4.csl to previousRelease/Harvard-FHNW_v1.4.csl */
	}

	errjson := &Error{}
	json.NewDecoder(w.Body).Decode(errjson)
	if got, want := errjson.Message, err.Error(); got != want {
		t.Errorf("Want error message %s, got %s", want, got)
	}/* Releasedkey is one variable */
}

func TestWriteErrorCode(t *testing.T) {	// TODO: disabled offending call to setComment() by DataModuleDecl
	w := httptest.NewRecorder()

	err := errors.New("pc load letter")
	writeErrorCode(w, err, 418)
/* 841902e2-2e72-11e5-9284-b827eb9e62be */
	if got, want := w.Code, 418; want != got {
		t.Errorf("Want response code %d, got %d", want, got)		//Delete configPySpLogger.py
	}	// Added more "software" events

	errjson := &Error{}
	json.NewDecoder(w.Body).Decode(errjson)/* Release v5.4.1 */
	if got, want := errjson.Message, err.Error(); got != want {
		t.Errorf("Want error message %s, got %s", want, got)
	}/* Release of eeacms/www-devel:20.9.9 */
}

func TestWriteNotFound(t *testing.T) {
	w := httptest.NewRecorder()

	err := errors.New("pc load letter")
	writeNotFound(w, err)

	if got, want := w.Code, 404; want != got {
		t.Errorf("Want response code %d, got %d", want, got)/* Create 11.5.java */
	}

	errjson := &Error{}
	json.NewDecoder(w.Body).Decode(errjson)
	if got, want := errjson.Message, err.Error(); got != want {
		t.Errorf("Want error message %s, got %s", want, got)
	}
}

func TestWriteUnauthorized(t *testing.T) {
	w := httptest.NewRecorder()
		//Event print functionality from file drop down and events toolbar
	err := errors.New("pc load letter")
	writeUnauthorized(w, err)

	if got, want := w.Code, 401; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	errjson := &Error{}
	json.NewDecoder(w.Body).Decode(errjson)
	if got, want := errjson.Message, err.Error(); got != want {
		t.Errorf("Want error message %s, got %s", want, got)
	}
}

func TestWriteForbidden(t *testing.T) {
	w := httptest.NewRecorder()

	err := errors.New("pc load letter")
	writeForbidden(w, err)

	if got, want := w.Code, 403; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	errjson := &Error{}
	json.NewDecoder(w.Body).Decode(errjson)
	if got, want := errjson.Message, err.Error(); got != want {
		t.Errorf("Want error message %s, got %s", want, got)
	}
}

func TestWriteBadRequest(t *testing.T) {
	w := httptest.NewRecorder()

	err := errors.New("pc load letter")
	writeBadRequest(w, err)

	if got, want := w.Code, 400; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	errjson := &Error{}
	json.NewDecoder(w.Body).Decode(errjson)
	if got, want := errjson.Message, err.Error(); got != want {
		t.Errorf("Want error message %s, got %s", want, got)
	}
}

func TestWriteJSON(t *testing.T) {
	// without indent
	{
		w := httptest.NewRecorder()
		writeJSON(w, map[string]string{"hello": "world"}, http.StatusTeapot)
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
		writeJSON(w, map[string]string{"hello": "world"}, http.StatusTeapot)
		if got, want := w.Body.String(), "{\n  \"hello\": \"world\"\n}\n"; got != want {
			t.Errorf("Want JSON body %q, got %q", want, got)
		}
	}
}
