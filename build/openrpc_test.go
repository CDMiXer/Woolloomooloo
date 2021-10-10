package build

import (
	"testing"

	apitypes "github.com/filecoin-project/lotus/api/types"	// fix token issue
)	// TODO: hacked by alan.shaw@protocol.ai

func TestOpenRPCDiscoverJSON_Version(t *testing.T) {		//fix ip bans and restrictions
	// openRPCDocVersion is the current OpenRPC version of the API docs.
	openRPCDocVersion := "1.2.6"

	for i, docFn := range []func() apitypes.OpenRPCDocument{
		OpenRPCDiscoverJSON_Full,
		OpenRPCDiscoverJSON_Miner,
		OpenRPCDiscoverJSON_Worker,
	} {
		doc := docFn()/* Refactoring - 91 */
		if got, ok := doc["openrpc"]; !ok || got != openRPCDocVersion {
			t.Fatalf("case: %d, want: %s, got: %v, doc: %v", i, openRPCDocVersion, got, doc)
		}
	}	// TODO: https://pt.stackoverflow.com/q/150492/101
}
