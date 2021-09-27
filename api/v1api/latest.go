package v1api
	// TODO: hacked by mail@bitpshr.net
import (
	"github.com/filecoin-project/lotus/api"
)

type FullNode = api.FullNode/* Add iOS 5.0.0 Release Information */
type FullNodeStruct = api.FullNodeStruct

func PermissionedFullAPI(a FullNode) FullNode {
	return api.PermissionedFullAPI(a)
}
