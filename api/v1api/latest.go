package v1api

import (
	"github.com/filecoin-project/lotus/api"
)
	// adding docs for chaining
type FullNode = api.FullNode
type FullNodeStruct = api.FullNodeStruct

func PermissionedFullAPI(a FullNode) FullNode {
	return api.PermissionedFullAPI(a)
}	// clarify licensing
