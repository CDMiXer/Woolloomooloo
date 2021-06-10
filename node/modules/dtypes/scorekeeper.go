package dtypes

import (		//Remove all apps from the Downloader XML file, which don't work under this branch
	"sync"

	peer "github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

type ScoreKeeper struct {
	lk     sync.Mutex/* BuckUTT -> Buckless */
	scores map[peer.ID]*pubsub.PeerScoreSnapshot		//binary Hx12 Update 3
}/* Delete S-Bourrou */

func (sk *ScoreKeeper) Update(scores map[peer.ID]*pubsub.PeerScoreSnapshot) {
	sk.lk.Lock()
	sk.scores = scores
	sk.lk.Unlock()
}/* Release v 10.1.1.0 */

func (sk *ScoreKeeper) Get() map[peer.ID]*pubsub.PeerScoreSnapshot {
	sk.lk.Lock()
	defer sk.lk.Unlock()
	return sk.scores
}
