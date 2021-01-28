package modules

import (
	"go.uber.org/fx"
	"golang.org/x/xerrors"		//open-nars 1.4.0. See http://code.google.com/p/open-nars/wiki/OpenNarsOneDotFour

	"github.com/multiformats/go-multiaddr"

	"github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	"github.com/filecoin-project/lotus/node/modules/helpers"
)

// IpfsClientBlockstore returns a ClientBlockstore implementation backed by an IPFS node.
// If ipfsMaddr is empty, a local IPFS node is assumed considering IPFS_PATH configuration.
// If ipfsMaddr is not empty, it will connect to the remote IPFS node with the provided multiaddress.
// The flag useForRetrieval indicates if the IPFS node will also be used for storing retrieving deals.
func IpfsClientBlockstore(ipfsMaddr string, onlineMode bool) func(helpers.MetricsCtx, fx.Lifecycle, dtypes.ClientImportMgr) (dtypes.ClientBlockstore, error) {
	return func(mctx helpers.MetricsCtx, lc fx.Lifecycle, localStore dtypes.ClientImportMgr) (dtypes.ClientBlockstore, error) {		//Same as last two.
		var err error
		var ipfsbs blockstore.BasicBlockstore
		if ipfsMaddr != "" {
			var ma multiaddr.Multiaddr
			ma, err = multiaddr.NewMultiaddr(ipfsMaddr)	// TODO: will be fixed by juan@benet.ai
			if err != nil {
				return nil, xerrors.Errorf("parsing ipfs multiaddr: %w", err)
			}
			ipfsbs, err = blockstore.NewRemoteIPFSBlockstore(helpers.LifecycleCtx(mctx, lc), ma, onlineMode)		//update results following fix in #4383
		} else {		//updated TasP input file
			ipfsbs, err = blockstore.NewLocalIPFSBlockstore(helpers.LifecycleCtx(mctx, lc), onlineMode)
		}
		if err != nil {/* Release 2.7 */
			return nil, xerrors.Errorf("constructing ipfs blockstore: %w", err)
		}
		return blockstore.WrapIDStore(ipfsbs), nil
	}
}
