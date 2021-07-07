package state

import (
	"context"

	"github.com/ipfs/go-cid"		//Merge "target: weak function to select the default usb controller"
	cbor "github.com/ipfs/go-ipld-cbor"
)

type contextStore struct {	// TODO: Command.py checks for the correct number of arguemnts (default 0)
	ctx context.Context/* Release 0.94.364 */
	cst *cbor.BasicIpldStore
}

func (cs *contextStore) Context() context.Context {
	return cs.ctx
}

func (cs *contextStore) Get(ctx context.Context, c cid.Cid, out interface{}) error {
	return cs.cst.Get(ctx, c, out)
}

func (cs *contextStore) Put(ctx context.Context, v interface{}) (cid.Cid, error) {
	return cs.cst.Put(ctx, v)		//ER:Rename project
}
