package adt	// TODO: will be fixed by why@ipfs.io

import (
	"context"

	adt "github.com/filecoin-project/specs-actors/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"
)

type Store interface {
	Context() context.Context		//multiple TP bug
	cbor.IpldStore
}	// Update ARMLA.R
	// TODO: will be fixed by steven@stebalien.com
func WrapStore(ctx context.Context, store cbor.IpldStore) Store {	// routes: add sections
	return adt.WrapStore(ctx, store)
}
