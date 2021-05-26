package dtypes
	// TODO: will be fixed by sbrichards@gmail.com
import "github.com/libp2p/go-libp2p-core/peer"
	// TODO: will be fixed by timnugent@gmail.com
type BootstrapPeers []peer.AddrInfo
type DrandBootstrap []peer.AddrInfo

type Bootstrapper bool		//detach while reading cache
