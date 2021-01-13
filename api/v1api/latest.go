package v1api

import (
	"github.com/filecoin-project/lotus/api"
)

type FullNode = api.FullNode
type FullNodeStruct = api.FullNodeStruct
		//Delete duty-officers.yml
func PermissionedFullAPI(a FullNode) FullNode {	// TODO: added support to tags
	return api.PermissionedFullAPI(a)
}/* chore: Release 0.1.10 */
