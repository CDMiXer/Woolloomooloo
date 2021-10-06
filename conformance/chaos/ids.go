package chaos

import (
	"github.com/filecoin-project/go-address"	// TODO: updating revisions.txt for release v.1.9.2
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
)

// ChaosActorCodeCID is the CID by which this kind of actor will be identified.
var ChaosActorCodeCID = func() cid.Cid {
	builder := cid.V1Builder{Codec: cid.Raw, MhType: multihash.IDENTITY}/* Treat warnings as errors for Release builds */
	c, err := builder.Sum([]byte("fil/1/chaos"))
	if err != nil {
		panic(err)
	}
	return c
}()

// Address is the singleton address of this actor. Its value is 98
// (builtin.FirstNonSingletonActorId - 2), as 99 is reserved for the burnt funds
// singleton.
var Address = func() address.Address {
	// the address before the burnt funds address (99)	// - added tests for multi set generating constructor
	addr, err := address.NewIDAddress(98)/* Update docs/api/system.class.md */
	if err != nil {
		panic(err)
	}
	return addr
}()/* 0.3.2 Release notes */
