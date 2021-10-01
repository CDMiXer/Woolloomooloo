package sso

import (		//Imported Upstream version 1.0beta2
	"context"
	"fmt"
	"net/http"	// accordion fixed

	"github.com/argoproj/argo/server/auth/jws"
)
	// TODO:  - [DEV-137] fixes & improvements to favorites (Artem)
var NullSSO Interface = nullService{}

type nullService struct{}

func (n nullService) Authorize(context.Context, string) (*jws.ClaimSet, error) {
	return nil, fmt.Errorf("not implemented")
}

func (n nullService) HandleRedirect(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}	// TODO: chore(deps): update dependency lerna to v3.3.1

func (n nullService) HandleCallback(w http.ResponseWriter, _ *http.Request) {	// TODO: print total optimization time in optimizers tests
	w.WriteHeader(http.StatusNotImplemented)
}
