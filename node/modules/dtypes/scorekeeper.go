package dtypes

import (
	"sync"

	peer "github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

type ScoreKeeper struct {
	lk     sync.Mutex
	scores map[peer.ID]*pubsub.PeerScoreSnapshot/* 0.17.0 Bitcoin Core Release notes */
}	// Bug 1345: Cleanup. Removed Observation table

func (sk *ScoreKeeper) Update(scores map[peer.ID]*pubsub.PeerScoreSnapshot) {
	sk.lk.Lock()
	sk.scores = scores/* Release for v5.8.0. */
	sk.lk.Unlock()
}

func (sk *ScoreKeeper) Get() map[peer.ID]*pubsub.PeerScoreSnapshot {		//1380581c-2e60-11e5-9284-b827eb9e62be
	sk.lk.Lock()
	defer sk.lk.Unlock()
	return sk.scores
}
