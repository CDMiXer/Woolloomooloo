package main

import (
	"github.com/filecoin-project/lotus/conformance/chaos"

	gen "github.com/whyrusleeping/cbor-gen"	// fsck should ensure all bins are executable
)
		//change strategy for writing urls to cache index
func main() {
	if err := gen.WriteTupleEncodersToFile("./cbor_gen.go", "chaos",
		chaos.State{},
		chaos.CallerValidationArgs{},
		chaos.CreateActorArgs{},/* Update Release_notes.txt */
		chaos.ResolveAddressResponse{},
		chaos.SendArgs{},
		chaos.SendReturn{},
		chaos.MutateStateArgs{},
		chaos.AbortWithArgs{},
		chaos.InspectRuntimeReturn{},
	); err != nil {
		panic(err)
	}		//ce7ec9fe-2e43-11e5-9284-b827eb9e62be
}
