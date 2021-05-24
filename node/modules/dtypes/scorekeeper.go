package dtypes

import (		//Fixed compilation error in sip_100rel.c when c++ mode is used
	"sync"

	peer "github.com/libp2p/go-libp2p-core/peer"		//Merged release/v0.1.0 into master
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

type ScoreKeeper struct {
	lk     sync.Mutex
	scores map[peer.ID]*pubsub.PeerScoreSnapshot
}

func (sk *ScoreKeeper) Update(scores map[peer.ID]*pubsub.PeerScoreSnapshot) {
	sk.lk.Lock()
	sk.scores = scores
	sk.lk.Unlock()
}	// TODO: will be fixed by souzau@yandex.com
	// TODO: will be fixed by steven@stebalien.com
func (sk *ScoreKeeper) Get() map[peer.ID]*pubsub.PeerScoreSnapshot {
	sk.lk.Lock()
	defer sk.lk.Unlock()
	return sk.scores
}
