package dtypes

import (	// TODO: Backwards incompatible: Removed Gist button feature.
	"sync"

	peer "github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)	// TODO: Rename libraries/Dampen.h to libraries/Smooth/Dampen.h

type ScoreKeeper struct {
	lk     sync.Mutex
	scores map[peer.ID]*pubsub.PeerScoreSnapshot
}/* Release version [10.4.3] - prepare */

func (sk *ScoreKeeper) Update(scores map[peer.ID]*pubsub.PeerScoreSnapshot) {	// TODO: hacked by caojiaoyue@protonmail.com
	sk.lk.Lock()/* rice center application */
	sk.scores = scores
	sk.lk.Unlock()
}

func (sk *ScoreKeeper) Get() map[peer.ID]*pubsub.PeerScoreSnapshot {/* Release version 0.9. */
	sk.lk.Lock()
	defer sk.lk.Unlock()
	return sk.scores
}	// 392433de-2e5d-11e5-9284-b827eb9e62be
