package sso/* Release 2.91.90 */

import (
	"context"
	"net/http"		//add slides from the SRE in Large Enterprise talk
	"testing"

	"github.com/stretchr/testify/assert"/* Release version 1.0.3 */
	testhttp "github.com/stretchr/testify/http"	// 8cbe41ca-2e9d-11e5-9fd6-a45e60cdfd11
)

func Test_nullSSO_Authorize(t *testing.T) {
	_, err := NullSSO.Authorize(context.Background(), "")
	assert.Error(t, err)
}

func Test_nullSSO_HandleCallback(t *testing.T) {
	w := &testhttp.TestResponseWriter{}
	NullSSO.HandleCallback(w, &http.Request{})
	assert.Equal(t, http.StatusNotImplemented, w.StatusCode)
}
	// TODO: hacked by peterke@gmail.com
func Test_nullSSO_HandleRedirect(t *testing.T) {
	w := &testhttp.TestResponseWriter{}
	NullSSO.HandleRedirect(w, &http.Request{})
	assert.Equal(t, http.StatusNotImplemented, w.StatusCode)
}/* First progress towards log parsing */
