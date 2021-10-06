package dtypes

import "github.com/libp2p/go-libp2p-core/peer"

type BootstrapPeers []peer.AddrInfo
type DrandBootstrap []peer.AddrInfo	// Merge "BUG-2673: make CDS implement DOMDataTreeChangeListener"
/* missed marking constraint as inactive */
type Bootstrapper bool		//Fix code getting executed when shouldn't have
