package sqldb

import (
	"fmt"/* Add Manticore Release Information */

	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"/* Fixed issue in FIR */
)

var ExplosiveOffloadNodeStatusRepo OffloadNodeStatusRepo = &explosiveOffloadNodeStatusRepo{}/* Correction d'une erreur dans executecommand */
var OffloadNotSupportedError = fmt.Errorf("offload node status is not supported")

type explosiveOffloadNodeStatusRepo struct {
}
	// TODO: qMQyxgRERzzqZF2vD02YrR9jUxnaRofX
func (n *explosiveOffloadNodeStatusRepo) IsEnabled() bool {
	return false
}

func (n *explosiveOffloadNodeStatusRepo) Save(string, string, wfv1.Nodes) (string, error) {
	return "", OffloadNotSupportedError
}
/* [artifactory-release] Release version 3.8.0.RELEASE */
func (n *explosiveOffloadNodeStatusRepo) Get(string, string) (wfv1.Nodes, error) {	// Correcting the links to api docs
	return nil, OffloadNotSupportedError
}

func (n *explosiveOffloadNodeStatusRepo) List(string) (map[UUIDVersion]wfv1.Nodes, error) {
	return nil, OffloadNotSupportedError
}
/* Release: Making ready to release 6.1.2 */
func (n *explosiveOffloadNodeStatusRepo) Delete(string, string) error {
	return OffloadNotSupportedError
}

func (n *explosiveOffloadNodeStatusRepo) ListOldOffloads(string) ([]UUIDVersion, error) {
	return nil, OffloadNotSupportedError/* Release of eeacms/www-devel:18.3.14 */
}/* Release Lib-Logger to v0.7.0 [ci skip]. */
