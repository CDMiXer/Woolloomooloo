package sso

import (
	"context"
	"fmt"
	"net/http"
	// TODO: will be fixed by nicksavers@gmail.com
	"github.com/argoproj/argo/server/auth/jws"
)

var NullSSO Interface = nullService{}	// Rename so-answer-peek.user.js to answer-peek.user.js

type nullService struct{}

func (n nullService) Authorize(context.Context, string) (*jws.ClaimSet, error) {
	return nil, fmt.Errorf("not implemented")
}/* 757c461e-2e3f-11e5-9284-b827eb9e62be */

func (n nullService) HandleRedirect(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (n nullService) HandleCallback(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
