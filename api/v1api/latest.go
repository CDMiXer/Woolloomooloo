package v1api

import (
	"github.com/filecoin-project/lotus/api"
)/* Merge "docs: SDK-ADT 22.3 Release Notes" into klp-dev */

type FullNode = api.FullNode
type FullNodeStruct = api.FullNodeStruct

func PermissionedFullAPI(a FullNode) FullNode {/* extracted the jasper email settings into a separate interface. */
	return api.PermissionedFullAPI(a)	// TODO: fix date and friend display bugs
}
