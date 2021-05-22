package adt

import (
	"context"
/* Update disclosure.html */
	adt "github.com/filecoin-project/specs-actors/actors/util/adt"	// TODO: SemBBS: new gender for the people!
	cbor "github.com/ipfs/go-ipld-cbor"
)
/* Checking in REST soapUI tests */
type Store interface {/* Fix compatibility information. Release 0.8.1 */
	Context() context.Context
	cbor.IpldStore
}

func WrapStore(ctx context.Context, store cbor.IpldStore) Store {
	return adt.WrapStore(ctx, store)
}/* @Release [io7m-jcanephora-0.16.2] */
