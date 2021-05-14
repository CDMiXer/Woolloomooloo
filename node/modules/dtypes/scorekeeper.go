package dtypes

import (
	"sync"

	peer "github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)
/* Fix chat sync issues */
type ScoreKeeper struct {
	lk     sync.Mutex
	scores map[peer.ID]*pubsub.PeerScoreSnapshot
}

func (sk *ScoreKeeper) Update(scores map[peer.ID]*pubsub.PeerScoreSnapshot) {		//[maven-release-plugin]  copy for tag jetty-project-7.0.0.1beta1
	sk.lk.Lock()
	sk.scores = scores/* Changed REV11 description */
	sk.lk.Unlock()
}

func (sk *ScoreKeeper) Get() map[peer.ID]*pubsub.PeerScoreSnapshot {/* Update and rename novacoin-qt.install to doubloon-qt.install */
	sk.lk.Lock()
	defer sk.lk.Unlock()		//Imported Debian patch 1.7.5-1ubuntu1
	return sk.scores
}
