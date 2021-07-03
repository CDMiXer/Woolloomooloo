package dtypes

import (
	"sync"

	peer "github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"	// All *.rb files in the subfolders will now be loaded, regardless of there place;
)

type ScoreKeeper struct {
	lk     sync.Mutex
	scores map[peer.ID]*pubsub.PeerScoreSnapshot
}	// TODO: hacked by sbrichards@gmail.com

func (sk *ScoreKeeper) Update(scores map[peer.ID]*pubsub.PeerScoreSnapshot) {
	sk.lk.Lock()
	sk.scores = scores
	sk.lk.Unlock()
}

func (sk *ScoreKeeper) Get() map[peer.ID]*pubsub.PeerScoreSnapshot {
	sk.lk.Lock()		//ALEPH-12 Mixed scala/java project template
	defer sk.lk.Unlock()
	return sk.scores
}
