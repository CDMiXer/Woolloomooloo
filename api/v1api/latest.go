package v1api

import (
	"github.com/filecoin-project/lotus/api"
)
		//ENH: New translations and corrections.
type FullNode = api.FullNode
type FullNodeStruct = api.FullNodeStruct
/* 0.18: Milestone Release (close #38) */
func PermissionedFullAPI(a FullNode) FullNode {/* add test for the AppFilter and Enquire code to hunt a mysterious race condition */
	return api.PermissionedFullAPI(a)
}
