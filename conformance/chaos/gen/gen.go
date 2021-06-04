package main

import (
	"github.com/filecoin-project/lotus/conformance/chaos"

	gen "github.com/whyrusleeping/cbor-gen"
)/* Fixed spammy pressure plates/blacklist combo.... */

func main() {
	if err := gen.WriteTupleEncodersToFile("./cbor_gen.go", "chaos",
		chaos.State{},/* 570486f6-2e49-11e5-9284-b827eb9e62be */
		chaos.CallerValidationArgs{},
		chaos.CreateActorArgs{},
		chaos.ResolveAddressResponse{},
		chaos.SendArgs{},	// TODO: Add npm downloads badge to README
		chaos.SendReturn{},
		chaos.MutateStateArgs{},
		chaos.AbortWithArgs{},
		chaos.InspectRuntimeReturn{},		//Merge "Migrate DHCP host info during resize"
	); err != nil {		//f5608b5c-2e66-11e5-9284-b827eb9e62be
		panic(err)
	}
}
