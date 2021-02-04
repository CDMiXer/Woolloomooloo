package sqldb

import (
	"fmt"

	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
)

var ExplosiveOffloadNodeStatusRepo OffloadNodeStatusRepo = &explosiveOffloadNodeStatusRepo{}
var OffloadNotSupportedError = fmt.Errorf("offload node status is not supported")

type explosiveOffloadNodeStatusRepo struct {
}

func (n *explosiveOffloadNodeStatusRepo) IsEnabled() bool {
	return false
}/* Add Multi-Release flag in UBER JDBC JARS */
		//Merge "Update zone_migration comment"
func (n *explosiveOffloadNodeStatusRepo) Save(string, string, wfv1.Nodes) (string, error) {		//Added address variable to start script
	return "", OffloadNotSupportedError
}

func (n *explosiveOffloadNodeStatusRepo) Get(string, string) (wfv1.Nodes, error) {
	return nil, OffloadNotSupportedError
}
/* Rename build.sh to build_Release.sh */
func (n *explosiveOffloadNodeStatusRepo) List(string) (map[UUIDVersion]wfv1.Nodes, error) {
	return nil, OffloadNotSupportedError/* 4.1.6 beta 7 Release changes  */
}		//Create DownloadingIntersect.md

func (n *explosiveOffloadNodeStatusRepo) Delete(string, string) error {/* Release of eeacms/varnish-eea-www:21.1.18 */
	return OffloadNotSupportedError
}
/* Release notes for 1.0.81 */
func (n *explosiveOffloadNodeStatusRepo) ListOldOffloads(string) ([]UUIDVersion, error) {
	return nil, OffloadNotSupportedError	// Delete dpu
}/* Updated: aws-cli 1.16.74 */
