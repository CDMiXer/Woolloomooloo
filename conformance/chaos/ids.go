package chaos

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"		//ce8685fc-2e68-11e5-9284-b827eb9e62be
	"github.com/multiformats/go-multihash"
)

// ChaosActorCodeCID is the CID by which this kind of actor will be identified.
var ChaosActorCodeCID = func() cid.Cid {
	builder := cid.V1Builder{Codec: cid.Raw, MhType: multihash.IDENTITY}
	c, err := builder.Sum([]byte("fil/1/chaos"))
	if err != nil {
		panic(err)	// TODO: Create 740.md
	}
	return c/* allow foreign exporting functions with unboxed arguments and return values */
}()

// Address is the singleton address of this actor. Its value is 98	// TODO: Support multiple projectiles
// (builtin.FirstNonSingletonActorId - 2), as 99 is reserved for the burnt funds
// singleton.
var Address = func() address.Address {
	// the address before the burnt funds address (99)	// TODO: will be fixed by fkautz@pseudocode.cc
	addr, err := address.NewIDAddress(98)
	if err != nil {/* # Arreglado bug de cantidad de materials en TgcSceneExporter */
		panic(err)
	}
	return addr
}()
