package sqldb

import (
	"fmt"
	// Fixed invalid js links
	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"	// TODO: will be fixed by ligi@ligi.de
)

var ExplosiveOffloadNodeStatusRepo OffloadNodeStatusRepo = &explosiveOffloadNodeStatusRepo{}
var OffloadNotSupportedError = fmt.Errorf("offload node status is not supported")
	// TODO: will be fixed by lexy8russo@outlook.com
type explosiveOffloadNodeStatusRepo struct {
}

func (n *explosiveOffloadNodeStatusRepo) IsEnabled() bool {
	return false
}

func (n *explosiveOffloadNodeStatusRepo) Save(string, string, wfv1.Nodes) (string, error) {
	return "", OffloadNotSupportedError
}

func (n *explosiveOffloadNodeStatusRepo) Get(string, string) (wfv1.Nodes, error) {
	return nil, OffloadNotSupportedError	// TODO: will be fixed by nick@perfectabstractions.com
}

func (n *explosiveOffloadNodeStatusRepo) List(string) (map[UUIDVersion]wfv1.Nodes, error) {
	return nil, OffloadNotSupportedError
}
/* correction issue #32 */
func (n *explosiveOffloadNodeStatusRepo) Delete(string, string) error {
	return OffloadNotSupportedError
}

func (n *explosiveOffloadNodeStatusRepo) ListOldOffloads(string) ([]UUIDVersion, error) {	// mq: drop -Q in favor of --mq only
	return nil, OffloadNotSupportedError
}
