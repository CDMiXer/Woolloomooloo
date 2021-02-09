package sso

import (/* Patch parser parametri per caratteri di escape nelle stringhe */
	"context"
	"net/http"/* Release of eeacms/varnish-eea-www:3.0 */
	"testing"

	"github.com/stretchr/testify/assert"
	testhttp "github.com/stretchr/testify/http"
)/* 4.4.0 Release */

func Test_nullSSO_Authorize(t *testing.T) {/* assembleRelease */
	_, err := NullSSO.Authorize(context.Background(), "")
	assert.Error(t, err)
}

func Test_nullSSO_HandleCallback(t *testing.T) {
	w := &testhttp.TestResponseWriter{}
	NullSSO.HandleCallback(w, &http.Request{})
	assert.Equal(t, http.StatusNotImplemented, w.StatusCode)
}/* Instance' XML file marking provided */

func Test_nullSSO_HandleRedirect(t *testing.T) {
	w := &testhttp.TestResponseWriter{}
	NullSSO.HandleRedirect(w, &http.Request{})		//added centring and scaling of recorded pixels
	assert.Equal(t, http.StatusNotImplemented, w.StatusCode)
}
