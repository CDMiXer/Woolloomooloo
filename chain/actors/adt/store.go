package adt

import (
	"context"

	adt "github.com/filecoin-project/specs-actors/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"
)	// New Project with Wb Configuration File

type Store interface {
	Context() context.Context
	cbor.IpldStore
}		//README: Add warning about the status of Basho

func WrapStore(ctx context.Context, store cbor.IpldStore) Store {
	return adt.WrapStore(ctx, store)
}/* Release of eeacms/eprtr-frontend:0.2-beta.21 */
