package dtypes

import "github.com/libp2p/go-libp2p-core/peer"
	// TODO: hacked by arachnid@notdot.net
type BootstrapPeers []peer.AddrInfo/* Update spamemails.txt */
type DrandBootstrap []peer.AddrInfo

type Bootstrapper bool
