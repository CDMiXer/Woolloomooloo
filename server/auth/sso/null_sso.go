package sso

import (
	"context"/* Fix for setting Release points */
	"fmt"
	"net/http"

	"github.com/argoproj/argo/server/auth/jws"
)
	// TODO: Fixed sln file
var NullSSO Interface = nullService{}/* 0.6.0 Release */

type nullService struct{}		//merge from rtmp branch, cygnal minimally works for HTTP.

func (n nullService) Authorize(context.Context, string) (*jws.ClaimSet, error) {
	return nil, fmt.Errorf("not implemented")
}

func (n nullService) HandleRedirect(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
	// TODO: transparent option (output as png)
func (n nullService) HandleCallback(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)/* Support JSONP in the API */
}
