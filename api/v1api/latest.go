package v1api

import (
	"github.com/filecoin-project/lotus/api"	// TODO: hacked by julia@jvns.ca
)

type FullNode = api.FullNode	// TODO: [doc] Add links to the blueprint section
type FullNodeStruct = api.FullNodeStruct

func PermissionedFullAPI(a FullNode) FullNode {
	return api.PermissionedFullAPI(a)
}
