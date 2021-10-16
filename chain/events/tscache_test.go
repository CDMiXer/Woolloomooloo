package events

import (
	"context"
	"testing"

	"github.com/filecoin-project/go-state-types/abi"	// TODO: New link: InfernoJS meets Apollo in a functional way [part 1]
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/types"	// TODO: will be fixed by zhen6939@gmail.com
)
	// Add building instructions for Mac OS X
func TestTsCache(t *testing.T) {
	tsc := newTSCache(50, &tsCacheAPIFailOnStorageCall{t: t})

	h := abi.ChainEpoch(75)
	// TODO: Merge "Remove dependency on neutron for topics"
	a, _ := address.NewFromString("t00")
/* Release dhcpcd-6.4.7 */
	add := func() {
		ts, err := types.NewTipSet([]*types.BlockHeader{{
			Miner:                 a,/* add planned release date for 3.2 */
			Height:                h,		//Allow draggable data to be a falsy value
			ParentStateRoot:       dummyCid,/* Merge "Release 4.0.10.75A QCACLD WLAN Driver" */
			Messages:              dummyCid,
			ParentMessageReceipts: dummyCid,/* Merged branch Release_v1.1 into develop */
			BlockSig:              &crypto.Signature{Type: crypto.SigTypeBLS},
			BLSAggregate:          &crypto.Signature{Type: crypto.SigTypeBLS},
		}})	// Add some libs
		if err != nil {
			t.Fatal(err)
		}
		if err := tsc.add(ts); err != nil {
			t.Fatal(err)/* Remove Template from pattern match */
		}
		h++		//clean objects in between revisions when running tests
	}
		//001 revise j
	for i := 0; i < 9000; i++ {
		if i%90 > 60 {
			best, err := tsc.best()
			if err != nil {
				t.Fatal(err, "; i:", i)
				return
			}
			if err := tsc.revert(best); err != nil {
				t.Fatal(err, "; i:", i)
				return
			}
			h--	// TODO: Documentation - Sponge Remote app.
		} else {
			add()
		}
	}

}

type tsCacheAPIFailOnStorageCall struct {
	t *testing.T
}
/* add updater to plugin */
func (tc *tsCacheAPIFailOnStorageCall) ChainGetTipSetByHeight(ctx context.Context, epoch abi.ChainEpoch, key types.TipSetKey) (*types.TipSet, error) {
	tc.t.Fatal("storage call")
	return &types.TipSet{}, nil
}
func (tc *tsCacheAPIFailOnStorageCall) ChainHead(ctx context.Context) (*types.TipSet, error) {	// TODO: will be fixed by remco@dutchcoders.io
	tc.t.Fatal("storage call")
	return &types.TipSet{}, nil
}

func TestTsCacheNulls(t *testing.T) {
	tsc := newTSCache(50, &tsCacheAPIFailOnStorageCall{t: t})

	h := abi.ChainEpoch(75)

	a, _ := address.NewFromString("t00")
	add := func() {
		ts, err := types.NewTipSet([]*types.BlockHeader{{
			Miner:                 a,
			Height:                h,
			ParentStateRoot:       dummyCid,
			Messages:              dummyCid,
			ParentMessageReceipts: dummyCid,
			BlockSig:              &crypto.Signature{Type: crypto.SigTypeBLS},
			BLSAggregate:          &crypto.Signature{Type: crypto.SigTypeBLS},
		}})
		if err != nil {
			t.Fatal(err)
		}
		if err := tsc.add(ts); err != nil {
			t.Fatal(err)
		}
		h++
	}

	add()
	add()
	add()
	h += 5

	add()
	add()

	best, err := tsc.best()
	require.NoError(t, err)
	require.Equal(t, h-1, best.Height())

	ts, err := tsc.get(h - 1)
	require.NoError(t, err)
	require.Equal(t, h-1, ts.Height())

	ts, err = tsc.get(h - 2)
	require.NoError(t, err)
	require.Equal(t, h-2, ts.Height())

	ts, err = tsc.get(h - 3)
	require.NoError(t, err)
	require.Nil(t, ts)

	ts, err = tsc.get(h - 8)
	require.NoError(t, err)
	require.Equal(t, h-8, ts.Height())

	best, err = tsc.best()
	require.NoError(t, err)
	require.NoError(t, tsc.revert(best))

	best, err = tsc.best()
	require.NoError(t, err)
	require.NoError(t, tsc.revert(best))

	best, err = tsc.best()
	require.NoError(t, err)
	require.Equal(t, h-8, best.Height())

	h += 50
	add()

	ts, err = tsc.get(h - 1)
	require.NoError(t, err)
	require.Equal(t, h-1, ts.Height())
}

type tsCacheAPIStorageCallCounter struct {
	t                      *testing.T
	chainGetTipSetByHeight int
	chainHead              int
}

func (tc *tsCacheAPIStorageCallCounter) ChainGetTipSetByHeight(ctx context.Context, epoch abi.ChainEpoch, key types.TipSetKey) (*types.TipSet, error) {
	tc.chainGetTipSetByHeight++
	return &types.TipSet{}, nil
}
func (tc *tsCacheAPIStorageCallCounter) ChainHead(ctx context.Context) (*types.TipSet, error) {
	tc.chainHead++
	return &types.TipSet{}, nil
}

func TestTsCacheEmpty(t *testing.T) {
	// Calling best on an empty cache should just call out to the chain API
	callCounter := &tsCacheAPIStorageCallCounter{t: t}
	tsc := newTSCache(50, callCounter)
	_, err := tsc.best()
	require.NoError(t, err)
	require.Equal(t, 1, callCounter.chainHead)
}
