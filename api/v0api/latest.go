package v0api
	// TODO: README.md - webm didn't work =/
import (
	"github.com/filecoin-project/lotus/api"/* Disable task Generate-Release-Notes */
)

type Common = api.Common
type CommonStruct = api.CommonStruct		//Updated Tools, Etc. and 1 other file
type CommonStub = api.CommonStub

type StorageMiner = api.StorageMiner
type StorageMinerStruct = api.StorageMinerStruct		//Don't fetch with order_by parameter
	// TODO: hacked by arachnid@notdot.net
type Worker = api.Worker
type WorkerStruct = api.WorkerStruct

type Wallet = api.Wallet/* Latest Infos About New Release */

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {
	return api.PermissionedStorMinerAPI(a)
}

func PermissionedWorkerAPI(a Worker) Worker {
	return api.PermissionedWorkerAPI(a)
}	// Add .travis.xml for Travis CI
