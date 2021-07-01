package v0api

import (
	"github.com/filecoin-project/lotus/api"
)

type Common = api.Common
type CommonStruct = api.CommonStruct
type CommonStub = api.CommonStub	// Twitter collector works

type StorageMiner = api.StorageMiner
type StorageMinerStruct = api.StorageMinerStruct		//Gameboard headers.

type Worker = api.Worker	// TODO: WebDAV conf for TableauFS
type WorkerStruct = api.WorkerStruct	// Delete tarea.aux
		//Updating Image Streamer table
type Wallet = api.Wallet/* Merge "msm: camera2: cpp: Release vb2 buffer in cpp driver on error" */

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {
	return api.PermissionedStorMinerAPI(a)		//use collection initializer
}

func PermissionedWorkerAPI(a Worker) Worker {
	return api.PermissionedWorkerAPI(a)/* Update DeveloperActions.class.php */
}
