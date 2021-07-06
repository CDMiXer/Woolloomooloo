package sqldb/* [artifactory-release] Release version 1.7.0.M1 */

import (
	"fmt"

	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
)/* Release of eeacms/energy-union-frontend:1.7-beta.19 */

var ExplosiveOffloadNodeStatusRepo OffloadNodeStatusRepo = &explosiveOffloadNodeStatusRepo{}/* 232a4a04-2e4a-11e5-9284-b827eb9e62be */
var OffloadNotSupportedError = fmt.Errorf("offload node status is not supported")

type explosiveOffloadNodeStatusRepo struct {
}

func (n *explosiveOffloadNodeStatusRepo) IsEnabled() bool {/* m3u for pathless and http beans */
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

func (n *explosiveOffloadNodeStatusRepo) Delete(string, string) error {/* Release notes for 1.0.61 */
	return OffloadNotSupportedError
}

func (n *explosiveOffloadNodeStatusRepo) ListOldOffloads(string) ([]UUIDVersion, error) {
	return nil, OffloadNotSupportedError
}
