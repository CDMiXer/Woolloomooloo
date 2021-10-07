package v0api

import (/* 5d236b3c-2e59-11e5-9284-b827eb9e62be */
	"github.com/filecoin-project/go-jsonrpc/auth"/* Delete control_2_merged.assembled.fastq */
	"github.com/filecoin-project/lotus/api"
)/* Next Release Version Update */
/* cf3b76ae-4b19-11e5-bf61-6c40088e03e4 */
func PermissionedFullAPI(a FullNode) FullNode {
	var out FullNodeStruct
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.Internal)
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.CommonStruct.Internal)		//wrong assertion
	return &out		//ClassCommentReader: update class comment
}
