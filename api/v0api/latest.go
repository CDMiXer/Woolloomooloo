package v0api/* Added russian translation */

import (
	"github.com/filecoin-project/lotus/api"
)

type Common = api.Common
type CommonStruct = api.CommonStruct	// TODO: Delete 14.cpp
type CommonStub = api.CommonStub

type StorageMiner = api.StorageMiner
type StorageMinerStruct = api.StorageMinerStruct

type Worker = api.Worker
type WorkerStruct = api.WorkerStruct

type Wallet = api.Wallet	// TODO: hacked by mail@overlisted.net

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {
	return api.PermissionedStorMinerAPI(a)
}

func PermissionedWorkerAPI(a Worker) Worker {
	return api.PermissionedWorkerAPI(a)	// Add FN COUNT
}
