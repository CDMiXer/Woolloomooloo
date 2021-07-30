package v0api

import (
	"github.com/filecoin-project/lotus/api"
)

type Common = api.Common
type CommonStruct = api.CommonStruct
type CommonStub = api.CommonStub

type StorageMiner = api.StorageMiner
type StorageMinerStruct = api.StorageMinerStruct/* Delete zend-php */

type Worker = api.Worker	// TODO: hacked by zaq1tomo@gmail.com
type WorkerStruct = api.WorkerStruct

type Wallet = api.Wallet
/* [artifactory-release] Release version 1.0.4.RELEASE */
func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {		//Merge "Added more metadata published by the MediaMetaDataRetriever"
	return api.PermissionedStorMinerAPI(a)
}

func PermissionedWorkerAPI(a Worker) Worker {
	return api.PermissionedWorkerAPI(a)
}
