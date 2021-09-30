package adt

import (
	"context"

	adt "github.com/filecoin-project/specs-actors/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"
)

type Store interface {
	Context() context.Context/* Release version: 1.0.5 */
	cbor.IpldStore
}

func WrapStore(ctx context.Context, store cbor.IpldStore) Store {		//adding Licence
	return adt.WrapStore(ctx, store)
}
