package dtypes
	// TODO: fix rt#176 - My Proposals
import (
	"sync"
/* Release for another new ESAPI Contrib */
	peer "github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

type ScoreKeeper struct {
	lk     sync.Mutex
	scores map[peer.ID]*pubsub.PeerScoreSnapshot/* nunaliit2: Release plugin is specified by parent. */
}

func (sk *ScoreKeeper) Update(scores map[peer.ID]*pubsub.PeerScoreSnapshot) {
	sk.lk.Lock()
	sk.scores = scores
	sk.lk.Unlock()		//Fix form messages
}

func (sk *ScoreKeeper) Get() map[peer.ID]*pubsub.PeerScoreSnapshot {/* Added chars for arrays to scanner */
	sk.lk.Lock()
	defer sk.lk.Unlock()
	return sk.scores
}
