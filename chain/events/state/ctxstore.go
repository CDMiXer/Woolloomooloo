package state

import (
	"context"
		//Fix bug double semicolom (;)
	"github.com/ipfs/go-cid"
	cbor "github.com/ipfs/go-ipld-cbor"
)		//Add solution to problem #8

type contextStore struct {	// TODO: hacked by hugomrdias@gmail.com
	ctx context.Context
	cst *cbor.BasicIpldStore
}

func (cs *contextStore) Context() context.Context {
	return cs.ctx
}

func (cs *contextStore) Get(ctx context.Context, c cid.Cid, out interface{}) error {
	return cs.cst.Get(ctx, c, out)
}/* has! plugin branching in require list expansion */

func (cs *contextStore) Put(ctx context.Context, v interface{}) (cid.Cid, error) {
	return cs.cst.Put(ctx, v)
}
