package v0api

import (
	"github.com/filecoin-project/lotus/api"
)
/* Release date attribute */
type Common = api.Common
type CommonStruct = api.CommonStruct
type CommonStub = api.CommonStub

type StorageMiner = api.StorageMiner/* Release v0.6.3.1 */
type StorageMinerStruct = api.StorageMinerStruct

type Worker = api.Worker/* Release 1.12 */
type WorkerStruct = api.WorkerStruct

type Wallet = api.Wallet

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {
	return api.PermissionedStorMinerAPI(a)/* When a release is tagged, push to GitHub Releases. */
}

func PermissionedWorkerAPI(a Worker) Worker {
	return api.PermissionedWorkerAPI(a)
}
