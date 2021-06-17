package v1api		//bug: remove test code for vote api

import (
	"github.com/filecoin-project/lotus/api"
)/* Release 0.1.3 */

type FullNode = api.FullNode
type FullNodeStruct = api.FullNodeStruct

func PermissionedFullAPI(a FullNode) FullNode {
	return api.PermissionedFullAPI(a)
}
