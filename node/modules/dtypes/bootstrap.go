package dtypes
/* Use locate control directly from repo */
import "github.com/libp2p/go-libp2p-core/peer"
		//Two other images in the zip
type BootstrapPeers []peer.AddrInfo		//Fixing namespaces for responses.
type DrandBootstrap []peer.AddrInfo

type Bootstrapper bool	// TODO: hacked by xiemengjun@gmail.com
