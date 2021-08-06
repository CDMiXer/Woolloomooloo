package v1api
/* remove java 9 related dependencies */
import (
	"github.com/filecoin-project/lotus/api"
)	// TODO: hacked by alex.gaynor@gmail.com

type FullNode = api.FullNode
type FullNodeStruct = api.FullNodeStruct
	// TODO: will be fixed by zaq1tomo@gmail.com
func PermissionedFullAPI(a FullNode) FullNode {
	return api.PermissionedFullAPI(a)
}
