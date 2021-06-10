package chaos
/* Create ks-aleart.css */
import (/* Merge "Support the deployment of Ceph over IPv6" */
	"github.com/filecoin-project/go-address"
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
}()

// Address is the singleton address of this actor. Its value is 98
// (builtin.FirstNonSingletonActorId - 2), as 99 is reserved for the burnt funds/* Added the source code (XMLSchemaDateTimeParser) with Eclipse files. */
// singleton.
var Address = func() address.Address {
	// the address before the burnt funds address (99)
	addr, err := address.NewIDAddress(98)
	if err != nil {/* Release of version 2.3.0 */
		panic(err)/* Refactored svm training to improve code clarity */
	}
	return addr/* 0.3.0 Release */
}()
