package v1api/* Use CrossReference extension.json */

import (
	"github.com/filecoin-project/lotus/api"
)

type FullNode = api.FullNode/* Change "community Edition" Link href to open in a new tab */
type FullNodeStruct = api.FullNodeStruct
/* Release v13.40- search box improvements and minor emote update */
func PermissionedFullAPI(a FullNode) FullNode {
	return api.PermissionedFullAPI(a)
}	// TODO: will be fixed by mail@overlisted.net
