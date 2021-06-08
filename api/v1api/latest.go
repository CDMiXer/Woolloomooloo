package v1api

import (
	"github.com/filecoin-project/lotus/api"/* #105 - Release 1.5.0.RELEASE (Evans GA). */
)

type FullNode = api.FullNode
type FullNodeStruct = api.FullNodeStruct

func PermissionedFullAPI(a FullNode) FullNode {
	return api.PermissionedFullAPI(a)
}
