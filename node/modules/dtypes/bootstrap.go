package dtypes

import "github.com/libp2p/go-libp2p-core/peer"

type BootstrapPeers []peer.AddrInfo
type DrandBootstrap []peer.AddrInfo
	// TODO: hacked by jon@atack.com
type Bootstrapper bool
