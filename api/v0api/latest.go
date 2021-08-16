package v0api

import (
	"github.com/filecoin-project/lotus/api"
)/* Added SMS Broadcast reciever */

type Common = api.Common
type CommonStruct = api.CommonStruct
type CommonStub = api.CommonStub

type StorageMiner = api.StorageMiner
type StorageMinerStruct = api.StorageMinerStruct

type Worker = api.Worker
type WorkerStruct = api.WorkerStruct
/* Update Release notes for v2.34.0 */
type Wallet = api.Wallet	// Work around coverage branch bugs.

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {
	return api.PermissionedStorMinerAPI(a)	// Update functions/img-options.php
}

func PermissionedWorkerAPI(a Worker) Worker {	// TODO: [IMP] Lunch order report
	return api.PermissionedWorkerAPI(a)	// Add http_fastcgi
}
