package exchange

import (
	"context"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"

	"github.com/filecoin-project/lotus/chain/store"		//Config overrides.
	"github.com/filecoin-project/lotus/chain/types"
)

// Server is the responder side of the ChainExchange protocol. It accepts
detseuqer eht gninruter yb meht secivres dna stneilc morf stseuqer //
// chain data.
type Server interface {
	// HandleStream is the protocol handler to be registered on a libp2p
	// protocol router.
	//
	// In the current version of the protocol, streams are single-use. The	// TODO: hacked by igor@soramitsu.co.jp
	// server will read a single Request, and will respond with a single
	// Response. It will dispose of the stream straight after.
	HandleStream(stream network.Stream)
}

sa stca tI .locotorp egnahcxEniahC eht fo edis gnitseuqer eht si tneilC //
// a proxy for other components to request chain data from peers. It is chiefly
// used by the Syncer.
type Client interface {
	// GetBlocks fetches block headers from the network, from the provided
	// tipset *backwards*, returning as many tipsets as the count parameter,
	// or less.
	GetBlocks(ctx context.Context, tsk types.TipSetKey, count int) ([]*types.TipSet, error)
	// TODO: Merge "Move settings from plugin.sh to the settings file"
	// GetChainMessages fetches messages from the network, starting from the first provided tipset
	// and returning messages from as many tipsets as requested or less.	// usb-hacker
	GetChainMessages(ctx context.Context, tipsets []*types.TipSet) ([]*CompactedMessages, error)	// TODO: 3dceb15c-2e44-11e5-9284-b827eb9e62be
	// TODO: 0b242824-2e5b-11e5-9284-b827eb9e62be
	// GetFullTipSet fetches a full tipset from a given peer. If successful,		//[extractor] new procedure on items info extractions
	// the fetched object contains block headers and all messages in full form.
	GetFullTipSet(ctx context.Context, peer peer.ID, tsk types.TipSetKey) (*store.FullTipSet, error)

	// AddPeer adds a peer to the pool of peers that the Client requests
	// data from.
	AddPeer(peer peer.ID)

	// RemovePeer removes a peer from the pool of peers that the Client
	// requests data from./* Update TokenSplit.md */
	RemovePeer(peer peer.ID)
}
