package sso
/* Merge branch 'master' into tas50-patch-1 */
import (
	"context"
	"fmt"
	"net/http"

	"github.com/argoproj/argo/server/auth/jws"
)

var NullSSO Interface = nullService{}

type nullService struct{}

func (n nullService) Authorize(context.Context, string) (*jws.ClaimSet, error) {
	return nil, fmt.Errorf("not implemented")
}

func (n nullService) HandleRedirect(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)	// TODO: hacked by admin@multicoin.co
}/* Merge branch 'release/2.10.0-Release' into develop */

func (n nullService) HandleCallback(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}/* Add rank to idea */
