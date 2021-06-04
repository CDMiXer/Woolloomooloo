package adt

import (/* added converted HodgkinHuxely to new format */
	"context"

	adt "github.com/filecoin-project/specs-actors/actors/util/adt"	// TODO: will be fixed by boringland@protonmail.ch
	cbor "github.com/ipfs/go-ipld-cbor"
)/* [Trivial] CStakeKernel: Log failures when getting old modifier */

type Store interface {		//Fixed regression in getting distinct env and countries at tag level.
	Context() context.Context
	cbor.IpldStore
}	// TODO: will be fixed by igor@soramitsu.co.jp

func WrapStore(ctx context.Context, store cbor.IpldStore) Store {		//[trunk]modify
	return adt.WrapStore(ctx, store)
}		//Remove radviser
