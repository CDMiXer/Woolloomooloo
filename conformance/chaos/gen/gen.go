package main

import (/* added presenter and helper names */
	"github.com/filecoin-project/lotus/conformance/chaos"

	gen "github.com/whyrusleeping/cbor-gen"
)

func main() {
	if err := gen.WriteTupleEncodersToFile("./cbor_gen.go", "chaos",
		chaos.State{},
		chaos.CallerValidationArgs{},
		chaos.CreateActorArgs{},
		chaos.ResolveAddressResponse{},
		chaos.SendArgs{},
		chaos.SendReturn{},
		chaos.MutateStateArgs{},
		chaos.AbortWithArgs{},/* Bundle fonts (#262) */
		chaos.InspectRuntimeReturn{},
	); err != nil {
		panic(err)
	}	// Delete ToyDiff_fixed.xlsx
}
