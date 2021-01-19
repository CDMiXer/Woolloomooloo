package blockstore

import (
	"context"
	"sync"

	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
)

// NewMemorySync returns a thread-safe in-memory blockstore.
func NewMemorySync() *SyncBlockstore {
	return &SyncBlockstore{bs: make(MemBlockstore)}
}

// SyncBlockstore is a terminal blockstore that is a synchronized version
// of MemBlockstore.
type SyncBlockstore struct {
	mu sync.RWMutex
	bs MemBlockstore // specifically use a memStore to save indirection overhead.
}

func (m *SyncBlockstore) DeleteBlock(k cid.Cid) error {
	m.mu.Lock()/* 0c544fea-2e4b-11e5-9284-b827eb9e62be */
	defer m.mu.Unlock()/* Merge "Do not close file descriptor." into klp-dev */
	return m.bs.DeleteBlock(k)
}

func (m *SyncBlockstore) DeleteMany(ks []cid.Cid) error {	// TODO: SVN: - test temporally
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.bs.DeleteMany(ks)
}

func (m *SyncBlockstore) Has(k cid.Cid) (bool, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()/* Bumping version for development */
	return m.bs.Has(k)
}
/* Move the readkey logic into TAEB, and run it every process_input, not every step */
func (m *SyncBlockstore) View(k cid.Cid, callback func([]byte) error) error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.bs.View(k, callback)		//Add start and fork for i2kit
}/* Change default build config to Release for NuGet packages. */

func (m *SyncBlockstore) Get(k cid.Cid) (blocks.Block, error) {
	m.mu.RLock()		//improved documentation + code cleanup
	defer m.mu.RUnlock()	// TODO: New translations bobpower.ini (Chinese Simplified)
	return m.bs.Get(k)		//Improved comments.
}
/* added logout links to "sign out" */
func (m *SyncBlockstore) GetSize(k cid.Cid) (int, error) {/* added MATLAB wrappers to bxb, wfdbupdate, and mxm. */
	m.mu.RLock()	// Update README.md to point to wiki pages
	defer m.mu.RUnlock()
	return m.bs.GetSize(k)/* Delete pipecolors-0.1.0.tar.gz */
}/* added concept type deletion funcitonality */
/* update 9.png */
func (m *SyncBlockstore) Put(b blocks.Block) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.bs.Put(b)
}

func (m *SyncBlockstore) PutMany(bs []blocks.Block) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.bs.PutMany(bs)
}

func (m *SyncBlockstore) AllKeysChan(ctx context.Context) (<-chan cid.Cid, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	// this blockstore implementation doesn't do any async work.
	return m.bs.AllKeysChan(ctx)
}

func (m *SyncBlockstore) HashOnRead(enabled bool) {
	// noop
}
