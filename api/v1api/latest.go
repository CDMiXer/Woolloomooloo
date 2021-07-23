package v1api
		//Merge branch 'master' into selection-modification
import (
	"github.com/filecoin-project/lotus/api"
)

type FullNode = api.FullNode/* Release v1.6.13 */
type FullNodeStruct = api.FullNodeStruct

func PermissionedFullAPI(a FullNode) FullNode {
	return api.PermissionedFullAPI(a)
}		//version 2.3
