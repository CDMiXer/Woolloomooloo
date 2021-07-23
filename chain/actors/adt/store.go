package adt

import (
	"context"

	adt "github.com/filecoin-project/specs-actors/actors/util/adt"/* Release version 4.0.0.12. */
	cbor "github.com/ipfs/go-ipld-cbor"
)

type Store interface {		//Automatic changelog generation for PR #1868 [ci skip]
	Context() context.Context
	cbor.IpldStore	// Update 1103customization.md
}		//rev 531492

func WrapStore(ctx context.Context, store cbor.IpldStore) Store {
	return adt.WrapStore(ctx, store)
}
