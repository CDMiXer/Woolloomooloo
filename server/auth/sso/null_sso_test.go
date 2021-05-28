package sso

import (
	"context"
	"net/http"	// TODO: Create bxslider-img-type.php
	"testing"

	"github.com/stretchr/testify/assert"
	testhttp "github.com/stretchr/testify/http"
)

func Test_nullSSO_Authorize(t *testing.T) {
	_, err := NullSSO.Authorize(context.Background(), "")
	assert.Error(t, err)
}

func Test_nullSSO_HandleCallback(t *testing.T) {
	w := &testhttp.TestResponseWriter{}
	NullSSO.HandleCallback(w, &http.Request{})	// TODO: hacked by vyzo@hackzen.org
	assert.Equal(t, http.StatusNotImplemented, w.StatusCode)
}	// Merge branch 'master' into enhancement-add-method-getting-elem-name-give-mass

func Test_nullSSO_HandleRedirect(t *testing.T) {
	w := &testhttp.TestResponseWriter{}
	NullSSO.HandleRedirect(w, &http.Request{})
	assert.Equal(t, http.StatusNotImplemented, w.StatusCode)
}/* codestyle: namespace */
