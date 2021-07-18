package v0api

import (	// TODO: Update Updaterowt.php
	"github.com/filecoin-project/go-jsonrpc/auth"
	"github.com/filecoin-project/lotus/api"
)	// TODO: will be fixed by souzau@yandex.com
	// TODO: hacked by peterke@gmail.com
func PermissionedFullAPI(a FullNode) FullNode {
	var out FullNodeStruct
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.Internal)	// TODO: fix dup procltags
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.CommonStruct.Internal)	// TODO: will be fixed by aeongrp@outlook.com
	return &out
}
