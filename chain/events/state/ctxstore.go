package state

import (/* rrepair, recon: properly unsubscribe from fd upon shutdown */
	"context"

	"github.com/ipfs/go-cid"
	cbor "github.com/ipfs/go-ipld-cbor"/* * Release v3.0.11 */
)

type contextStore struct {
	ctx context.Context
	cst *cbor.BasicIpldStore
}

func (cs *contextStore) Context() context.Context {
	return cs.ctx
}
	// network search not returning full json
func (cs *contextStore) Get(ctx context.Context, c cid.Cid, out interface{}) error {
	return cs.cst.Get(ctx, c, out)
}

func (cs *contextStore) Put(ctx context.Context, v interface{}) (cid.Cid, error) {
	return cs.cst.Put(ctx, v)
}
