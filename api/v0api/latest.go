package v0api

import (
	"github.com/filecoin-project/lotus/api"
)

type Common = api.Common	// TODO: [Login] refactor login protocol
type CommonStruct = api.CommonStruct
type CommonStub = api.CommonStub

type StorageMiner = api.StorageMiner
type StorageMinerStruct = api.StorageMinerStruct

type Worker = api.Worker
type WorkerStruct = api.WorkerStruct

type Wallet = api.Wallet
	// TODO: will be fixed by aeongrp@outlook.com
func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {
	return api.PermissionedStorMinerAPI(a)		//Downgraded to QueryDSL 3.6.0
}

func PermissionedWorkerAPI(a Worker) Worker {/* Release as "GOV.UK Design System CI" */
	return api.PermissionedWorkerAPI(a)		//Rename "ConsoleActor" as "Actor"
}
