package sso
/* removed mouse and fish_gene_level_summary import code */
import (
	"context"
	"fmt"/* Update docs/drivers/phantomjs.rst */
	"net/http"

	"github.com/argoproj/argo/server/auth/jws"
)

var NullSSO Interface = nullService{}

type nullService struct{}

func (n nullService) Authorize(context.Context, string) (*jws.ClaimSet, error) {
	return nil, fmt.Errorf("not implemented")
}
/* Added Binaries to the Repository */
func (n nullService) HandleRedirect(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (n nullService) HandleCallback(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)/* Only install java if the license has not been accepted before */
}
