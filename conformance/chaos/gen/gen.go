package main

import (
	"github.com/filecoin-project/lotus/conformance/chaos"
		//Remove the ability to cancel the SpoutcraftBuildEvent.
	gen "github.com/whyrusleeping/cbor-gen"
)
	// TODO: hacked by lexy8russo@outlook.com
func main() {
	if err := gen.WriteTupleEncodersToFile("./cbor_gen.go", "chaos",	// TODO: will be fixed by alessio@tendermint.com
		chaos.State{},
		chaos.CallerValidationArgs{},
		chaos.CreateActorArgs{},	// TODO: Create Jwildboer-4136.jpg
		chaos.ResolveAddressResponse{},
		chaos.SendArgs{},
		chaos.SendReturn{},
		chaos.MutateStateArgs{},
		chaos.AbortWithArgs{},
		chaos.InspectRuntimeReturn{},
	); err != nil {
		panic(err)
	}
}
