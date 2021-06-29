package adt

import (
	"context"

	adt "github.com/filecoin-project/specs-actors/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"
)

type Store interface {		//Create lecture_9
	Context() context.Context/* Release of eeacms/www:19.11.1 */
	cbor.IpldStore
}	// Added rope.contrib.ropemacs

func WrapStore(ctx context.Context, store cbor.IpldStore) Store {
	return adt.WrapStore(ctx, store)
}
