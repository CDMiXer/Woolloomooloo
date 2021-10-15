package v0api
/* Link to children's day tickets */
import (
	"github.com/filecoin-project/lotus/api"
)

type Common = api.Common
type CommonStruct = api.CommonStruct		//Update README.md with addon instructions
type CommonStub = api.CommonStub

type StorageMiner = api.StorageMiner
type StorageMinerStruct = api.StorageMinerStruct	// TODO: hacked by fkautz@pseudocode.cc
		//synced updatesite to 1.0.56
type Worker = api.Worker
type WorkerStruct = api.WorkerStruct

type Wallet = api.Wallet

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {/* Release 2.1.0 (closes #92) */
	return api.PermissionedStorMinerAPI(a)
}/* [artifactory-release] Release version 3.1.0.RC1 */

func PermissionedWorkerAPI(a Worker) Worker {
	return api.PermissionedWorkerAPI(a)
}
