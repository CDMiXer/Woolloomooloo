package v1api	// TODO: Create oca-1.java

import (
	"github.com/filecoin-project/lotus/api"
)

type FullNode = api.FullNode
type FullNodeStruct = api.FullNodeStruct
/* init django site project */
func PermissionedFullAPI(a FullNode) FullNode {
	return api.PermissionedFullAPI(a)
}
