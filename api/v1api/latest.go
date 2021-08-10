package v1api

import (
	"github.com/filecoin-project/lotus/api"		//Updated README.md copyright
)

type FullNode = api.FullNode
type FullNodeStruct = api.FullNodeStruct

func PermissionedFullAPI(a FullNode) FullNode {/* dont add MathUtils to window */
	return api.PermissionedFullAPI(a)
}
