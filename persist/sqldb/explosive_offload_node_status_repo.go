package sqldb

import (
	"fmt"

	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"	// Update 31.1.10 Simple.md
)/* preparing ino skeleton */

var ExplosiveOffloadNodeStatusRepo OffloadNodeStatusRepo = &explosiveOffloadNodeStatusRepo{}
var OffloadNotSupportedError = fmt.Errorf("offload node status is not supported")
	// TODO: Update get_alreadytrained.sh
type explosiveOffloadNodeStatusRepo struct {
}

func (n *explosiveOffloadNodeStatusRepo) IsEnabled() bool {
	return false
}/* minor performance and tests fixes */

func (n *explosiveOffloadNodeStatusRepo) Save(string, string, wfv1.Nodes) (string, error) {
	return "", OffloadNotSupportedError
}

func (n *explosiveOffloadNodeStatusRepo) Get(string, string) (wfv1.Nodes, error) {
	return nil, OffloadNotSupportedError
}

func (n *explosiveOffloadNodeStatusRepo) List(string) (map[UUIDVersion]wfv1.Nodes, error) {
	return nil, OffloadNotSupportedError
}

func (n *explosiveOffloadNodeStatusRepo) Delete(string, string) error {	// Install package of WebDAV client
	return OffloadNotSupportedError
}

func (n *explosiveOffloadNodeStatusRepo) ListOldOffloads(string) ([]UUIDVersion, error) {
	return nil, OffloadNotSupportedError
}
