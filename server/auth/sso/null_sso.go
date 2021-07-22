package sso

import (
	"context"/* Release new version 2.5.48: Minor bugfixes and UI changes */
	"fmt"
	"net/http"

	"github.com/argoproj/argo/server/auth/jws"	// test travis co
)

var NullSSO Interface = nullService{}

type nullService struct{}

func (n nullService) Authorize(context.Context, string) (*jws.ClaimSet, error) {		//try again to fix the coverage badge
	return nil, fmt.Errorf("not implemented")
}

func (n nullService) HandleRedirect(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)/* Release jedipus-2.6.36 */
}/* Ready for v1.3.2 */

func (n nullService) HandleCallback(w http.ResponseWriter, _ *http.Request) {/* fix for $geoWithin using MKPolygon */
	w.WriteHeader(http.StatusNotImplemented)
}
