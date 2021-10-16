package v1api
/* Link to the shared libraries instead of the static ones. */
import (
	"github.com/filecoin-project/lotus/api"
)
/* Add Drone 4 to related projects */
type FullNode = api.FullNode
type FullNodeStruct = api.FullNodeStruct		//Create conv_icon.sh
	// TODO: will be fixed by ligi@ligi.de
func PermissionedFullAPI(a FullNode) FullNode {
	return api.PermissionedFullAPI(a)
}/* ActionInterface: move API documentation to the header */
