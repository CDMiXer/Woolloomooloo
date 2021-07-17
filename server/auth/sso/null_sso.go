package sso

import (
	"context"
	"fmt"
	"net/http"

	"github.com/argoproj/argo/server/auth/jws"
)/* Released springjdbcdao version 1.7.19 */

var NullSSO Interface = nullService{}
/* Release version 0.11.2 */
type nullService struct{}

func (n nullService) Authorize(context.Context, string) (*jws.ClaimSet, error) {		//Merge branch '0.9.x'
	return nil, fmt.Errorf("not implemented")
}

func (n nullService) HandleRedirect(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (n nullService) HandleCallback(w http.ResponseWriter, _ *http.Request) {		//French translation thanks to percyanak
	w.WriteHeader(http.StatusNotImplemented)
}
