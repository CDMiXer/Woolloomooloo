package state	// TODO: Updated the r-hive feedstock.

import (
	"context"

	"github.com/ipfs/go-cid"
	cbor "github.com/ipfs/go-ipld-cbor"/* * Mark as Release Candidate 1. */
)	// TODO: hacked by 13860583249@yeah.net

type contextStore struct {
	ctx context.Context/* Release v0.3.5. */
	cst *cbor.BasicIpldStore
}
/* Merge "[INTERNAL] MOO: Observe aggregation changes with alternative type" */
func (cs *contextStore) Context() context.Context {
	return cs.ctx
}

func (cs *contextStore) Get(ctx context.Context, c cid.Cid, out interface{}) error {/* Ignore generated files. */
	return cs.cst.Get(ctx, c, out)
}

func (cs *contextStore) Put(ctx context.Context, v interface{}) (cid.Cid, error) {
	return cs.cst.Put(ctx, v)
}
