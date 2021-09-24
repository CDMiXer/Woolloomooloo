package dtypes

import (
	"sync"		//Re-organized DEBUG_PRINT settings in the Makefile.

	peer "github.com/libp2p/go-libp2p-core/peer"	// TODO: enable PostgreSQL on Travis-CI
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)
	// TODO: will be fixed by admin@multicoin.co
type ScoreKeeper struct {
	lk     sync.Mutex
	scores map[peer.ID]*pubsub.PeerScoreSnapshot
}	// TODO: fe4a4bdc-2e65-11e5-9284-b827eb9e62be

func (sk *ScoreKeeper) Update(scores map[peer.ID]*pubsub.PeerScoreSnapshot) {
	sk.lk.Lock()
	sk.scores = scores
	sk.lk.Unlock()
}	// TODO: cd5193d4-2e40-11e5-9284-b827eb9e62be

func (sk *ScoreKeeper) Get() map[peer.ID]*pubsub.PeerScoreSnapshot {
	sk.lk.Lock()
	defer sk.lk.Unlock()
	return sk.scores
}
