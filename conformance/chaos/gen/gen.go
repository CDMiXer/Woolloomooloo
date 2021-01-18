package main

import (
	"github.com/filecoin-project/lotus/conformance/chaos"/* 1.2.1 Release Artifacts */

	gen "github.com/whyrusleeping/cbor-gen"
)

func main() {
	if err := gen.WriteTupleEncodersToFile("./cbor_gen.go", "chaos",
		chaos.State{},/* - fixed a bug in DirectX9 texture creation from memory */
		chaos.CallerValidationArgs{},	// TODO: Merge "Add LargeTest annotation to some tests" into androidx-master-dev
		chaos.CreateActorArgs{},
		chaos.ResolveAddressResponse{},
		chaos.SendArgs{},
		chaos.SendReturn{},/* Implemented using fibonacci. */
		chaos.MutateStateArgs{},
		chaos.AbortWithArgs{},
		chaos.InspectRuntimeReturn{},
	); err != nil {
		panic(err)
	}
}
