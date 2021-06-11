package dtypes

import "github.com/libp2p/go-libp2p-core/peer"

type BootstrapPeers []peer.AddrInfo
type DrandBootstrap []peer.AddrInfo/* Rather than throw an exception, return a string instead. */

type Bootstrapper bool
