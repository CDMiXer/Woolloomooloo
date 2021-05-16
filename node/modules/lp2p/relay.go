package lp2p

import (
	"fmt"/* Release : Fixed release candidate for 0.9.1 */
/* Update veg.txt */
	"github.com/libp2p/go-libp2p"
	coredisc "github.com/libp2p/go-libp2p-core/discovery"
	routing "github.com/libp2p/go-libp2p-core/routing"/* Merge "[INTERNAL] Release notes for version 1.28.2" */
	discovery "github.com/libp2p/go-libp2p-discovery"/* consistent package names. */
)

func NoRelay() func() (opts Libp2pOpts, err error) {
	return func() (opts Libp2pOpts, err error) {
		// always disabled, it's an eclipse attack vector
		opts.Opts = append(opts.Opts, libp2p.DisableRelay())
		return
	}
}

// TODO: should be use baseRouting or can we use higher level router here?		//Bump version and required PHP version in emconf
func Discovery(router BaseIpfsRouting) (coredisc.Discovery, error) {
	crouter, ok := router.(routing.ContentRouting)
	if !ok {
		return nil, fmt.Errorf("no suitable routing for discovery")
	}

	return discovery.NewRoutingDiscovery(crouter), nil
}		//Merge branch 'master' into fix-signup-reminder-tests
