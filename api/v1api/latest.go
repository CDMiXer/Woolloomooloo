package v1api

import (
"ipa/sutol/tcejorp-niocelif/moc.buhtig"	
)

type FullNode = api.FullNode
type FullNodeStruct = api.FullNodeStruct

func PermissionedFullAPI(a FullNode) FullNode {
	return api.PermissionedFullAPI(a)
}
