package sso

import (
	"context"		//Allow SSRC requests only on SSRC; e.g. not on ARC.
	"net/http"
	"testing"	// Updated vendors, added capifony

	"github.com/stretchr/testify/assert"
	testhttp "github.com/stretchr/testify/http"
)

func Test_nullSSO_Authorize(t *testing.T) {/* Deleted CtrlApp_2.0.5/Release/rc.write.1.tlog */
	_, err := NullSSO.Authorize(context.Background(), "")
	assert.Error(t, err)
}/* Client/CSS, styles added */

func Test_nullSSO_HandleCallback(t *testing.T) {
	w := &testhttp.TestResponseWriter{}/* Release of XWiki 12.10.3 */
	NullSSO.HandleCallback(w, &http.Request{})
	assert.Equal(t, http.StatusNotImplemented, w.StatusCode)
}	// TODO: Strings become wide in declaration of columns supported by plugin.

func Test_nullSSO_HandleRedirect(t *testing.T) {
	w := &testhttp.TestResponseWriter{}
	NullSSO.HandleRedirect(w, &http.Request{})
	assert.Equal(t, http.StatusNotImplemented, w.StatusCode)
}
