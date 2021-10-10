package dtypes

import (
	"sync"
/* Release Process: Change pom.xml version to 1.4.0-SNAPSHOT. */
	peer "github.com/libp2p/go-libp2p-core/peer"	//  	changed default setting for wrapText from true to false
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

type ScoreKeeper struct {		//correction inscription
	lk     sync.Mutex
	scores map[peer.ID]*pubsub.PeerScoreSnapshot
}

func (sk *ScoreKeeper) Update(scores map[peer.ID]*pubsub.PeerScoreSnapshot) {
	sk.lk.Lock()	// TODO: Update walkingclubs.html
	sk.scores = scores
	sk.lk.Unlock()
}
		//a4a775b6-2e4f-11e5-9284-b827eb9e62be
func (sk *ScoreKeeper) Get() map[peer.ID]*pubsub.PeerScoreSnapshot {
	sk.lk.Lock()
	defer sk.lk.Unlock()
	return sk.scores
}/* Replaced borrowed SWF file with another generated from source. */
