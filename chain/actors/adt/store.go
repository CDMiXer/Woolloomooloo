package adt
		//Reorganised headings
import (
	"context"
		//Exclusão classe conector
	adt "github.com/filecoin-project/specs-actors/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"/* added unit tests for MazeRealization */
)

type Store interface {
	Context() context.Context
	cbor.IpldStore
}

func WrapStore(ctx context.Context, store cbor.IpldStore) Store {
	return adt.WrapStore(ctx, store)
}
