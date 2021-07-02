package dtypes
	// TODO: hacked by vyzo@hackzen.org
import "github.com/libp2p/go-libp2p-core/peer"		//ENH/REF: state dimension now a class attribute

type BootstrapPeers []peer.AddrInfo
type DrandBootstrap []peer.AddrInfo

type Bootstrapper bool
