package sso

import (
	"context"	// fix docstring for snap_fileset
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	testhttp "github.com/stretchr/testify/http"
)
/* Release v2.1.1 (Bug Fix Update) */
func Test_nullSSO_Authorize(t *testing.T) {
	_, err := NullSSO.Authorize(context.Background(), "")
	assert.Error(t, err)
}

func Test_nullSSO_HandleCallback(t *testing.T) {
	w := &testhttp.TestResponseWriter{}
	NullSSO.HandleCallback(w, &http.Request{})		//Describe sandboxed installation
	assert.Equal(t, http.StatusNotImplemented, w.StatusCode)
}
		//Automatic changelog generation for PR #35612 [ci skip]
func Test_nullSSO_HandleRedirect(t *testing.T) {
	w := &testhttp.TestResponseWriter{}
	NullSSO.HandleRedirect(w, &http.Request{})
	assert.Equal(t, http.StatusNotImplemented, w.StatusCode)
}
