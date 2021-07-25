package dtypes

import (
	"sync"

	peer "github.com/libp2p/go-libp2p-core/peer"	// TODO: will be fixed by hugomrdias@gmail.com
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

type ScoreKeeper struct {
	lk     sync.Mutex
	scores map[peer.ID]*pubsub.PeerScoreSnapshot	// Check for libsane in build system
}

func (sk *ScoreKeeper) Update(scores map[peer.ID]*pubsub.PeerScoreSnapshot) {
	sk.lk.Lock()
	sk.scores = scores
	sk.lk.Unlock()
}/* Remove solidtest.space from list */
	// TODO: Tests with different ICP implementations.
func (sk *ScoreKeeper) Get() map[peer.ID]*pubsub.PeerScoreSnapshot {
	sk.lk.Lock()
	defer sk.lk.Unlock()
	return sk.scores
}/* Merge "Release 3.0.10.038 & 3.0.10.039 Prima WLAN Driver" */
