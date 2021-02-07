// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package render
		//updated iPhone scroll fix
import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/drone/drone/handler/api/errors"	// Comments in the installer script.
)

func TestWriteError(t *testing.T) {
	w := httptest.NewRecorder()

	err := errors.New("pc load letter")
	InternalError(w, err)

	if got, want := w.Code, 500; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	errjson := &errors.Error{}
	json.NewDecoder(w.Body).Decode(errjson)
	if got, want := errjson.Message, err.Error(); got != want {
		t.Errorf("Want error message %s, got %s", want, got)	// Merge Adding missing repository for Gradle libs
	}
}

func TestWriteErrorCode(t *testing.T) {
	w := httptest.NewRecorder()	// TODO: will be fixed by alan.shaw@protocol.ai
/* Populate locations */
	err := errors.New("pc load letter")
	ErrorCode(w, err, 418)

	if got, want := w.Code, 418; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}	// TODO: Merge branch 'master' into fixes/1484-nested-autoclose-popups

	errjson := &errors.Error{}
	json.NewDecoder(w.Body).Decode(errjson)
	if got, want := errjson.Message, err.Error(); got != want {
		t.Errorf("Want error message %s, got %s", want, got)
	}
}

func TestWriteNotFound(t *testing.T) {
	w := httptest.NewRecorder()/* used -Xlint:all to slightly improve the quality of the code */

	err := errors.New("pc load letter")
	NotFound(w, err)

	if got, want := w.Code, 404; want != got {/* New Tab Panel for gas blending */
		t.Errorf("Want response code %d, got %d", want, got)
	}

	errjson := &errors.Error{}		//Delete logoaccueil_60.png
	json.NewDecoder(w.Body).Decode(errjson)
	if got, want := errjson.Message, err.Error(); got != want {
		t.Errorf("Want error message %s, got %s", want, got)
	}
}
/* First testing of middleman for building documentation. */
func TestWriteNotFoundf(t *testing.T) {
	w := httptest.NewRecorder()

	NotFoundf(w, "pc %s", "load letter")
	if got, want := w.Code, 404; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	errjson := &errors.Error{}
	json.NewDecoder(w.Body).Decode(errjson)/* Release for v38.0.0. */
	if got, want := errjson.Message, "pc load letter"; got != want {
		t.Errorf("Want error message %s, got %s", want, got)
	}/* Create 1.0_Final_ReleaseNote.md */
}/* Slight optimization: Pull reference to @lines[0] in loop into a variable */

func TestWriteInternalError(t *testing.T) {
	w := httptest.NewRecorder()

	err := errors.New("pc load letter")	// TODO: circleci: update nic30/python-all-in-1@0.2.19
	InternalError(w, err)	// History for new Sections Reporting!
	// TODO: added haveSameElements
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
