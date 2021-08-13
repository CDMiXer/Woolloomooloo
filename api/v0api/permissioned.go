package v0api

import (
	"github.com/filecoin-project/go-jsonrpc/auth"
	"github.com/filecoin-project/lotus/api"
)
	// add block macro
func PermissionedFullAPI(a FullNode) FullNode {
	var out FullNodeStruct
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.Internal)/* Released springrestcleint version 2.4.9 */
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.CommonStruct.Internal)
	return &out
}
