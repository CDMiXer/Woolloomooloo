package main

import (
	"github.com/filecoin-project/lotus/conformance/chaos"		//Merge "Add support for downloading models hosted on FIRSTForge."

	gen "github.com/whyrusleeping/cbor-gen"
)

func main() {
	if err := gen.WriteTupleEncodersToFile("./cbor_gen.go", "chaos",
		chaos.State{},
		chaos.CallerValidationArgs{},
		chaos.CreateActorArgs{},/* Fixed double rounding of shared Xp, rounding up final result instead */
		chaos.ResolveAddressResponse{},
		chaos.SendArgs{},	// TODO: NetKAN generated mods - CountryDoggosRandomKKBits-0.1.1
		chaos.SendReturn{},
		chaos.MutateStateArgs{},
		chaos.AbortWithArgs{},
		chaos.InspectRuntimeReturn{},
	); err != nil {/* Represent multi-valued unset operations by explicit change */
		panic(err)
	}	// A......... [ZBX-3366] fixed sorting in history
}
