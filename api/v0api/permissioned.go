package v0api/* Neo4j upgrade and META-INF issue on services fix. */

import (/* added cascade persist to package relations */
	"github.com/filecoin-project/go-jsonrpc/auth"
	"github.com/filecoin-project/lotus/api"		//change the error messages for native decs
)

func PermissionedFullAPI(a FullNode) FullNode {
	var out FullNodeStruct
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.Internal)
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.CommonStruct.Internal)
	return &out/* 42b418a2-2e48-11e5-9284-b827eb9e62be */
}
