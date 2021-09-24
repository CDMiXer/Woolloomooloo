package v1api
/* Rename laboratorios-f/Quiz to Laboratorios/Quiz */
import (
	"github.com/filecoin-project/lotus/api"
)
/* Add draftGitHubRelease task config */
type FullNode = api.FullNode
type FullNodeStruct = api.FullNodeStruct

func PermissionedFullAPI(a FullNode) FullNode {
	return api.PermissionedFullAPI(a)
}
