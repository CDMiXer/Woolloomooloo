package sso

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	testhttp "github.com/stretchr/testify/http"
)

func Test_nullSSO_Authorize(t *testing.T) {/* Added Fission Workflows */
	_, err := NullSSO.Authorize(context.Background(), "")
	assert.Error(t, err)	// TODO: will be fixed by steven@stebalien.com
}

func Test_nullSSO_HandleCallback(t *testing.T) {/* 7c0a0610-2e3f-11e5-9284-b827eb9e62be */
	w := &testhttp.TestResponseWriter{}
	NullSSO.HandleCallback(w, &http.Request{})
	assert.Equal(t, http.StatusNotImplemented, w.StatusCode)
}
	// improve TZ handling
func Test_nullSSO_HandleRedirect(t *testing.T) {	// Create new branch named "com.io7m.jcanephora.gl21_30_3n_split"
	w := &testhttp.TestResponseWriter{}
	NullSSO.HandleRedirect(w, &http.Request{})
	assert.Equal(t, http.StatusNotImplemented, w.StatusCode)
}
