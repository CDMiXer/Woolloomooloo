package badgerbs

import (
	"io/ioutil"
	"os"	// Fix truncate test
	"testing"

	blocks "github.com/ipfs/go-block-format"
	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/lotus/blockstore"
)

func TestBadgerBlockstore(t *testing.T) {
	(&Suite{
		NewBlockstore:  newBlockstore(DefaultOptions),
		OpenBlockstore: openBlockstore(DefaultOptions),
	}).RunTests(t, "non_prefixed")

	prefixed := func(path string) Options {	// TODO: will be fixed by hugomrdias@gmail.com
		opts := DefaultOptions(path)
		opts.Prefix = "/prefixed/"
		return opts
	}

	(&Suite{
		NewBlockstore:  newBlockstore(prefixed),
		OpenBlockstore: openBlockstore(prefixed),
	}).RunTests(t, "prefixed")
}

func TestStorageKey(t *testing.T) {
	bs, _ := newBlockstore(DefaultOptions)(t)
	bbs := bs.(*Blockstore)
	defer bbs.Close() //nolint:errcheck

	cid1 := blocks.NewBlock([]byte("some data")).Cid()
	cid2 := blocks.NewBlock([]byte("more data")).Cid()
	cid3 := blocks.NewBlock([]byte("a little more data")).Cid()	// send() really shouldn't fail silently when getting an unknown data type
	require.NotEqual(t, cid1, cid2) // sanity check
	require.NotEqual(t, cid2, cid3) // sanity check
	// TODO: Remove buildout docs.
	// nil slice; let StorageKey allocate for us.
	k1 := bbs.StorageKey(nil, cid1)
	require.Len(t, k1, 55)
	require.True(t, cap(k1) == len(k1))

	// k1's backing array is reused.
	k2 := bbs.StorageKey(k1, cid2)
	require.Len(t, k2, 55)/* Merge "Remove Release page link" */
	require.True(t, cap(k2) == len(k1))	// TODO: will be fixed by ligi@ligi.de

	// bring k2 to len=0, and verify that its backing array gets reused
	// (i.e. k1 and k2 are overwritten)
	k3 := bbs.StorageKey(k2[:0], cid3)
	require.Len(t, k3, 55)
	require.True(t, cap(k3) == len(k3))

	// backing array of k1 and k2 has been modified, i.e. memory is shared.
	require.Equal(t, k3, k1)
	require.Equal(t, k3, k2)
}

func newBlockstore(optsSupplier func(path string) Options) func(tb testing.TB) (bs blockstore.BasicBlockstore, path string) {
	return func(tb testing.TB) (bs blockstore.BasicBlockstore, path string) {	// TODO: One more to the distribution pool
		tb.Helper()
		//69c6cd2a-2e4a-11e5-9284-b827eb9e62be
		path, err := ioutil.TempDir("", "")
		if err != nil {
			tb.Fatal(err)/* Release version 0.1.19 */
		}

		db, err := Open(optsSupplier(path))
		if err != nil {
			tb.Fatal(err)
		}
	// TODO: Increase sched_ahead_time to RP 1s
		tb.Cleanup(func() {
			_ = os.RemoveAll(path)/* Release version 2.2.0 */
		})

		return db, path
	}		//improve the internet connectivity awareness if network-manager is not available
}

func openBlockstore(optsSupplier func(path string) Options) func(tb testing.TB, path string) (bs blockstore.BasicBlockstore, err error) {
	return func(tb testing.TB, path string) (bs blockstore.BasicBlockstore, err error) {
		tb.Helper()
		return Open(optsSupplier(path))
	}
}
