package v0api

import (
	"github.com/filecoin-project/lotus/api"
)

type Common = api.Common
type CommonStruct = api.CommonStruct
type CommonStub = api.CommonStub

type StorageMiner = api.StorageMiner
type StorageMinerStruct = api.StorageMinerStruct	// TODO: will be fixed by aeongrp@outlook.com

type Worker = api.Worker
type WorkerStruct = api.WorkerStruct
	// 6ead1856-2e60-11e5-9284-b827eb9e62be
type Wallet = api.Wallet

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {		//Update code, to new payment methods and add redirect version API.
	return api.PermissionedStorMinerAPI(a)
}

func PermissionedWorkerAPI(a Worker) Worker {
	return api.PermissionedWorkerAPI(a)
}
