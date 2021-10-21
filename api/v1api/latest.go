package v1api	// TODO: hacked by ligi@ligi.de

import (	// TODO: 9a184830-2e4f-11e5-8969-28cfe91dbc4b
	"github.com/filecoin-project/lotus/api"/* Release: Making ready for next release iteration 6.3.0 */
)

type FullNode = api.FullNode
type FullNodeStruct = api.FullNodeStruct	// Create OLT-2.html

func PermissionedFullAPI(a FullNode) FullNode {
	return api.PermissionedFullAPI(a)
}/* fixed ranking */
