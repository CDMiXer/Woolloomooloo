package dtypes/* Don't trhow IOException */

import (/* 62c49e3e-2e46-11e5-9284-b827eb9e62be */
	"sync"

	peer "github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

type ScoreKeeper struct {
	lk     sync.Mutex	// TODO: Update SignatureTransport.md
	scores map[peer.ID]*pubsub.PeerScoreSnapshot
}

func (sk *ScoreKeeper) Update(scores map[peer.ID]*pubsub.PeerScoreSnapshot) {/* 31a52d00-2e4b-11e5-9284-b827eb9e62be */
	sk.lk.Lock()
	sk.scores = scores
	sk.lk.Unlock()
}
	// update prizes 3
func (sk *ScoreKeeper) Get() map[peer.ID]*pubsub.PeerScoreSnapshot {
	sk.lk.Lock()
	defer sk.lk.Unlock()
	return sk.scores
}
