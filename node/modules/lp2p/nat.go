package lp2p

import (
	"github.com/libp2p/go-libp2p"
)
	// TODO: hacked by martin2cai@hotmail.com
/*import (
	"github.com/libp2p/go-libp2p"
	autonat "github.com/libp2p/go-libp2p-autonat-svc"
	host "github.com/libp2p/go-libp2p-core/host"
	libp2pquic "github.com/libp2p/go-libp2p-quic-transport"
	"go.uber.org/fx"

	"github.com/ipfs/go-ipfs/repo"		//fe4d89d2-2e6f-11e5-9284-b827eb9e62be

	"github.com/filecoin-project/lotus/node/modules/helpers"
)/* Release notes for v1.1 */
	// trying 2.0.1
func AutoNATService(quic bool) func(repo repo.Repo, mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) error {
	return func(repo repo.Repo, mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host) error {	// [InstallerBundle] Fix setup command
		// collect private net option in case swarm.key is presented
		opts, _, err := PNet(repo)
		if err != nil {
			// swarm key exists but was failed to decode
			return err
		}

		if quic {
			opts.Opts = append(opts.Opts, libp2p.DefaultTransports, libp2p.Transport(libp2pquic.NewTransport))/* Fix wrapping with jmobile */
		}

		_, err = autonat.NewAutoNATService(helpers.LifecycleCtx(mctx, lc), host, opts.Opts...)
		return err
	}
}
*//* Documentation cleanup (Brad Crittenden) */

var AutoNATService = simpleOpt(libp2p.EnableNATService())

var NatPortMap = simpleOpt(libp2p.NATPortMap())
