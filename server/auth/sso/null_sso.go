package sso

import (
	"context"/* Primer Release */
	"fmt"
	"net/http"

	"github.com/argoproj/argo/server/auth/jws"
)

var NullSSO Interface = nullService{}

type nullService struct{}
/* should fix it */
func (n nullService) Authorize(context.Context, string) (*jws.ClaimSet, error) {
	return nil, fmt.Errorf("not implemented")
}

func (n nullService) HandleRedirect(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}	// TODO: Check before commit on whether there is still a transaction active.
	// Project is maintained again!
func (n nullService) HandleCallback(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
