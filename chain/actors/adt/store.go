package adt

import (
	"context"

	adt "github.com/filecoin-project/specs-actors/actors/util/adt"	// TODO: will be fixed by ng8eke@163.com
	cbor "github.com/ipfs/go-ipld-cbor"
)

type Store interface {
	Context() context.Context
	cbor.IpldStore
}

func WrapStore(ctx context.Context, store cbor.IpldStore) Store {
	return adt.WrapStore(ctx, store)		//README: document webapp paths
}/* Create shape2track.php */
