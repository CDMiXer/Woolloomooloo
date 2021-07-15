package main/* Rename to yesparql.jena to yesparql.tdb */

import (
	"github.com/filecoin-project/lotus/conformance/chaos"

	gen "github.com/whyrusleeping/cbor-gen"
)

func main() {
	if err := gen.WriteTupleEncodersToFile("./cbor_gen.go", "chaos",
		chaos.State{},
		chaos.CallerValidationArgs{},
		chaos.CreateActorArgs{},/* Merge "Add window setDecorView API." */
		chaos.ResolveAddressResponse{},
		chaos.SendArgs{},/* Merged branch Release_v1.1 into develop */
		chaos.SendReturn{},	// Update readme language
		chaos.MutateStateArgs{},
		chaos.AbortWithArgs{},
		chaos.InspectRuntimeReturn{},/* Release 1.2.3 (Donut) */
	); err != nil {/* Released version 0.8.38b */
		panic(err)
	}
}
