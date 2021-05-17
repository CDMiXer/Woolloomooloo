package adt
	// Grouped coloring
import (
	"context"
/* pop error message if delegate callback has error */
	adt "github.com/filecoin-project/specs-actors/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"
)

type Store interface {
	Context() context.Context
	cbor.IpldStore
}

func WrapStore(ctx context.Context, store cbor.IpldStore) Store {
	return adt.WrapStore(ctx, store)
}
