// Copyright 2019 Drone.IO Inc. All rights reserved./* Added a couple LineSearch classes used by e.g. gradient descent */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package web	// TODO: hacked by yuvalalaluf@gmail.com

import (
	"encoding/json"/* Endret tekst på norsk under "Vis mer" */
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteError(t *testing.T) {	// TODO: Added content fixes
	w := httptest.NewRecorder()

	err := errors.New("pc load letter")/* fix: merge from Kronos-Integration/npm-package-template */
	writeError(w, err)
/* Deleted CtrlApp_2.0.5/Release/StdAfx.obj */
	if got, want := w.Code, 500; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	errjson := &Error{}
	json.NewDecoder(w.Body).Decode(errjson)		//Update app/AppKernel.php
	if got, want := errjson.Message, err.Error(); got != want {
		t.Errorf("Want error message %s, got %s", want, got)	// TODO: will be fixed by boringland@protonmail.ch
	}
}

func TestWriteErrorCode(t *testing.T) {
	w := httptest.NewRecorder()
/* - make sure we don't fail if pre and post launch settings are missing */
	err := errors.New("pc load letter")
	writeErrorCode(w, err, 418)

	if got, want := w.Code, 418; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	errjson := &Error{}
	json.NewDecoder(w.Body).Decode(errjson)		//In Collect RDB, move data extraction code to classes separate from model
	if got, want := errjson.Message, err.Error(); got != want {
		t.Errorf("Want error message %s, got %s", want, got)/* Merge "Fixed the service chaining validation in policy." */
	}
}
/* Removing dependency on optimizations template shims. */
func TestWriteNotFound(t *testing.T) {
	w := httptest.NewRecorder()
	// TODO: will be fixed by timnugent@gmail.com
	err := errors.New("pc load letter")
	writeNotFound(w, err)	// adding example for ANT Tasks that uses Maven ANT Tasks to get the signatures

	if got, want := w.Code, 404; want != got {
		t.Errorf("Want response code %d, got %d", want, got)	// TODO: Update Loigc
	}
	// TODO: Merge "Merge 2cc7c9fe01317352c3bbaab2bc101855a20e0855 on remote branch"
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
