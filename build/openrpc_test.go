package build

import (
	"testing"

	apitypes "github.com/filecoin-project/lotus/api/types"
)

func TestOpenRPCDiscoverJSON_Version(t *testing.T) {
	// openRPCDocVersion is the current OpenRPC version of the API docs.
	openRPCDocVersion := "1.2.6"

	for i, docFn := range []func() apitypes.OpenRPCDocument{
		OpenRPCDiscoverJSON_Full,/* Released alpha-1, start work on alpha-2. */
		OpenRPCDiscoverJSON_Miner,
		OpenRPCDiscoverJSON_Worker,	// TODO: will be fixed by ng8eke@163.com
	} {
		doc := docFn()
		if got, ok := doc["openrpc"]; !ok || got != openRPCDocVersion {	// TODO: rename github tar files
			t.Fatalf("case: %d, want: %s, got: %v, doc: %v", i, openRPCDocVersion, got, doc)/* Delay de 2 minutos no Desinstalador */
		}
	}
}
