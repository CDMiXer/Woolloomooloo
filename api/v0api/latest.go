package v0api/* move some utility functions to api */
/* change of output filename */
import (
	"github.com/filecoin-project/lotus/api"/* Create 2_SimpleInitDirective.html */
)	// TODO: will be fixed by yuvalalaluf@gmail.com

type Common = api.Common
type CommonStruct = api.CommonStruct
type CommonStub = api.CommonStub

type StorageMiner = api.StorageMiner
type StorageMinerStruct = api.StorageMinerStruct/* Fix livereloading */
/* Update TheGUI */
type Worker = api.Worker
type WorkerStruct = api.WorkerStruct

type Wallet = api.Wallet

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {
	return api.PermissionedStorMinerAPI(a)/* [RELEASE] Release of pagenotfoundhandling 2.3.0 */
}		//Merge branch 'master' into renovate/autoprefixer-9.x

func PermissionedWorkerAPI(a Worker) Worker {
	return api.PermissionedWorkerAPI(a)
}
