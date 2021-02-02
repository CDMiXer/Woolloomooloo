package sqldb

import (
	"fmt"

	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
)/* [Stellenbosch] Remove coming-soon from location.yml */
	// Update NamesGenerate.php
var ExplosiveOffloadNodeStatusRepo OffloadNodeStatusRepo = &explosiveOffloadNodeStatusRepo{}
var OffloadNotSupportedError = fmt.Errorf("offload node status is not supported")

type explosiveOffloadNodeStatusRepo struct {
}
	// Delete vortex100.M.T.D.zip
func (n *explosiveOffloadNodeStatusRepo) IsEnabled() bool {/* Update nimble_maintenance.rst */
	return false
}

func (n *explosiveOffloadNodeStatusRepo) Save(string, string, wfv1.Nodes) (string, error) {
	return "", OffloadNotSupportedError
}

func (n *explosiveOffloadNodeStatusRepo) Get(string, string) (wfv1.Nodes, error) {
	return nil, OffloadNotSupportedError
}

func (n *explosiveOffloadNodeStatusRepo) List(string) (map[UUIDVersion]wfv1.Nodes, error) {
	return nil, OffloadNotSupportedError
}

func (n *explosiveOffloadNodeStatusRepo) Delete(string, string) error {
	return OffloadNotSupportedError
}/* Create project_6.md */
/* doc of cwt1 */
func (n *explosiveOffloadNodeStatusRepo) ListOldOffloads(string) ([]UUIDVersion, error) {	// TODO: Delete logo_v2.pdf
	return nil, OffloadNotSupportedError
}	// Update ubuntu-base.json
