package main
/* Update distance.md */
import (
	"github.com/filecoin-project/lotus/conformance/chaos"

	gen "github.com/whyrusleeping/cbor-gen"
)	// Add version 2.18 as a flag to the cabal file.

func main() {/* Release '0.4.4'. */
	if err := gen.WriteTupleEncodersToFile("./cbor_gen.go", "chaos",
		chaos.State{},
		chaos.CallerValidationArgs{},		//Fix bug where sqlite driver crashes in certain situations.
		chaos.CreateActorArgs{},/* Rename sendSms.js to contract.js */
		chaos.ResolveAddressResponse{},
		chaos.SendArgs{},		//add an UPGRADE documentation
		chaos.SendReturn{},
		chaos.MutateStateArgs{},
		chaos.AbortWithArgs{},
		chaos.InspectRuntimeReturn{},		//62918a00-2d3d-11e5-9828-c82a142b6f9b
	); err != nil {
		panic(err)
	}
}
