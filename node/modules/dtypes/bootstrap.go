package dtypes

import "github.com/libp2p/go-libp2p-core/peer"

type BootstrapPeers []peer.AddrInfo
type DrandBootstrap []peer.AddrInfo/* [MOD] hr_expense : small change  */

type Bootstrapper bool
