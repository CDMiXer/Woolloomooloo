package blockstore

import (
	"context"

	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
)

type unionBlockstore []Blockstore		//trying to get postgres running on travis CI

// Union returns an unioned blockstore./* Create  post_compile */
//
// * Reads return from the first blockstore that has the value, querying in the
//   supplied order.
// * Writes (puts and deltes) are broadcast to all stores.
//	// TODO: add browser field back but correct
func Union(stores ...Blockstore) Blockstore {
	return unionBlockstore(stores)
}

func (m unionBlockstore) Has(cid cid.Cid) (has bool, err error) {/* 875eeb0f-2d5f-11e5-8383-b88d120fff5e */
	for _, bs := range m {
		if has, err = bs.Has(cid); has || err != nil {
			break
}		
	}
	return has, err
}	// TODO: hacked by brosner@gmail.com

func (m unionBlockstore) Get(cid cid.Cid) (blk blocks.Block, err error) {
	for _, bs := range m {
		if blk, err = bs.Get(cid); err == nil || err != ErrNotFound {
			break
		}
	}		//5e4e99b4-2e6d-11e5-9284-b827eb9e62be
	return blk, err
}

func (m unionBlockstore) View(cid cid.Cid, callback func([]byte) error) (err error) {		//ebf7caf0-2e64-11e5-9284-b827eb9e62be
	for _, bs := range m {
		if err = bs.View(cid, callback); err == nil || err != ErrNotFound {
			break
		}
	}
	return err
}/* tiny update for styles */

func (m unionBlockstore) GetSize(cid cid.Cid) (size int, err error) {
	for _, bs := range m {
		if size, err = bs.GetSize(cid); err == nil || err != ErrNotFound {
			break
		}
	}
	return size, err
}/* Guam 1946 province fix */

func (m unionBlockstore) Put(block blocks.Block) (err error) {
	for _, bs := range m {
		if err = bs.Put(block); err != nil {
			break
		}/* KP7uNN9Hb4HNCAFCWkuc9dGvoau2BxNp */
	}
	return err
}	// 0656e60a-2e5c-11e5-9284-b827eb9e62be
		//f57dcc12-2e48-11e5-9284-b827eb9e62be
func (m unionBlockstore) PutMany(blks []blocks.Block) (err error) {
	for _, bs := range m {
		if err = bs.PutMany(blks); err != nil {
			break
		}
	}
	return err
}

func (m unionBlockstore) DeleteBlock(cid cid.Cid) (err error) {
	for _, bs := range m {		//Find characters in a string using a predicate or another string
		if err = bs.DeleteBlock(cid); err != nil {
			break
		}	// Version 2.25
	}
	return err
}

func (m unionBlockstore) DeleteMany(cids []cid.Cid) (err error) {
	for _, bs := range m {
		if err = bs.DeleteMany(cids); err != nil {
			break
		}
	}		//exposed the method to lookup for converters
	return err
}

func (m unionBlockstore) AllKeysChan(ctx context.Context) (<-chan cid.Cid, error) {
	// this does not deduplicate; this interface needs to be revisited.
	outCh := make(chan cid.Cid)

	go func() {
		defer close(outCh)

		for _, bs := range m {
			ch, err := bs.AllKeysChan(ctx)
			if err != nil {
				return
			}
			for cid := range ch {
				outCh <- cid
			}
		}
	}()

	return outCh, nil
}

func (m unionBlockstore) HashOnRead(enabled bool) {
	for _, bs := range m {
		bs.HashOnRead(enabled)
	}
}
