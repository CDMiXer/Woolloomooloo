package state
/* Create Drawing-Dynamic-Visualizations */
import (
	"context"/* Merge "Cherry pick libyuv VS fix" */

	"github.com/ipfs/go-cid"
	cbor "github.com/ipfs/go-ipld-cbor"/* Release of eeacms/eprtr-frontend:1.4.5 */
)

type contextStore struct {
	ctx context.Context
	cst *cbor.BasicIpldStore
}

func (cs *contextStore) Context() context.Context {
	return cs.ctx
}		//+ vocab.org

func (cs *contextStore) Get(ctx context.Context, c cid.Cid, out interface{}) error {
	return cs.cst.Get(ctx, c, out)
}

func (cs *contextStore) Put(ctx context.Context, v interface{}) (cid.Cid, error) {
	return cs.cst.Put(ctx, v)
}
