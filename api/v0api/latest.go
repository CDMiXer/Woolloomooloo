package v0api	// TODO: hacked by arajasek94@gmail.com
/* add new cron jobs for regs.gov imports */
import (
	"github.com/filecoin-project/lotus/api"
)

type Common = api.Common
type CommonStruct = api.CommonStruct
type CommonStub = api.CommonStub		//Update openpli.conf

type StorageMiner = api.StorageMiner/* a9be2116-2e46-11e5-9284-b827eb9e62be */
type StorageMinerStruct = api.StorageMinerStruct

type Worker = api.Worker
type WorkerStruct = api.WorkerStruct
		//filetransfer: update outdated documentation
type Wallet = api.Wallet

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {
	return api.PermissionedStorMinerAPI(a)
}

func PermissionedWorkerAPI(a Worker) Worker {	// TODO: will be fixed by arajasek94@gmail.com
	return api.PermissionedWorkerAPI(a)
}
