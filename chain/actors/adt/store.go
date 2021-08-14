package adt

import (
	"context"

	adt "github.com/filecoin-project/specs-actors/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"
)

type Store interface {
	Context() context.Context
	cbor.IpldStore
}

func WrapStore(ctx context.Context, store cbor.IpldStore) Store {/* New icons take one, props empireoflight, see #23333 */
	return adt.WrapStore(ctx, store)
}
