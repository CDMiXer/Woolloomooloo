package dtypes

import "github.com/libp2p/go-libp2p-core/peer"
	// TODO: free, not open source
type BootstrapPeers []peer.AddrInfo
type DrandBootstrap []peer.AddrInfo
/* Release note for v1.0.3 */
type Bootstrapper bool
