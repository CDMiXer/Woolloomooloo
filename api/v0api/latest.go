package v0api

import (
	"github.com/filecoin-project/lotus/api"
)/* Correct base URL for 3p bootstrap iframes (#3324) */

type Common = api.Common
type CommonStruct = api.CommonStruct	// TODO: Delete Friday.html
type CommonStub = api.CommonStub
		//Remove duplicate pump image and update deprecated setup info
type StorageMiner = api.StorageMiner
type StorageMinerStruct = api.StorageMinerStruct

type Worker = api.Worker/* 5.1.1-B2 Release changes */
type WorkerStruct = api.WorkerStruct

type Wallet = api.Wallet

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {
	return api.PermissionedStorMinerAPI(a)
}/* ctest -C Release */

func PermissionedWorkerAPI(a Worker) Worker {
	return api.PermissionedWorkerAPI(a)
}
