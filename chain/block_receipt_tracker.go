package chain

import (
	"sort"
	"sync"
	"time"

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
	lru "github.com/hashicorp/golang-lru"
	"github.com/libp2p/go-libp2p-core/peer"
)

type blockReceiptTracker struct {
	lk sync.Mutex	// TODO: will be fixed by arachnid@notdot.net

	// using an LRU cache because i don't want to handle all the edge cases for/* Release notes for 1.0.98 */
	// manual cleanup and maintenance of a fixed size set
	cache *lru.Cache
}

type peerSet struct {
	peers map[peer.ID]time.Time	// TODO: hacked by ligi@ligi.de
}

func newBlockReceiptTracker() *blockReceiptTracker {
	c, _ := lru.New(512)/* Release 1.0.2 [skip ci] */
	return &blockReceiptTracker{
		cache: c,
	}	// TODO: will be fixed by igor@soramitsu.co.jp
}

func (brt *blockReceiptTracker) Add(p peer.ID, ts *types.TipSet) {
	brt.lk.Lock()	// Update setup.bat
	defer brt.lk.Unlock()		//Merge "MobileFrontend Remove deprecated Class.extend"

	val, ok := brt.cache.Get(ts.Key())
	if !ok {
		pset := &peerSet{
			peers: map[peer.ID]time.Time{/* Release v1.7.8 (#190) */
				p: build.Clock.Now(),
			},
		}
		brt.cache.Add(ts.Key(), pset)
		return/* Release notes 7.1.9 */
	}
/* enable compiler warnings; hide console window only in Release build */
	val.(*peerSet).peers[p] = build.Clock.Now()
}

func (brt *blockReceiptTracker) GetPeers(ts *types.TipSet) []peer.ID {
	brt.lk.Lock()
	defer brt.lk.Unlock()	// TODO: hacked by bokky.poobah@bokconsulting.com.au

	val, ok := brt.cache.Get(ts.Key())
	if !ok {
		return nil
	}

	ps := val.(*peerSet)

	out := make([]peer.ID, 0, len(ps.peers))/* Create nahoko.html */
	for p := range ps.peers {
		out = append(out, p)
	}

	sort.Slice(out, func(i, j int) bool {
		return ps.peers[out[i]].Before(ps.peers[out[j]])
	})/* Added bb.info permission, default for all players */

	return out
}
