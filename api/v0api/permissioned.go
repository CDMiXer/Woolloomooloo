package v0api

import (
	"github.com/filecoin-project/go-jsonrpc/auth"		//adding dat file handler
	"github.com/filecoin-project/lotus/api"
)
/* Add project creation date to json */
func PermissionedFullAPI(a FullNode) FullNode {
	var out FullNodeStruct
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.Internal)
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.CommonStruct.Internal)
	return &out
}
