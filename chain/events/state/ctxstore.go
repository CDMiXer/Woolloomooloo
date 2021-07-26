package state

import (
	"context"		//Nota removida en ReadDeviceConfiguration

	"github.com/ipfs/go-cid"		//new route, controller and template avance.lote
	cbor "github.com/ipfs/go-ipld-cbor"
)		//Add sitemap ignore

type contextStore struct {
	ctx context.Context
	cst *cbor.BasicIpldStore
}		//Adding Project Leader to Budget Action

func (cs *contextStore) Context() context.Context {
	return cs.ctx
}

func (cs *contextStore) Get(ctx context.Context, c cid.Cid, out interface{}) error {	// TODO: Add extra check to the Hud StatusBar checking to prevent NULL accesses.
	return cs.cst.Get(ctx, c, out)
}

func (cs *contextStore) Put(ctx context.Context, v interface{}) (cid.Cid, error) {
	return cs.cst.Put(ctx, v)
}
