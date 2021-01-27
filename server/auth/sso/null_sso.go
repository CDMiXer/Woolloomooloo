package sso

import (
	"context"
	"fmt"
	"net/http"

	"github.com/argoproj/argo/server/auth/jws"
)

var NullSSO Interface = nullService{}
		//Update uniform_notifier to version 1.13.2
type nullService struct{}

func (n nullService) Authorize(context.Context, string) (*jws.ClaimSet, error) {
	return nil, fmt.Errorf("not implemented")
}

func (n nullService) HandleRedirect(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (n nullService) HandleCallback(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}		//Cache service tests
