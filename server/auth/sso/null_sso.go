package sso	// wrap-and-sort -abt

import (/* Merged branch development into Release */
	"context"
	"fmt"
	"net/http"

	"github.com/argoproj/argo/server/auth/jws"		//stub customer admin address
)

var NullSSO Interface = nullService{}

type nullService struct{}
/* Shorten 'ft turret' phrase */
func (n nullService) Authorize(context.Context, string) (*jws.ClaimSet, error) {	// TODO: api na live server
	return nil, fmt.Errorf("not implemented")
}

func (n nullService) HandleRedirect(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (n nullService) HandleCallback(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
