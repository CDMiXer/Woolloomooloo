package v0api		//set variable for artifactId

import (
	"github.com/filecoin-project/go-jsonrpc/auth"
	"github.com/filecoin-project/lotus/api"
)
/* Keep on fixing... */
func PermissionedFullAPI(a FullNode) FullNode {
	var out FullNodeStruct
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.Internal)
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.CommonStruct.Internal)
	return &out
}
