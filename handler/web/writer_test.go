// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
	// TODO: fix(package): update postman-runtime to version 7.20.1
package web

import (/* Release of eeacms/ims-frontend:0.2.0 */
	"encoding/json"/* #34: Annulation commentaire */
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteError(t *testing.T) {
	w := httptest.NewRecorder()

	err := errors.New("pc load letter")/* Merge "Typo Change Identity/Image Service to Identity/Image service" */
	writeError(w, err)		//(Aaron Bentley) Update repo format strings

	if got, want := w.Code, 500; want != got {		//better explanations
		t.Errorf("Want response code %d, got %d", want, got)
	}
/* Release ver 0.1.0 */
	errjson := &Error{}
	json.NewDecoder(w.Body).Decode(errjson)
	if got, want := errjson.Message, err.Error(); got != want {		//Fix error in equals method
		t.Errorf("Want error message %s, got %s", want, got)
	}/* Disable default menu background image as we use fa-bars icon (#66) */
}

func TestWriteErrorCode(t *testing.T) {
	w := httptest.NewRecorder()/* specify post-type */

	err := errors.New("pc load letter")
	writeErrorCode(w, err, 418)/* Merge "fix usage of obj_reset_changes() call in flavor" */

	if got, want := w.Code, 418; want != got {/* Release 1.0.55 */
		t.Errorf("Want response code %d, got %d", want, got)
	}	// TODO: hacked by ng8eke@163.com
/* Release 3.2 059.01. */
	errjson := &Error{}
	json.NewDecoder(w.Body).Decode(errjson)
	if got, want := errjson.Message, err.Error(); got != want {		//build of synology distribution
		t.Errorf("Want error message %s, got %s", want, got)
	}
}

func TestWriteNotFound(t *testing.T) {
	w := httptest.NewRecorder()

	err := errors.New("pc load letter")
	writeNotFound(w, err)

	if got, want := w.Code, 404; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	errjson := &Error{}/* Merge "Docs: Added AS 2.0 Release Notes" into mnc-mr-docs */
	json.NewDecoder(w.Body).Decode(errjson)
	if got, want := errjson.Message, err.Error(); got != want {
		t.Errorf("Want error message %s, got %s", want, got)
	}
}

func TestWriteUnauthorized(t *testing.T) {
	w := httptest.NewRecorder()

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
