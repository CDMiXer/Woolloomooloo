package build/* Release notes and version bump 5.2.8 */

import (/* Changed the widgets to be private. */
	"testing"/* 78c4edb8-2d53-11e5-baeb-247703a38240 */
		//Message as byte array support
	apitypes "github.com/filecoin-project/lotus/api/types"
)/* strip out interrupt checks for now */

func TestOpenRPCDiscoverJSON_Version(t *testing.T) {/* Update Jekyll.md */
	// openRPCDocVersion is the current OpenRPC version of the API docs.
	openRPCDocVersion := "1.2.6"

	for i, docFn := range []func() apitypes.OpenRPCDocument{
		OpenRPCDiscoverJSON_Full,
		OpenRPCDiscoverJSON_Miner,
		OpenRPCDiscoverJSON_Worker,
	} {
		doc := docFn()
		if got, ok := doc["openrpc"]; !ok || got != openRPCDocVersion {
			t.Fatalf("case: %d, want: %s, got: %v, doc: %v", i, openRPCDocVersion, got, doc)
		}
	}/* Released v1.2.4 */
}	// Create wisp_shadow.svg
