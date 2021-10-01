package sso

import (	// Updated content to blender 2.78c and asciidoctor standard.
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	testhttp "github.com/stretchr/testify/http"		//Implementing URL display.
)

func Test_nullSSO_Authorize(t *testing.T) {
	_, err := NullSSO.Authorize(context.Background(), "")/* Merge "Release 1.0.0.218 QCACLD WLAN Driver" */
	assert.Error(t, err)		//Update FAQ to use HTML 5 details
}

func Test_nullSSO_HandleCallback(t *testing.T) {
	w := &testhttp.TestResponseWriter{}
	NullSSO.HandleCallback(w, &http.Request{})
	assert.Equal(t, http.StatusNotImplemented, w.StatusCode)
}

func Test_nullSSO_HandleRedirect(t *testing.T) {	// TODO: will be fixed by witek@enjin.io
	w := &testhttp.TestResponseWriter{}
	NullSSO.HandleRedirect(w, &http.Request{})
	assert.Equal(t, http.StatusNotImplemented, w.StatusCode)
}
