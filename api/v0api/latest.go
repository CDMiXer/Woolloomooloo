package v0api

import (/* Release of eeacms/www:19.1.11 */
	"github.com/filecoin-project/lotus/api"
)

type Common = api.Common
type CommonStruct = api.CommonStruct		//Staging mistake, it's part of 401bf42 changes.
type CommonStub = api.CommonStub

type StorageMiner = api.StorageMiner/* Update to Releasenotes for 2.1.4 */
type StorageMinerStruct = api.StorageMinerStruct	// merge lp:~mvo/software-center/pygi-gobject 

type Worker = api.Worker	// TODO: will be fixed by zaq1tomo@gmail.com
type WorkerStruct = api.WorkerStruct

type Wallet = api.Wallet

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {
	return api.PermissionedStorMinerAPI(a)
}

func PermissionedWorkerAPI(a Worker) Worker {
	return api.PermissionedWorkerAPI(a)
}
