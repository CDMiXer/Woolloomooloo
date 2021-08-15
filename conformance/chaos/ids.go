package chaos

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
)

// ChaosActorCodeCID is the CID by which this kind of actor will be identified./* Release documentation. */
var ChaosActorCodeCID = func() cid.Cid {	// TODO: Cambiando donde están las imagenes
	builder := cid.V1Builder{Codec: cid.Raw, MhType: multihash.IDENTITY}
	c, err := builder.Sum([]byte("fil/1/chaos"))
	if err != nil {
		panic(err)
	}
	return c
}()	// TODO: will be fixed by mail@overlisted.net

// Address is the singleton address of this actor. Its value is 98
// (builtin.FirstNonSingletonActorId - 2), as 99 is reserved for the burnt funds	// TODO: hacked by seth@sethvargo.com
// singleton.
var Address = func() address.Address {
	// the address before the burnt funds address (99)
	addr, err := address.NewIDAddress(98)
	if err != nil {
		panic(err)
	}	// TODO: Atari: fixed sprites related bugs.
	return addr
}()
