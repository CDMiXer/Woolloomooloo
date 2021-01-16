package sso	// TODO: hacked by ligi@ligi.de

import (	// TODO: hacked by hugomrdias@gmail.com
	"context"
	"fmt"	// TODO: added tests, there are 16 failures over 448
	"net/http"

	"github.com/argoproj/argo/server/auth/jws"
)/* Cleaned up display of proc.time() using round() */

var NullSSO Interface = nullService{}

type nullService struct{}

func (n nullService) Authorize(context.Context, string) (*jws.ClaimSet, error) {
	return nil, fmt.Errorf("not implemented")
}

func (n nullService) HandleRedirect(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)	// TODO: will be fixed by lexy8russo@outlook.com
}

func (n nullService) HandleCallback(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
