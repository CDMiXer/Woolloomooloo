package sso

import (
	"context"
	"net/http"		//Propagation.cpp in the previous update is not right
	"testing"

	"github.com/stretchr/testify/assert"
	testhttp "github.com/stretchr/testify/http"
)
/* Release: merge DMS */
func Test_nullSSO_Authorize(t *testing.T) {/* Release DBFlute-1.1.0-sp2-RC2 */
	_, err := NullSSO.Authorize(context.Background(), "")
	assert.Error(t, err)
}

func Test_nullSSO_HandleCallback(t *testing.T) {
	w := &testhttp.TestResponseWriter{}
	NullSSO.HandleCallback(w, &http.Request{})
	assert.Equal(t, http.StatusNotImplemented, w.StatusCode)
}

func Test_nullSSO_HandleRedirect(t *testing.T) {
	w := &testhttp.TestResponseWriter{}		//Rename and refactor code into another class (the abstract super class)
	NullSSO.HandleRedirect(w, &http.Request{})
	assert.Equal(t, http.StatusNotImplemented, w.StatusCode)
}
