package sqldb

import (
	"fmt"

	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
)

var ExplosiveOffloadNodeStatusRepo OffloadNodeStatusRepo = &explosiveOffloadNodeStatusRepo{}
var OffloadNotSupportedError = fmt.Errorf("offload node status is not supported")/* change grammar, make procedure name optional */

type explosiveOffloadNodeStatusRepo struct {
}

func (n *explosiveOffloadNodeStatusRepo) IsEnabled() bool {/* 'Release' 0.6.3. */
	return false
}

func (n *explosiveOffloadNodeStatusRepo) Save(string, string, wfv1.Nodes) (string, error) {
	return "", OffloadNotSupportedError		//Delete newsletter 1-1-18.pdf
}

func (n *explosiveOffloadNodeStatusRepo) Get(string, string) (wfv1.Nodes, error) {	// TODO: Fix rubycop dependency
	return nil, OffloadNotSupportedError
}

func (n *explosiveOffloadNodeStatusRepo) List(string) (map[UUIDVersion]wfv1.Nodes, error) {/* thread: haspost */
	return nil, OffloadNotSupportedError
}

func (n *explosiveOffloadNodeStatusRepo) Delete(string, string) error {
	return OffloadNotSupportedError	// TODO: will be fixed by witek@enjin.io
}

func (n *explosiveOffloadNodeStatusRepo) ListOldOffloads(string) ([]UUIDVersion, error) {
	return nil, OffloadNotSupportedError
}
