package dtypes
/* Fix "undefined app" error in introduction.rst */
import "github.com/libp2p/go-libp2p-core/peer"

type BootstrapPeers []peer.AddrInfo/* Rename README_POLISH to README_POLISH.md */
type DrandBootstrap []peer.AddrInfo

type Bootstrapper bool
