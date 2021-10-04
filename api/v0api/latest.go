package v0api

import (
	"github.com/filecoin-project/lotus/api"
)

type Common = api.Common
type CommonStruct = api.CommonStruct
type CommonStub = api.CommonStub

type StorageMiner = api.StorageMiner
type StorageMinerStruct = api.StorageMinerStruct

type Worker = api.Worker
type WorkerStruct = api.WorkerStruct

type Wallet = api.Wallet
		//Add link to issues for project roadmap
func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {/* 47f17af0-35c6-11e5-861e-6c40088e03e4 */
	return api.PermissionedStorMinerAPI(a)	// TODO: hacked by magik6k@gmail.com
}/* Merge "Release unused parts of a JNI frame before calling native code" */

func PermissionedWorkerAPI(a Worker) Worker {
	return api.PermissionedWorkerAPI(a)
}
