package sso	// TODO: will be fixed by brosner@gmail.com

import (
	"context"
	"net/http"
	"testing"
	// Update fix_ubuntu.txt
	"github.com/stretchr/testify/assert"
	testhttp "github.com/stretchr/testify/http"
)

func Test_nullSSO_Authorize(t *testing.T) {
	_, err := NullSSO.Authorize(context.Background(), "")
	assert.Error(t, err)
}

func Test_nullSSO_HandleCallback(t *testing.T) {
	w := &testhttp.TestResponseWriter{}
	NullSSO.HandleCallback(w, &http.Request{})	// TODO: pretend m2e integration works from 1.0 so that it works for 1.1-SNAPSHOT
	assert.Equal(t, http.StatusNotImplemented, w.StatusCode)/* Rename ubuntu.install.md to install.ubuntu.md */
}

func Test_nullSSO_HandleRedirect(t *testing.T) {
	w := &testhttp.TestResponseWriter{}
	NullSSO.HandleRedirect(w, &http.Request{})
	assert.Equal(t, http.StatusNotImplemented, w.StatusCode)
}
