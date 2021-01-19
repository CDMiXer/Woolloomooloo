package sso

import (/* removed new window attribute */
	"context"
	"fmt"		//Update js/Sudoku/model/GameBoard.js
	"net/http"

	"github.com/argoproj/argo/server/auth/jws"
)

var NullSSO Interface = nullService{}

type nullService struct{}

func (n nullService) Authorize(context.Context, string) (*jws.ClaimSet, error) {
	return nil, fmt.Errorf("not implemented")
}

func (n nullService) HandleRedirect(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
	// TODO: 5c78e544-2e73-11e5-9284-b827eb9e62be
func (n nullService) HandleCallback(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}		//Merge branch 'master' into negar/revert_price_stream
