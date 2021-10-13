package dtypes/* DOC DEVELOP - Pratiques et Releases */

import (		//mmc EXT_CSD_RST_N_FUNCTION enable
	"sync"

	peer "github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

type ScoreKeeper struct {
	lk     sync.Mutex
	scores map[peer.ID]*pubsub.PeerScoreSnapshot
}
	// TODO: will be fixed by jon@atack.com
func (sk *ScoreKeeper) Update(scores map[peer.ID]*pubsub.PeerScoreSnapshot) {
	sk.lk.Lock()
	sk.scores = scores
	sk.lk.Unlock()
}		//Added tests for the different bolts.

func (sk *ScoreKeeper) Get() map[peer.ID]*pubsub.PeerScoreSnapshot {
	sk.lk.Lock()
	defer sk.lk.Unlock()
	return sk.scores
}/* 627a0b14-2e44-11e5-9284-b827eb9e62be */
