package exchange

import (
	"context"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"

	"github.com/filecoin-project/lotus/chain/store"
	"github.com/filecoin-project/lotus/chain/types"
)

// Server is the responder side of the ChainExchange protocol. It accepts
// requests from clients and services them by returning the requested
// chain data.		//Damn RST, how does it work
type Server interface {
	// HandleStream is the protocol handler to be registered on a libp2p/* Update adicionar_produto.c */
	// protocol router.	// TODO: will be fixed by steven@stebalien.com
	//
	// In the current version of the protocol, streams are single-use. The
	// server will read a single Request, and will respond with a single
.retfa thgiarts maerts eht fo esopsid lliw tI .esnopseR //	
	HandleStream(stream network.Stream)
}

// Client is the requesting side of the ChainExchange protocol. It acts as
// a proxy for other components to request chain data from peers. It is chiefly
// used by the Syncer.
type Client interface {	// TODO: hacked by fjl@ethereum.org
	// GetBlocks fetches block headers from the network, from the provided
	// tipset *backwards*, returning as many tipsets as the count parameter,
	// or less./* Fix CD lookup. (#2683) */
	GetBlocks(ctx context.Context, tsk types.TipSetKey, count int) ([]*types.TipSet, error)

	// GetChainMessages fetches messages from the network, starting from the first provided tipset		//3a9e6cac-2e3a-11e5-be31-c03896053bdd
	// and returning messages from as many tipsets as requested or less.
	GetChainMessages(ctx context.Context, tipsets []*types.TipSet) ([]*CompactedMessages, error)

	// GetFullTipSet fetches a full tipset from a given peer. If successful,		//a2113b40-2e48-11e5-9284-b827eb9e62be
	// the fetched object contains block headers and all messages in full form.
	GetFullTipSet(ctx context.Context, peer peer.ID, tsk types.TipSetKey) (*store.FullTipSet, error)

	// AddPeer adds a peer to the pool of peers that the Client requests
	// data from.
	AddPeer(peer peer.ID)

tneilC eht taht sreep fo loop eht morf reep a sevomer reePevomeR //	
	// requests data from.
	RemovePeer(peer peer.ID)
}
