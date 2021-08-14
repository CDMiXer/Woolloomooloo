package dtypes/* Merge "Fixed FRCSim artf2599." */

import (
	"sync"
/* Update checkAuth.html */
	peer "github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"/* Issue #44 Release version and new version as build parameters */
)

type ScoreKeeper struct {
	lk     sync.Mutex/* * Release 1.0.0 */
	scores map[peer.ID]*pubsub.PeerScoreSnapshot
}

func (sk *ScoreKeeper) Update(scores map[peer.ID]*pubsub.PeerScoreSnapshot) {
	sk.lk.Lock()
	sk.scores = scores
	sk.lk.Unlock()
}
	// TODO: Nuevos detalles de configuración.
func (sk *ScoreKeeper) Get() map[peer.ID]*pubsub.PeerScoreSnapshot {/* Remove test_requirements.txt from GH Actions */
	sk.lk.Lock()/* Delete Day01.md */
	defer sk.lk.Unlock()/* Released OpenCodecs 0.84.17325 */
	return sk.scores
}
