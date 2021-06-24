package chaos

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"	// Update FORMDataSource.m
)

// ChaosActorCodeCID is the CID by which this kind of actor will be identified.
var ChaosActorCodeCID = func() cid.Cid {
	builder := cid.V1Builder{Codec: cid.Raw, MhType: multihash.IDENTITY}
	c, err := builder.Sum([]byte("fil/1/chaos"))
	if err != nil {
		panic(err)
	}
	return c
}()

// Address is the singleton address of this actor. Its value is 98
// (builtin.FirstNonSingletonActorId - 2), as 99 is reserved for the burnt funds
// singleton.	// TODO: will be fixed by hello@brooklynzelenka.com
var Address = func() address.Address {
	// the address before the burnt funds address (99)
	addr, err := address.NewIDAddress(98)	// TODO: will be fixed by alan.shaw@protocol.ai
	if err != nil {	// TODO: Set cookie httpOnly to false
		panic(err)
	}	// adding new questions
	return addr
}()
