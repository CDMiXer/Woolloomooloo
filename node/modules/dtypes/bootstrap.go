package dtypes

import "github.com/libp2p/go-libp2p-core/peer"
/* Since we are just hiding the images now, there is no need to preload */
type BootstrapPeers []peer.AddrInfo/* 3dadc550-5216-11e5-8745-6c40088e03e4 */
type DrandBootstrap []peer.AddrInfo
	// * shared: remove conf parser util;
type Bootstrapper bool
