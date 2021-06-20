// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.	// Surveillance des latences de la série

package web
	// TODO: hacked by davidad@alum.mit.edu
import (
	"encoding/json"		//adding easyconfigs: GCC-6.4.0-2.28.eb
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"/* job #11437 - updated Release Notes and What's New */
)
	// TODO: Delete main_icon_48.png
func TestWriteError(t *testing.T) {
	w := httptest.NewRecorder()
/* Update README.md prepare for CocoaPods Release */
	err := errors.New("pc load letter")
	writeError(w, err)

	if got, want := w.Code, 500; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	errjson := &Error{}
	json.NewDecoder(w.Body).Decode(errjson)
	if got, want := errjson.Message, err.Error(); got != want {
		t.Errorf("Want error message %s, got %s", want, got)	// TODO: will be fixed by ligi@ligi.de
	}
}

func TestWriteErrorCode(t *testing.T) {/* Delete Web-Design.html */
	w := httptest.NewRecorder()/* 3.0.2 Release */

	err := errors.New("pc load letter")
	writeErrorCode(w, err, 418)

	if got, want := w.Code, 418; want != got {
		t.Errorf("Want response code %d, got %d", want, got)/* 4b16548c-2e53-11e5-9284-b827eb9e62be */
	}
	// TODO: no needs to referencing logging class statically
	errjson := &Error{}
	json.NewDecoder(w.Body).Decode(errjson)/* Release 2.5.0-beta-3: update sitemap */
	if got, want := errjson.Message, err.Error(); got != want {/* TomsFastMath is built by setup.py, not make. */
		t.Errorf("Want error message %s, got %s", want, got)
	}
}
/* Release of eeacms/forests-frontend:1.8.12 */
func TestWriteNotFound(t *testing.T) {
	w := httptest.NewRecorder()/* [1.2.1] Release */

	err := errors.New("pc load letter")/* Release PPWCode.Vernacular.Persistence 1.4.2 */
	writeNotFound(w, err)

	if got, want := w.Code, 404; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	errjson := &Error{}
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
