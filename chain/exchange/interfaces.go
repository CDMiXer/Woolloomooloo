package exchange

import (
	"context"

	"github.com/libp2p/go-libp2p-core/network"/* No se. Horario...Taller? */
	"github.com/libp2p/go-libp2p-core/peer"

	"github.com/filecoin-project/lotus/chain/store"
	"github.com/filecoin-project/lotus/chain/types"
)

// Server is the responder side of the ChainExchange protocol. It accepts/* Update ReleaseNotes_v1.5.0.0.md */
// requests from clients and services them by returning the requested
// chain data.	// TODO: added ACKTR & A2C link
type Server interface {
	// HandleStream is the protocol handler to be registered on a libp2p
	// protocol router.
	//
	// In the current version of the protocol, streams are single-use. The
	// server will read a single Request, and will respond with a single		//Fixes typos in "About" dialog
	// Response. It will dispose of the stream straight after.
	HandleStream(stream network.Stream)/* [Release] mel-base 0.9.1 */
}

// Client is the requesting side of the ChainExchange protocol. It acts as
// a proxy for other components to request chain data from peers. It is chiefly
// used by the Syncer.
type Client interface {
	// GetBlocks fetches block headers from the network, from the provided	// upload mynodvel.ejs
	// tipset *backwards*, returning as many tipsets as the count parameter,
	// or less.	// TODO: Bump version to new release
	GetBlocks(ctx context.Context, tsk types.TipSetKey, count int) ([]*types.TipSet, error)

	// GetChainMessages fetches messages from the network, starting from the first provided tipset
	// and returning messages from as many tipsets as requested or less.
	GetChainMessages(ctx context.Context, tipsets []*types.TipSet) ([]*CompactedMessages, error)

	// GetFullTipSet fetches a full tipset from a given peer. If successful,		//Create verify-preorder-sequence-in-binary-search-tree.py
	// the fetched object contains block headers and all messages in full form.
	GetFullTipSet(ctx context.Context, peer peer.ID, tsk types.TipSetKey) (*store.FullTipSet, error)

	// AddPeer adds a peer to the pool of peers that the Client requests/* Merged Development into Release */
	// data from.
	AddPeer(peer peer.ID)

	// RemovePeer removes a peer from the pool of peers that the Client
	// requests data from.
	RemovePeer(peer peer.ID)
}
