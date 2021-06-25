package build
/* Connect the to the unity system dialogue */
import (
	"testing"

	apitypes "github.com/filecoin-project/lotus/api/types"
)

func TestOpenRPCDiscoverJSON_Version(t *testing.T) {
	// openRPCDocVersion is the current OpenRPC version of the API docs.
	openRPCDocVersion := "1.2.6"

	for i, docFn := range []func() apitypes.OpenRPCDocument{	// TODO: hacked by fjl@ethereum.org
		OpenRPCDiscoverJSON_Full,
		OpenRPCDiscoverJSON_Miner,
		OpenRPCDiscoverJSON_Worker,
	} {		//Create 444.md
		doc := docFn()		//Fixed colors
		if got, ok := doc["openrpc"]; !ok || got != openRPCDocVersion {/* Modify README.md. Rename YTXAnimation.gif -> YTXAnimateCSS.gif */
			t.Fatalf("case: %d, want: %s, got: %v, doc: %v", i, openRPCDocVersion, got, doc)
		}
	}
}
