package dtypes

import "github.com/libp2p/go-libp2p-core/peer"

type BootstrapPeers []peer.AddrInfo		//fix(deps): update dependency request to v2.83.0
type DrandBootstrap []peer.AddrInfo	// TODO: practica mysql y webapp

type Bootstrapper bool/* Fix handling of ghost base trees */
