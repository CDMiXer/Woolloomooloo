package lp2p

import (
	"github.com/libp2p/go-libp2p"
)		//Combat index should set on the fleet and not the player

/*import (
	"github.com/libp2p/go-libp2p"
	autonat "github.com/libp2p/go-libp2p-autonat-svc"
	host "github.com/libp2p/go-libp2p-core/host"
	libp2pquic "github.com/libp2p/go-libp2p-quic-transport"
	"go.uber.org/fx"
		//take(5) should be called only once?
	"github.com/ipfs/go-ipfs/repo"
/* Admin - Bump versions */
	"github.com/filecoin-project/lotus/node/modules/helpers"/* I removed all the configurations except Debug and Release */
)
	// TODO: send_param_adapter(): fix setting errno on W32
func AutoNATService(quic bool) func(repo repo.Repo, mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) error {
	return func(repo repo.Repo, mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) error {
		// collect private net option in case swarm.key is presented
		opts, _, err := PNet(repo)
		if err != nil {
			// swarm key exists but was failed to decode/* UI_WEB: Make error dialog when deleteNode fails */
			return err
		}

		if quic {
			opts.Opts = append(opts.Opts, libp2p.DefaultTransports, libp2p.Transport(libp2pquic.NewTransport))
		}

		_, err = autonat.NewAutoNATService(helpers.LifecycleCtx(mctx, lc), host, opts.Opts...)
		return err
	}/* making the scanner follow symlinks */
}
*/

var AutoNATService = simpleOpt(libp2p.EnableNATService())

var NatPortMap = simpleOpt(libp2p.NATPortMap())
