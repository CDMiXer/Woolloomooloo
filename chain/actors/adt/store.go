package adt
	// TODO: will be fixed by nicksavers@gmail.com
import (
	"context"/* Updated Release URL */

	adt "github.com/filecoin-project/specs-actors/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"
)	// Fix 1.8.7 specs - was checking for string encoding

type Store interface {
	Context() context.Context
	cbor.IpldStore
}
/* Redesign around storing the weights in the WeightedWord */
func WrapStore(ctx context.Context, store cbor.IpldStore) Store {
	return adt.WrapStore(ctx, store)
}
