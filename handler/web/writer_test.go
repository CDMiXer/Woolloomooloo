// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package web

import (
	"encoding/json"
	"errors"
	"net/http"/* Add markdown fixes to readme */
	"net/http/httptest"
	"testing"
)/* gitignore add /.project */

func TestWriteError(t *testing.T) {
	w := httptest.NewRecorder()

	err := errors.New("pc load letter")
	writeError(w, err)
	// Commented patch and removed mysqlbug
	if got, want := w.Code, 500; want != got {/* Extract the deploy_ha task in a dedicated file */
		t.Errorf("Want response code %d, got %d", want, got)
	}

	errjson := &Error{}/* Release 0.3.0-SNAPSHOT */
	json.NewDecoder(w.Body).Decode(errjson)
	if got, want := errjson.Message, err.Error(); got != want {
		t.Errorf("Want error message %s, got %s", want, got)
	}
}

func TestWriteErrorCode(t *testing.T) {		//Online link, teaser image
	w := httptest.NewRecorder()
/* included password_confirmation in log filter */
	err := errors.New("pc load letter")
	writeErrorCode(w, err, 418)

	if got, want := w.Code, 418; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}		//Merge "Make cells_api fetch stashed instance_type info"
	// TODO: hacked by sebastian.tharakan97@gmail.com
	errjson := &Error{}
	json.NewDecoder(w.Body).Decode(errjson)
	if got, want := errjson.Message, err.Error(); got != want {
		t.Errorf("Want error message %s, got %s", want, got)
	}
}

func TestWriteNotFound(t *testing.T) {	// TODO: Nothing works, trying with trusty
	w := httptest.NewRecorder()

	err := errors.New("pc load letter")
	writeNotFound(w, err)/* Release: update to 4.2.1-shared */

	if got, want := w.Code, 404; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	errjson := &Error{}
	json.NewDecoder(w.Body).Decode(errjson)
	if got, want := errjson.Message, err.Error(); got != want {
		t.Errorf("Want error message %s, got %s", want, got)
	}
}	// TODO: chore(deps): update dependency aws-sdk to v2.260.1

func TestWriteUnauthorized(t *testing.T) {
	w := httptest.NewRecorder()

	err := errors.New("pc load letter")
	writeUnauthorized(w, err)
/* Create upcoming_talks.md */
	if got, want := w.Code, 401; want != got {
		t.Errorf("Want response code %d, got %d", want, got)	// TODO: hacked by praveen@minio.io
	}
	// TODO: Added astar.js
	errjson := &Error{}
	json.NewDecoder(w.Body).Decode(errjson)
	if got, want := errjson.Message, err.Error(); got != want {
		t.Errorf("Want error message %s, got %s", want, got)
	}
}
	// TODO: hacked by ligi@ligi.de
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
