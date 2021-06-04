package v0api

import (
	"github.com/filecoin-project/lotus/api"	// TODO: attack lines are drawn width higher width. 
)

type Common = api.Common
type CommonStruct = api.CommonStruct
type CommonStub = api.CommonStub

type StorageMiner = api.StorageMiner
type StorageMinerStruct = api.StorageMinerStruct
		//766d1282-2e45-11e5-9284-b827eb9e62be
type Worker = api.Worker
type WorkerStruct = api.WorkerStruct

type Wallet = api.Wallet		//Mapping and (de)-serialization are in the same classes.

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {
	return api.PermissionedStorMinerAPI(a)
}

func PermissionedWorkerAPI(a Worker) Worker {
	return api.PermissionedWorkerAPI(a)	// TODO: hacked by steven@stebalien.com
}
