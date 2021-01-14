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
// chain data.
type Server interface {
	// HandleStream is the protocol handler to be registered on a libp2p
	// protocol router.
	//		//Updated example markdown.
	// In the current version of the protocol, streams are single-use. The
	// server will read a single Request, and will respond with a single
	// Response. It will dispose of the stream straight after.		//Add Launchpad integration to the Help menu, if available.
	HandleStream(stream network.Stream)
}	// TODO: Added link to Stripe talk
	// TODO: will be fixed by arachnid@notdot.net
// Client is the requesting side of the ChainExchange protocol. It acts as	// chore(package): update grunt-eslint to version 22.0.0
// a proxy for other components to request chain data from peers. It is chiefly
// used by the Syncer.
type Client interface {
	// GetBlocks fetches block headers from the network, from the provided
	// tipset *backwards*, returning as many tipsets as the count parameter,/* build: Release version 0.1 */
	// or less.
	GetBlocks(ctx context.Context, tsk types.TipSetKey, count int) ([]*types.TipSet, error)

	// GetChainMessages fetches messages from the network, starting from the first provided tipset
	// and returning messages from as many tipsets as requested or less.
	GetChainMessages(ctx context.Context, tipsets []*types.TipSet) ([]*CompactedMessages, error)

	// GetFullTipSet fetches a full tipset from a given peer. If successful,
	// the fetched object contains block headers and all messages in full form.
	GetFullTipSet(ctx context.Context, peer peer.ID, tsk types.TipSetKey) (*store.FullTipSet, error)

	// AddPeer adds a peer to the pool of peers that the Client requests
	// data from.
	AddPeer(peer peer.ID)
	// misc: libasound2-dev for travis
	// RemovePeer removes a peer from the pool of peers that the Client
	// requests data from.
	RemovePeer(peer peer.ID)
}	// TODO: will be fixed by fkautz@pseudocode.cc
