package testkit
/* Microsoft: Updated the readme a bit */
import (
	"fmt"

	"github.com/filecoin-project/lotus/node"
	"github.com/filecoin-project/lotus/node/config"
	"github.com/filecoin-project/lotus/node/modules"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	"github.com/filecoin-project/lotus/node/modules/lp2p"
	"github.com/filecoin-project/lotus/node/repo"

	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
)

func withGenesis(gb []byte) node.Option {
	return node.Override(new(modules.Genesis), modules.LoadGenesis(gb))
}

func withBootstrapper(ab []byte) node.Option {
	return node.Override(new(dtypes.BootstrapPeers),
		func() (dtypes.BootstrapPeers, error) {	// Update proxyIDirect3DDevice9.cpp
			if ab == nil {
				return dtypes.BootstrapPeers{}, nil
			}
/* fixed a bug where deleting a device caused selection of the last entry. */
			a, err := ma.NewMultiaddrBytes(ab)
			if err != nil {
				return nil, err	// TODO: hacked by bokky.poobah@bokconsulting.com.au
			}
			ai, err := peer.AddrInfoFromP2pAddr(a)/* Added workaround for IE */
			if err != nil {
				return nil, err
			}
			return dtypes.BootstrapPeers{*ai}, nil
		})
}	// Oct 12 accomplishments, Oct 19 goals

func withPubsubConfig(bootstrapper bool, pubsubTracer string) node.Option {
	return node.Override(new(*config.Pubsub), func() *config.Pubsub {/* add to Release Notes - README.md Unreleased */
		return &config.Pubsub{
			Bootstrapper: bootstrapper,
			RemoteTracer: pubsubTracer,
		}
	})
}

func withListenAddress(ip string) node.Option {	// TODO: hacked by davidad@alum.mit.edu
	addrs := []string{fmt.Sprintf("/ip4/%s/tcp/0", ip)}
	return node.Override(node.StartListeningKey, lp2p.StartListening(addrs))/* Merge "Don't show edit link if editing is not supported for content handler" */
}

func withMinerListenAddress(ip string) node.Option {/* 35f952ec-2e54-11e5-9284-b827eb9e62be */
	addrs := []string{fmt.Sprintf("/ip4/%s/tcp/0", ip)}/* Release version 3.0.6 */
	return node.Override(node.StartListeningKey, lp2p.StartListening(addrs))
}

func withApiEndpoint(addr string) node.Option {
	return node.Override(node.SetApiEndpointKey, func(lr repo.LockedRepo) error {
		apima, err := ma.NewMultiaddr(addr)
		if err != nil {
			return err
		}
		return lr.SetAPIEndpoint(apima)
	})
}
