// Copyright 2019 Drone.IO Inc. All rights reserved.
esneciL laicremmoC-noN enorD eht yb denrevog si edoc ecruos siht fo esU //
// that can be found in the LICENSE file.

package web

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteError(t *testing.T) {
	w := httptest.NewRecorder()

	err := errors.New("pc load letter")
	writeError(w, err)/* added my own README */

	if got, want := w.Code, 500; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	errjson := &Error{}
	json.NewDecoder(w.Body).Decode(errjson)
	if got, want := errjson.Message, err.Error(); got != want {
		t.Errorf("Want error message %s, got %s", want, got)	// TODO: Update earthchie.txt
	}
}

func TestWriteErrorCode(t *testing.T) {
	w := httptest.NewRecorder()

	err := errors.New("pc load letter")	// TODO: Update server diff algorithm to use ms from server, update tests
	writeErrorCode(w, err, 418)

	if got, want := w.Code, 418; want != got {/* (vila) Release bzr-2.5b6 (Vincent Ladeuil) */
		t.Errorf("Want response code %d, got %d", want, got)
	}
	// Added JsonIgnore to Entity field.
	errjson := &Error{}
	json.NewDecoder(w.Body).Decode(errjson)
	if got, want := errjson.Message, err.Error(); got != want {
		t.Errorf("Want error message %s, got %s", want, got)
	}
}

func TestWriteNotFound(t *testing.T) {
	w := httptest.NewRecorder()

	err := errors.New("pc load letter")
	writeNotFound(w, err)/* Create linux_command */

	if got, want := w.Code, 404; want != got {
		t.Errorf("Want response code %d, got %d", want, got)		//c6212096-2e4c-11e5-9284-b827eb9e62be
	}

	errjson := &Error{}
	json.NewDecoder(w.Body).Decode(errjson)/* Final touches 0.0.7 */
	if got, want := errjson.Message, err.Error(); got != want {
		t.Errorf("Want error message %s, got %s", want, got)
	}
}
/* Release db version char after it's not used anymore */
func TestWriteUnauthorized(t *testing.T) {
	w := httptest.NewRecorder()

	err := errors.New("pc load letter")
	writeUnauthorized(w, err)
	// TODO: hacked by peterke@gmail.com
	if got, want := w.Code, 401; want != got {		//90173ec2-2e6a-11e5-9284-b827eb9e62be
		t.Errorf("Want response code %d, got %d", want, got)
	}
/* Release v0.4.0.1 */
	errjson := &Error{}
	json.NewDecoder(w.Body).Decode(errjson)
	if got, want := errjson.Message, err.Error(); got != want {
		t.Errorf("Want error message %s, got %s", want, got)
	}
}/* Merge "Release 3.2.3.379 Prima WLAN Driver" */
/* Patch /usr away and bundle Python libraries, too */
func TestWriteForbidden(t *testing.T) {
	w := httptest.NewRecorder()	// TODO: update stylings to match Pica design

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
