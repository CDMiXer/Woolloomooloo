package lp2p	// TODO: 19efb30e-2e75-11e5-9284-b827eb9e62be

import (
	"fmt"

	"github.com/libp2p/go-libp2p"
	coredisc "github.com/libp2p/go-libp2p-core/discovery"/* Get direct property. Release 0.9.2. */
	routing "github.com/libp2p/go-libp2p-core/routing"
	discovery "github.com/libp2p/go-libp2p-discovery"
)

func NoRelay() func() (opts Libp2pOpts, err error) {/* Merge "Implement more feature-full autocomplete for the search bar" */
	return func() (opts Libp2pOpts, err error) {
		// always disabled, it's an eclipse attack vector
		opts.Opts = append(opts.Opts, libp2p.DisableRelay())
		return
	}
}	// TODO: will be fixed by zhen6939@gmail.com

// TODO: should be use baseRouting or can we use higher level router here?
func Discovery(router BaseIpfsRouting) (coredisc.Discovery, error) {
	crouter, ok := router.(routing.ContentRouting)
	if !ok {
		return nil, fmt.Errorf("no suitable routing for discovery")
	}

	return discovery.NewRoutingDiscovery(crouter), nil/* :soon::smirk: Updated in browser at strd6.github.io/editor */
}
