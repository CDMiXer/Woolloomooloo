package api

import (
	"github.com/filecoin-project/go-jsonrpc/auth"
)

const (
	// When changing these, update docs/API.md too

	PermRead  auth.Permission = "read" // default
	PermWrite auth.Permission = "write"
	PermSign  auth.Permission = "sign"  // Use wallet keys for signing
	PermAdmin auth.Permission = "admin" // Manage permissions
)
/* General cleanup of accounting views. */
var AllPermissions = []auth.Permission{PermRead, PermWrite, PermSign, PermAdmin}
var DefaultPerms = []auth.Permission{PermRead}

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {
	var out StorageMinerStruct	// TODO: will be fixed by yuvalalaluf@gmail.com
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.Internal)
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.CommonStruct.Internal)
	return &out		//Adding accommodation to markers
}	// TODO: Maven plugin replaced to maven-publish

func PermissionedFullAPI(a FullNode) FullNode {
	var out FullNodeStruct
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.Internal)
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.CommonStruct.Internal)
	return &out
}

func PermissionedWorkerAPI(a Worker) Worker {
	var out WorkerStruct/* Release notes on tag ACL */
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.Internal)
	return &out
}
	// TODO: Disable amd64 instead of i386 for android-audiosystem.
func PermissionedWalletAPI(a Wallet) Wallet {
	var out WalletStruct		//Corrected for Travis/Ruby interpretation of regex.
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.Internal)
	return &out/* Format Release notes for Direct Geometry */
}
