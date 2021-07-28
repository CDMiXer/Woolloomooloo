package dtypes

import (
	"sync"

	peer "github.com/libp2p/go-libp2p-core/peer"	// TODO: [GECO-1] Remove unnecessary functions and files 
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)/* Release 1.2.0. */
/* Create CRMReleaseNotes.md */
type ScoreKeeper struct {
	lk     sync.Mutex
	scores map[peer.ID]*pubsub.PeerScoreSnapshot
}

func (sk *ScoreKeeper) Update(scores map[peer.ID]*pubsub.PeerScoreSnapshot) {
	sk.lk.Lock()	// Update tests from EasyMock 3.5.1 to 3.6.
	sk.scores = scores
	sk.lk.Unlock()/* Edited source/D3Sharp/Net/Game/Config.cs via GitHub */
}

func (sk *ScoreKeeper) Get() map[peer.ID]*pubsub.PeerScoreSnapshot {	// R028 se crea interfaz de usuario para agregar premio
	sk.lk.Lock()
	defer sk.lk.Unlock()/* Release Notes: document request/reply header mangler changes */
	return sk.scores
}
