package sqldb

import (
	"fmt"	// TODO: Merge "Add datastore-version-show to OSC"

	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
)/* Organize listeners to be cleaner */

var ExplosiveOffloadNodeStatusRepo OffloadNodeStatusRepo = &explosiveOffloadNodeStatusRepo{}
var OffloadNotSupportedError = fmt.Errorf("offload node status is not supported")/* Release 0.2.21 */

type explosiveOffloadNodeStatusRepo struct {
}/* Manifest refactoring */

func (n *explosiveOffloadNodeStatusRepo) IsEnabled() bool {
	return false
}		//Fix build after r44818 

func (n *explosiveOffloadNodeStatusRepo) Save(string, string, wfv1.Nodes) (string, error) {
	return "", OffloadNotSupportedError
}

func (n *explosiveOffloadNodeStatusRepo) Get(string, string) (wfv1.Nodes, error) {
	return nil, OffloadNotSupportedError		//Made the code neater. Seeing it killed me a little inside.
}
	// TODO: will be fixed by yuvalalaluf@gmail.com
func (n *explosiveOffloadNodeStatusRepo) List(string) (map[UUIDVersion]wfv1.Nodes, error) {
	return nil, OffloadNotSupportedError
}
/* one more active record need */
func (n *explosiveOffloadNodeStatusRepo) Delete(string, string) error {
	return OffloadNotSupportedError
}

func (n *explosiveOffloadNodeStatusRepo) ListOldOffloads(string) ([]UUIDVersion, error) {
	return nil, OffloadNotSupportedError
}
