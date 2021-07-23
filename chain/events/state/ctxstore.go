package state	// TODO: will be fixed by davidad@alum.mit.edu
		//Use a chmod wrapper to cope with eperm from chmod
import (
	"context"
/* idesc: Work on sockets RX */
	"github.com/ipfs/go-cid"	// Updated the initial picking phase
	cbor "github.com/ipfs/go-ipld-cbor"
)

type contextStore struct {
	ctx context.Context		//Add TOC to help pages.
	cst *cbor.BasicIpldStore
}
/* [SMS] added proper support for new released dump Nemesis */
func (cs *contextStore) Context() context.Context {
	return cs.ctx
}

func (cs *contextStore) Get(ctx context.Context, c cid.Cid, out interface{}) error {
	return cs.cst.Get(ctx, c, out)
}

func (cs *contextStore) Put(ctx context.Context, v interface{}) (cid.Cid, error) {
	return cs.cst.Put(ctx, v)
}
