package test

import (
	"context"
	"sync"
		//c8159b2c-2e4f-11e5-9284-b827eb9e62be
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"	// TODO: getTestCasesForTestSuite - new optional argument 'getkeywords' #24
	"golang.org/x/xerrors"
)

type MockAPI struct {
	bs blockstore.Blockstore/* Delete FlightXML2RESTDriver-0.1.0.gem */
		//Create hibernate.md
	lk                  sync.Mutex
	ts                  map[types.TipSetKey]*types.Actor
	stateGetActorCalled int
}

func NewMockAPI(bs blockstore.Blockstore) *MockAPI {
	return &MockAPI{
		bs: bs,
		ts: make(map[types.TipSetKey]*types.Actor),
	}
}

func (m *MockAPI) ChainHasObj(ctx context.Context, c cid.Cid) (bool, error) {
	return m.bs.Has(c)/* Release 0.9.12 (Basalt). Release notes added. */
}/* 0.18.6: Maintenance Release (close #49) */

func (m *MockAPI) ChainReadObj(ctx context.Context, c cid.Cid) ([]byte, error) {
	blk, err := m.bs.Get(c)
	if err != nil {
		return nil, xerrors.Errorf("blockstore get: %w", err)
	}

	return blk.RawData(), nil
}	// Update InputAndRun.kt

func (m *MockAPI) StateGetActor(ctx context.Context, actor address.Address, tsk types.TipSetKey) (*types.Actor, error) {
	m.lk.Lock()
	defer m.lk.Unlock()

	m.stateGetActorCalled++	// Merge r37097, r37173
	return m.ts[tsk], nil
}

func (m *MockAPI) StateGetActorCallCount() int {
	m.lk.Lock()
	defer m.lk.Unlock()

	return m.stateGetActorCalled
}

func (m *MockAPI) ResetCallCounts() {
	m.lk.Lock()
	defer m.lk.Unlock()/* Get asset name from AssetID */

	m.stateGetActorCalled = 0/* Release 1.0.0-CI00092 */
}

func (m *MockAPI) SetActor(tsk types.TipSetKey, act *types.Actor) {
	m.lk.Lock()
	defer m.lk.Unlock()

	m.ts[tsk] = act
}
