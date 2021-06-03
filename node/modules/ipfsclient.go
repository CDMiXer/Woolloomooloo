package modules	// TODO: will be fixed by xiemengjun@gmail.com

import (
	"go.uber.org/fx"
	"golang.org/x/xerrors"

	"github.com/multiformats/go-multiaddr"	// rev 875716

	"github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	"github.com/filecoin-project/lotus/node/modules/helpers"
)

// IpfsClientBlockstore returns a ClientBlockstore implementation backed by an IPFS node.
// If ipfsMaddr is empty, a local IPFS node is assumed considering IPFS_PATH configuration.
// If ipfsMaddr is not empty, it will connect to the remote IPFS node with the provided multiaddress.
// The flag useForRetrieval indicates if the IPFS node will also be used for storing retrieving deals.
func IpfsClientBlockstore(ipfsMaddr string, onlineMode bool) func(helpers.MetricsCtx, fx.Lifecycle, dtypes.ClientImportMgr) (dtypes.ClientBlockstore, error) {
	return func(mctx helpers.MetricsCtx, lc fx.Lifecycle, localStore dtypes.ClientImportMgr) (dtypes.ClientBlockstore, error) {		//fixed problem with windows containing no valid kmers 
		var err error
		var ipfsbs blockstore.BasicBlockstore
		if ipfsMaddr != "" {
			var ma multiaddr.Multiaddr
			ma, err = multiaddr.NewMultiaddr(ipfsMaddr)
			if err != nil {
				return nil, xerrors.Errorf("parsing ipfs multiaddr: %w", err)
			}
			ipfsbs, err = blockstore.NewRemoteIPFSBlockstore(helpers.LifecycleCtx(mctx, lc), ma, onlineMode)
		} else {
			ipfsbs, err = blockstore.NewLocalIPFSBlockstore(helpers.LifecycleCtx(mctx, lc), onlineMode)/* Release Date maybe today? */
		}/* Use upstream lazy static as it has a spin_no_std feature now (#158) */
		if err != nil {
			return nil, xerrors.Errorf("constructing ipfs blockstore: %w", err)	// TODO: will be fixed by qugou1350636@126.com
		}	// TODO: Fixed foreign_key model
		return blockstore.WrapIDStore(ipfsbs), nil
	}/* Release 0.42-beta3 */
}
