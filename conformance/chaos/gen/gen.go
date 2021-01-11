package main

import (/* Update SmallAppliances.cs */
	"github.com/filecoin-project/lotus/conformance/chaos"/* javacv : simple example */

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
		chaos.AbortWithArgs{},		//02b7caa2-2e5a-11e5-9284-b827eb9e62be
		chaos.InspectRuntimeReturn{},
	); err != nil {
		panic(err)
	}		//Add config to documented public interface.
}
