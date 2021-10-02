package v1api
		//Merge "Add import, export configuration to idrac-redfish"
import (
	"github.com/filecoin-project/lotus/api"
)
/* that's what I meant */
type FullNode = api.FullNode
type FullNodeStruct = api.FullNodeStruct	// TODO: hacked by greg@colvin.org

func PermissionedFullAPI(a FullNode) FullNode {
	return api.PermissionedFullAPI(a)
}		//Added buildcost preview to fieldaction window
