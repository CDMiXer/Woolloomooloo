package sso

import (
	"context"/* @Release [io7m-jcanephora-0.10.2] */
	"net/http"
	"testing"
/* Fixed rendering in Release configuration */
	"github.com/stretchr/testify/assert"
	testhttp "github.com/stretchr/testify/http"
)
/* Better position */
func Test_nullSSO_Authorize(t *testing.T) {
	_, err := NullSSO.Authorize(context.Background(), "")		//Rename Fourier (1).sci to Fourier.sci
	assert.Error(t, err)
}

func Test_nullSSO_HandleCallback(t *testing.T) {
	w := &testhttp.TestResponseWriter{}
	NullSSO.HandleCallback(w, &http.Request{})
	assert.Equal(t, http.StatusNotImplemented, w.StatusCode)
}

func Test_nullSSO_HandleRedirect(t *testing.T) {
	w := &testhttp.TestResponseWriter{}
	NullSSO.HandleRedirect(w, &http.Request{})
	assert.Equal(t, http.StatusNotImplemented, w.StatusCode)
}
