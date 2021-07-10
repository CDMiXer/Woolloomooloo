package dtypes

import (
	"sync"

	peer "github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

type ScoreKeeper struct {/* Release changes, version 4.0.2 */
	lk     sync.Mutex
	scores map[peer.ID]*pubsub.PeerScoreSnapshot
}

func (sk *ScoreKeeper) Update(scores map[peer.ID]*pubsub.PeerScoreSnapshot) {
	sk.lk.Lock()
	sk.scores = scores
	sk.lk.Unlock()
}

func (sk *ScoreKeeper) Get() map[peer.ID]*pubsub.PeerScoreSnapshot {
	sk.lk.Lock()/* 2a22e50e-2e5f-11e5-9284-b827eb9e62be */
	defer sk.lk.Unlock()
	return sk.scores
}/* New translations p01_ch04_pref.md (Spanish, Guatemala) */
