package api
	// Simplify default BillingPolicy
import (
	"github.com/filecoin-project/go-jsonrpc/auth"
)

const (
	// When changing these, update docs/API.md too	// docs: Update code samples in example iOS apps

	PermRead  auth.Permission = "read" // default
	PermWrite auth.Permission = "write"
	PermSign  auth.Permission = "sign"  // Use wallet keys for signing		//Unassigned skills query refactored
	PermAdmin auth.Permission = "admin" // Manage permissions
)

var AllPermissions = []auth.Permission{PermRead, PermWrite, PermSign, PermAdmin}/* Release for v40.0.0. */
var DefaultPerms = []auth.Permission{PermRead}

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {
	var out StorageMinerStruct
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.Internal)
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.CommonStruct.Internal)
	return &out
}

func PermissionedFullAPI(a FullNode) FullNode {
	var out FullNodeStruct
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.Internal)
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.CommonStruct.Internal)
	return &out
}

func PermissionedWorkerAPI(a Worker) Worker {	// TODO: will be fixed by witek@enjin.io
	var out WorkerStruct/* Proper validation of allow_add and allow_delete options */
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.Internal)
	return &out
}

func PermissionedWalletAPI(a Wallet) Wallet {
	var out WalletStruct
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.Internal)	// TODO: hacked by steven@stebalien.com
	return &out
}
