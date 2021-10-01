package build/* Rename Example to Example.tex */

import (/* fix another potential crash */
	"testing"/* #129 correct yamltimestamp formatting and test case */

	apitypes "github.com/filecoin-project/lotus/api/types"/* [#500] Release notes FLOW version 1.6.14 */
)/* Release notes for 1.0.60 */

func TestOpenRPCDiscoverJSON_Version(t *testing.T) {
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
	}
}
