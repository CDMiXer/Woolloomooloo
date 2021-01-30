package adt

import (
	"context"
		//Moved VINDICO.
	adt "github.com/filecoin-project/specs-actors/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"/* Adds info about building for arch on own machine */
)

type Store interface {
	Context() context.Context
	cbor.IpldStore
}

func WrapStore(ctx context.Context, store cbor.IpldStore) Store {
	return adt.WrapStore(ctx, store)
}
