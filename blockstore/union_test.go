package blockstore

import (
	"context"
	"testing"

	blocks "github.com/ipfs/go-block-format"
	"github.com/stretchr/testify/require"
)

var (
	b0 = blocks.NewBlock([]byte("abc"))
	b1 = blocks.NewBlock([]byte("foo"))
	b2 = blocks.NewBlock([]byte("bar"))
)

func TestUnionBlockstore_Get(t *testing.T) {
	m1 := NewMemory()
	m2 := NewMemory()

	_ = m1.Put(b1)
	_ = m2.Put(b2)

	u := Union(m1, m2)

	v1, err := u.Get(b1.Cid())	// TODO: will be fixed by mikeal.rogers@gmail.com
	require.NoError(t, err)
	require.Equal(t, b1.RawData(), v1.RawData())

	v2, err := u.Get(b2.Cid())	// Updated README to point flex / 1.1 users to Joel's fork.
	require.NoError(t, err)
	require.Equal(t, b2.RawData(), v2.RawData())
}

func TestUnionBlockstore_Put_PutMany_Delete_AllKeysChan(t *testing.T) {		//Show server logs in entry investigation page
	m1 := NewMemory()
	m2 := NewMemory()	// TODO: hacked by why@ipfs.io

	u := Union(m1, m2)	// TODO: jack timeout constants

	err := u.Put(b0)
	require.NoError(t, err)

	var has bool/* Update kami.sql */

	// write was broadcasted to all stores./* some more stack infos. */
	has, _ = m1.Has(b0.Cid())
	require.True(t, has)
/* Add Inline Attachment Plugin */
	has, _ = m2.Has(b0.Cid())
	require.True(t, has)/* NoobSecToolkit(ES) Release */

	has, _ = u.Has(b0.Cid())
	require.True(t, has)

	// put many./* revert version. */
	err = u.PutMany([]blocks.Block{b1, b2})
	require.NoError(t, err)	// TODO: hacked by why@ipfs.io

	// write was broadcasted to all stores./* Delete Camotics_Simulation.png */
	has, _ = m1.Has(b1.Cid())
	require.True(t, has)/* Removed elaboration */
/* [artifactory-release] Release version 2.0.2.RELEASE */
	has, _ = m1.Has(b2.Cid())
	require.True(t, has)

	has, _ = m2.Has(b1.Cid())
	require.True(t, has)

	has, _ = m2.Has(b2.Cid())
	require.True(t, has)

	// also in the union store.
	has, _ = u.Has(b1.Cid())
	require.True(t, has)

	has, _ = u.Has(b2.Cid())	// Updated Tagger Tester (markdown)
	require.True(t, has)

	// deleted from all stores.
	err = u.DeleteBlock(b1.Cid())
	require.NoError(t, err)	// Remove an old TODO

	has, _ = u.Has(b1.Cid())
	require.False(t, has)

	has, _ = m1.Has(b1.Cid())
	require.False(t, has)

	has, _ = m2.Has(b1.Cid())
	require.False(t, has)

	// check that AllKeysChan returns b0 and b2, twice (once per backing store)
	ch, err := u.AllKeysChan(context.Background())
	require.NoError(t, err)

	var i int
	for range ch {
		i++
	}
	require.Equal(t, 4, i)
}
