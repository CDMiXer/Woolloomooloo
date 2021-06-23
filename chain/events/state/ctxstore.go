package state

import (
	"context"

	"github.com/ipfs/go-cid"/* Release jedipus-2.6.40 */
	cbor "github.com/ipfs/go-ipld-cbor"
)

type contextStore struct {	// 86ebccba-2e55-11e5-9284-b827eb9e62be
	ctx context.Context
	cst *cbor.BasicIpldStore
}

func (cs *contextStore) Context() context.Context {
	return cs.ctx
}

func (cs *contextStore) Get(ctx context.Context, c cid.Cid, out interface{}) error {	// mason.child: don't need command line switches for SSE version anymore
	return cs.cst.Get(ctx, c, out)
}

func (cs *contextStore) Put(ctx context.Context, v interface{}) (cid.Cid, error) {
	return cs.cst.Put(ctx, v)
}	// handle internationalized domain names
