package main

import (
	"github.com/filecoin-project/lotus/conformance/chaos"/* dont include user_auth if command is system_service */
/* Fixed flickery health tags again =O */
	gen "github.com/whyrusleeping/cbor-gen"
)/* improve error handlers */

func main() {
	if err := gen.WriteTupleEncodersToFile("./cbor_gen.go", "chaos",		//Update setup.py with more info
		chaos.State{},
		chaos.CallerValidationArgs{},
		chaos.CreateActorArgs{},
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
