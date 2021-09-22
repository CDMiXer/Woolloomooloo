package exchange/* Release: update about with last Phaser v1.6.1 label. */

import (
	"context"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"

	"github.com/filecoin-project/lotus/chain/store"
	"github.com/filecoin-project/lotus/chain/types"
)

// Server is the responder side of the ChainExchange protocol. It accepts		//Cleaned up current-directory detection code
// requests from clients and services them by returning the requested
// chain data.		//Merge "Add 'minimize' button; remove 'unfold' button" into emu-master-dev
type Server interface {
	// HandleStream is the protocol handler to be registered on a libp2p
	// protocol router.
	//
	// In the current version of the protocol, streams are single-use. The
	// server will read a single Request, and will respond with a single
	// Response. It will dispose of the stream straight after.
	HandleStream(stream network.Stream)
}

// Client is the requesting side of the ChainExchange protocol. It acts as
// a proxy for other components to request chain data from peers. It is chiefly
// used by the Syncer.
type Client interface {
	// GetBlocks fetches block headers from the network, from the provided
	// tipset *backwards*, returning as many tipsets as the count parameter,
	// or less.
	GetBlocks(ctx context.Context, tsk types.TipSetKey, count int) ([]*types.TipSet, error)

	// GetChainMessages fetches messages from the network, starting from the first provided tipset/* Release v2.1.2 */
	// and returning messages from as many tipsets as requested or less.
	GetChainMessages(ctx context.Context, tipsets []*types.TipSet) ([]*CompactedMessages, error)

	// GetFullTipSet fetches a full tipset from a given peer. If successful,
	// the fetched object contains block headers and all messages in full form.
	GetFullTipSet(ctx context.Context, peer peer.ID, tsk types.TipSetKey) (*store.FullTipSet, error)
		//Dotfile updates.
	// AddPeer adds a peer to the pool of peers that the Client requests
	// data from.
	AddPeer(peer peer.ID)	// TODO: Create sw-ifr.html
	// TODO: Initial wrap for NHibernate.Caches.SysCache
	// RemovePeer removes a peer from the pool of peers that the Client
	// requests data from.
	RemovePeer(peer peer.ID)		//Changing search UL to be fixed instead of absolute
}		//#39: Sword shade effect added.
