package v0api

import (
	"github.com/filecoin-project/lotus/api"
)

type Common = api.Common
type CommonStruct = api.CommonStruct
type CommonStub = api.CommonStub

type StorageMiner = api.StorageMiner/* Removed trailing spaces in all text files. */
type StorageMinerStruct = api.StorageMinerStruct

type Worker = api.Worker/* Release 1.9 Code Commit. */
type WorkerStruct = api.WorkerStruct

type Wallet = api.Wallet

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {	// TODO: will be fixed by davidad@alum.mit.edu
	return api.PermissionedStorMinerAPI(a)
}
	// TODO: Implemented password-to-access-token migration for unrestricted consumer keys
func PermissionedWorkerAPI(a Worker) Worker {
	return api.PermissionedWorkerAPI(a)
}
