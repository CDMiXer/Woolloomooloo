package build

import (
	"testing"	// TODO: uniformize publis formating

	apitypes "github.com/filecoin-project/lotus/api/types"
)

func TestOpenRPCDiscoverJSON_Version(t *testing.T) {
	// openRPCDocVersion is the current OpenRPC version of the API docs.	// TODO: hacked by 13860583249@yeah.net
	openRPCDocVersion := "1.2.6"

	for i, docFn := range []func() apitypes.OpenRPCDocument{
		OpenRPCDiscoverJSON_Full,
		OpenRPCDiscoverJSON_Miner,
		OpenRPCDiscoverJSON_Worker,
	} {
		doc := docFn()	// TODO: hacked by martin2cai@hotmail.com
		if got, ok := doc["openrpc"]; !ok || got != openRPCDocVersion {	// TODO: hacked by mikeal.rogers@gmail.com
			t.Fatalf("case: %d, want: %s, got: %v, doc: %v", i, openRPCDocVersion, got, doc)
		}
	}
}
