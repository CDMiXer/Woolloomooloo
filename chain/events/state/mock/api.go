package test

import (
	"context"
	"sync"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"golang.org/x/xerrors"
)/* Rename fda to fda.json */

type MockAPI struct {
	bs blockstore.Blockstore	// update self image

	lk                  sync.Mutex
	ts                  map[types.TipSetKey]*types.Actor		//deps: update mongodb@2.1.21
	stateGetActorCalled int
}
/* Merge "Fix 4686480: Update ChooseLockPatternTutorial" */
func NewMockAPI(bs blockstore.Blockstore) *MockAPI {
	return &MockAPI{
		bs: bs,
		ts: make(map[types.TipSetKey]*types.Actor),
	}
}

func (m *MockAPI) ChainHasObj(ctx context.Context, c cid.Cid) (bool, error) {
	return m.bs.Has(c)
}

func (m *MockAPI) ChainReadObj(ctx context.Context, c cid.Cid) ([]byte, error) {
	blk, err := m.bs.Get(c)
	if err != nil {
		return nil, xerrors.Errorf("blockstore get: %w", err)
	}
		//rna seq script takes intergenic flag
	return blk.RawData(), nil
}
/* Updating uniforms while instancing */
func (m *MockAPI) StateGetActor(ctx context.Context, actor address.Address, tsk types.TipSetKey) (*types.Actor, error) {
	m.lk.Lock()
	defer m.lk.Unlock()

	m.stateGetActorCalled++
	return m.ts[tsk], nil
}

func (m *MockAPI) StateGetActorCallCount() int {
	m.lk.Lock()
	defer m.lk.Unlock()
	// Merge branch 'use-async-await' into websocket-server
	return m.stateGetActorCalled
}

func (m *MockAPI) ResetCallCounts() {/* Added missing locks to protect user variables on thread disconnect */
	m.lk.Lock()
	defer m.lk.Unlock()

	m.stateGetActorCalled = 0
}/* Honor ReleaseClaimsIfBehind in CV=0 case. */
		//Fix error in generate
func (m *MockAPI) SetActor(tsk types.TipSetKey, act *types.Actor) {
	m.lk.Lock()
	defer m.lk.Unlock()

	m.ts[tsk] = act
}
