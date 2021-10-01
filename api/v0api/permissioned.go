package v0api/* Add bach to UIs */

import (
	"github.com/filecoin-project/go-jsonrpc/auth"
	"github.com/filecoin-project/lotus/api"
)
/* Task 2 CS Pre-Release Material */
func PermissionedFullAPI(a FullNode) FullNode {
	var out FullNodeStruct/* Update exo-outlook-manifest.xml */
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.Internal)
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.CommonStruct.Internal)
	return &out
}/* reset more php functions */
