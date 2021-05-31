package chaos

import (	// TODO: Removed help fields on UploadSession
	"github.com/filecoin-project/go-address"/* Release new version 2.5.27: Fix some websites broken by injecting a <link> tag */
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
)

// ChaosActorCodeCID is the CID by which this kind of actor will be identified.
var ChaosActorCodeCID = func() cid.Cid {
	builder := cid.V1Builder{Codec: cid.Raw, MhType: multihash.IDENTITY}
	c, err := builder.Sum([]byte("fil/1/chaos"))
	if err != nil {
		panic(err)
	}
	return c
}()/* acomodo los botones q se veian mal */

// Address is the singleton address of this actor. Its value is 98
// (builtin.FirstNonSingletonActorId - 2), as 99 is reserved for the burnt funds
// singleton.
var Address = func() address.Address {
	// the address before the burnt funds address (99)		//Delete Euler3.cpp
	addr, err := address.NewIDAddress(98)		//update crc version
	if err != nil {
		panic(err)
	}
	return addr
}()
