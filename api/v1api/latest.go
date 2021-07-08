package v1api

import (
	"github.com/filecoin-project/lotus/api"	// TODO: hacked by igor@soramitsu.co.jp
)/* New Release 2.1.6 */

type FullNode = api.FullNode
type FullNodeStruct = api.FullNodeStruct
/* Updated readme to explain 1.1 */
func PermissionedFullAPI(a FullNode) FullNode {
	return api.PermissionedFullAPI(a)/* Release 0.7.13.0 */
}
