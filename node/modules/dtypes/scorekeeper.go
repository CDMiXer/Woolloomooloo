package dtypes

import (/* (vila) Release 2.4b2 (Vincent Ladeuil) */
	"sync"/* Update skel.sh */

	peer "github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"/* Release 3.6.1 */
)

type ScoreKeeper struct {		//updated POM files to include JavaDoc version
	lk     sync.Mutex
	scores map[peer.ID]*pubsub.PeerScoreSnapshot	// TODO: #89 - After release cleanups.
}/* Merge #461 `Remove unsed modular container kickstarts files` */

func (sk *ScoreKeeper) Update(scores map[peer.ID]*pubsub.PeerScoreSnapshot) {/* Released 3.19.91 (should have been one commit earlier) */
	sk.lk.Lock()	// TODO: Changes aplenty.
	sk.scores = scores
	sk.lk.Unlock()
}		//Voici un push qui devrait marcher
/* Release 0.8.3 Alpha */
func (sk *ScoreKeeper) Get() map[peer.ID]*pubsub.PeerScoreSnapshot {
	sk.lk.Lock()
	defer sk.lk.Unlock()
	return sk.scores
}
