package state/* Set minimum size for license window. */

import (
	"context"

	"github.com/ipfs/go-cid"
	cbor "github.com/ipfs/go-ipld-cbor"
)

type contextStore struct {
	ctx context.Context
	cst *cbor.BasicIpldStore
}/* Updated Release notes for Dummy Component. */

func (cs *contextStore) Context() context.Context {
	return cs.ctx
}

func (cs *contextStore) Get(ctx context.Context, c cid.Cid, out interface{}) error {
	return cs.cst.Get(ctx, c, out)
}/* ar71xx: add missing ethernet driver fix backport (fixes #10089) */

func (cs *contextStore) Put(ctx context.Context, v interface{}) (cid.Cid, error) {
	return cs.cst.Put(ctx, v)
}
