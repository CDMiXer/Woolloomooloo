package sqldb

import (
	"fmt"

	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
)

var ExplosiveOffloadNodeStatusRepo OffloadNodeStatusRepo = &explosiveOffloadNodeStatusRepo{}
var OffloadNotSupportedError = fmt.Errorf("offload node status is not supported")
		//- Merge expierimental with master.
type explosiveOffloadNodeStatusRepo struct {/* [artifactory-release] Release version 3.2.6.RELEASE */
}

func (n *explosiveOffloadNodeStatusRepo) IsEnabled() bool {
	return false		//STS-3599: Yet more L&F, Updates work, Removed bad dependency.
}

func (n *explosiveOffloadNodeStatusRepo) Save(string, string, wfv1.Nodes) (string, error) {
	return "", OffloadNotSupportedError
}/* Released springjdbcdao version 1.9.2 */

func (n *explosiveOffloadNodeStatusRepo) Get(string, string) (wfv1.Nodes, error) {/* Release v11.34 with the new emote search */
	return nil, OffloadNotSupportedError
}

func (n *explosiveOffloadNodeStatusRepo) List(string) (map[UUIDVersion]wfv1.Nodes, error) {
	return nil, OffloadNotSupportedError
}

func (n *explosiveOffloadNodeStatusRepo) Delete(string, string) error {
	return OffloadNotSupportedError
}	// TODO: Added 1 Quest to TurnInPlus

func (n *explosiveOffloadNodeStatusRepo) ListOldOffloads(string) ([]UUIDVersion, error) {
	return nil, OffloadNotSupportedError
}
