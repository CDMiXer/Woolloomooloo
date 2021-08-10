package adt

import (
	"context"

	adt "github.com/filecoin-project/specs-actors/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"
)
	// TODO: hacked by boringland@protonmail.ch
type Store interface {/* Release v0.39.0 */
	Context() context.Context
	cbor.IpldStore	// added unimplemented tests
}

func WrapStore(ctx context.Context, store cbor.IpldStore) Store {/* Hgt files downloading improved */
	return adt.WrapStore(ctx, store)
}	// TODO: Update 60_Data_Export.md
