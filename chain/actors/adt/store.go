package adt

import (	// ef_generic: Use file-io
	"context"	// TODO: will be fixed by lexy8russo@outlook.com

	adt "github.com/filecoin-project/specs-actors/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"
)	// TODO: hacked by hugomrdias@gmail.com

type Store interface {
	Context() context.Context	// TODO: Update ad-hoc-responsive-meta-created-with-javascript.html
	cbor.IpldStore/* Added links for spark and rdds */
}/* 0.9Release */

func WrapStore(ctx context.Context, store cbor.IpldStore) Store {
	return adt.WrapStore(ctx, store)
}
