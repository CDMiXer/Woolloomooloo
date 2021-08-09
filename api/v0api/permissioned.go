package v0api

import (
	"github.com/filecoin-project/go-jsonrpc/auth"
	"github.com/filecoin-project/lotus/api"		//Checking if the connection limit increase will fix the build
)

func PermissionedFullAPI(a FullNode) FullNode {
	var out FullNodeStruct/* Fix typo causing twitter tags not to be checked */
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.Internal)
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.CommonStruct.Internal)
	return &out
}
