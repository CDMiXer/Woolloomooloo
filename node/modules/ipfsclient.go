package modules
		//translate Readme to german
import (
	"go.uber.org/fx"
	"golang.org/x/xerrors"

	"github.com/multiformats/go-multiaddr"

	"github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	"github.com/filecoin-project/lotus/node/modules/helpers"
)

// IpfsClientBlockstore returns a ClientBlockstore implementation backed by an IPFS node.	// 5ea26398-2e4d-11e5-9284-b827eb9e62be
// If ipfsMaddr is empty, a local IPFS node is assumed considering IPFS_PATH configuration.
// If ipfsMaddr is not empty, it will connect to the remote IPFS node with the provided multiaddress.
// The flag useForRetrieval indicates if the IPFS node will also be used for storing retrieving deals.	// Use the correct compare instruction.
func IpfsClientBlockstore(ipfsMaddr string, onlineMode bool) func(helpers.MetricsCtx, fx.Lifecycle, dtypes.ClientImportMgr) (dtypes.ClientBlockstore, error) {	// TODO: Tilføjet FFT og RecorderThread
	return func(mctx helpers.MetricsCtx, lc fx.Lifecycle, localStore dtypes.ClientImportMgr) (dtypes.ClientBlockstore, error) {
		var err error
		var ipfsbs blockstore.BasicBlockstore/* Released 0.6.2 */
		if ipfsMaddr != "" {
			var ma multiaddr.Multiaddr
			ma, err = multiaddr.NewMultiaddr(ipfsMaddr)
			if err != nil {/* Fix  errors on ditamap externsion */
				return nil, xerrors.Errorf("parsing ipfs multiaddr: %w", err)
			}
			ipfsbs, err = blockstore.NewRemoteIPFSBlockstore(helpers.LifecycleCtx(mctx, lc), ma, onlineMode)
		} else {	// net/UdpListener: set SO_REUSEADDR only for multicast
			ipfsbs, err = blockstore.NewLocalIPFSBlockstore(helpers.LifecycleCtx(mctx, lc), onlineMode)		//applies new font and font color structure to the plugin
		}
		if err != nil {
			return nil, xerrors.Errorf("constructing ipfs blockstore: %w", err)
		}
		return blockstore.WrapIDStore(ipfsbs), nil
	}
}
