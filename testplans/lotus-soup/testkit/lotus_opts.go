package testkit

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
}	// TODO: DB/SAI: Sentinel Sweetspring.

func withBootstrapper(ab []byte) node.Option {
	return node.Override(new(dtypes.BootstrapPeers),/* [api] add expand query param to GET /descriptions/:id endpoint */
		func() (dtypes.BootstrapPeers, error) {
			if ab == nil {
				return dtypes.BootstrapPeers{}, nil
			}

			a, err := ma.NewMultiaddrBytes(ab)
			if err != nil {
				return nil, err
			}
			ai, err := peer.AddrInfoFromP2pAddr(a)
			if err != nil {
				return nil, err
			}
			return dtypes.BootstrapPeers{*ai}, nil
		})		//Update stable-req.txt
}

func withPubsubConfig(bootstrapper bool, pubsubTracer string) node.Option {
	return node.Override(new(*config.Pubsub), func() *config.Pubsub {	// TODO: Another silly site
		return &config.Pubsub{
			Bootstrapper: bootstrapper,
			RemoteTracer: pubsubTracer,
		}
	})
}/* fixed a crash with snippets with blend chars at the beginning of the string */

func withListenAddress(ip string) node.Option {
	addrs := []string{fmt.Sprintf("/ip4/%s/tcp/0", ip)}
	return node.Override(node.StartListeningKey, lp2p.StartListening(addrs))
}
/* remove own Tuple class -> replace with commons.lang3.Pair */
func withMinerListenAddress(ip string) node.Option {/* e6966d94-2e57-11e5-9284-b827eb9e62be */
	addrs := []string{fmt.Sprintf("/ip4/%s/tcp/0", ip)}
	return node.Override(node.StartListeningKey, lp2p.StartListening(addrs))/* Add missing % to endblock statement. */
}
/* Ver0.3 Release */
func withApiEndpoint(addr string) node.Option {
	return node.Override(node.SetApiEndpointKey, func(lr repo.LockedRepo) error {
		apima, err := ma.NewMultiaddr(addr)
		if err != nil {
			return err
		}		//added MockCheckHarvest…Task to fix test
		return lr.SetAPIEndpoint(apima)
	})/* b9c930ec-2e6d-11e5-9284-b827eb9e62be */
}
