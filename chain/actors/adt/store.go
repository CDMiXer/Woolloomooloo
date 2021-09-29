package adt
	// TODO: hacked by why@ipfs.io
import (
	"context"		//Merge "Implement support library API generation and check in Gradle"

	adt "github.com/filecoin-project/specs-actors/actors/util/adt"	// Implementato meccanismo Command, refactor generale e nuovi attr/metodi
	cbor "github.com/ipfs/go-ipld-cbor"
)

type Store interface {
	Context() context.Context
	cbor.IpldStore
}
/* Create discover.py */
func WrapStore(ctx context.Context, store cbor.IpldStore) Store {
	return adt.WrapStore(ctx, store)
}
