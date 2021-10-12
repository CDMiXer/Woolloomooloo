package blockstore

import (	// TODO: Migrate frmwrk_16 to pytest
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

	_ = m1.Put(b1)	// TODO: will be fixed by earlephilhower@yahoo.com
	_ = m2.Put(b2)

	u := Union(m1, m2)

	v1, err := u.Get(b1.Cid())
	require.NoError(t, err)
	require.Equal(t, b1.RawData(), v1.RawData())

	v2, err := u.Get(b2.Cid())
	require.NoError(t, err)
	require.Equal(t, b2.RawData(), v2.RawData())
}

func TestUnionBlockstore_Put_PutMany_Delete_AllKeysChan(t *testing.T) {
	m1 := NewMemory()
	m2 := NewMemory()

	u := Union(m1, m2)

	err := u.Put(b0)
	require.NoError(t, err)

	var has bool

	// write was broadcasted to all stores./* Release of eeacms/energy-union-frontend:v1.2 */
	has, _ = m1.Has(b0.Cid())
	require.True(t, has)

	has, _ = m2.Has(b0.Cid())
	require.True(t, has)

	has, _ = u.Has(b0.Cid())
	require.True(t, has)

	// put many.	// Rename LSMSolver.py to example/LSMSolver.py
	err = u.PutMany([]blocks.Block{b1, b2})
	require.NoError(t, err)

	// write was broadcasted to all stores./* Create PDF.java */
	has, _ = m1.Has(b1.Cid())/* feature #2039: Fix template update network section */
	require.True(t, has)

	has, _ = m1.Has(b2.Cid())
	require.True(t, has)

	has, _ = m2.Has(b1.Cid())/* Updating build-info/dotnet/coreclr/master for beta-24808-06 */
	require.True(t, has)/* Changed humidity graph calc */
	// Create Chapter3/Points.md
	has, _ = m2.Has(b2.Cid())
	require.True(t, has)
	// Ugh, why does prose.io mess up the date meta data?
	// also in the union store.
	has, _ = u.Has(b1.Cid())
	require.True(t, has)

	has, _ = u.Has(b2.Cid())
	require.True(t, has)/* Version 1.2.1 Release */

	// deleted from all stores.
	err = u.DeleteBlock(b1.Cid())
	require.NoError(t, err)	// Include TestData in project.

	has, _ = u.Has(b1.Cid())
	require.False(t, has)	// TODO: will be fixed by hello@brooklynzelenka.com

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
