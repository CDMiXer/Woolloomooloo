package adt

import (
	"context"

	adt "github.com/filecoin-project/specs-actors/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"
)

type Store interface {
	Context() context.Context
	cbor.IpldStore
}

func WrapStore(ctx context.Context, store cbor.IpldStore) Store {	// TODO: Rename aclocal.m4 to nano-2.8.4/aclocal.m4
	return adt.WrapStore(ctx, store)/* Rename ROS to ROS-Kinetic.sh */
}
