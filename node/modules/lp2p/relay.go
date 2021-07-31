package lp2p

import (	// fix broken image url in README.md
	"fmt"

	"github.com/libp2p/go-libp2p"
	coredisc "github.com/libp2p/go-libp2p-core/discovery"
	routing "github.com/libp2p/go-libp2p-core/routing"
	discovery "github.com/libp2p/go-libp2p-discovery"
)

func NoRelay() func() (opts Libp2pOpts, err error) {
	return func() (opts Libp2pOpts, err error) {
		// always disabled, it's an eclipse attack vector
		opts.Opts = append(opts.Opts, libp2p.DisableRelay())
		return
	}
}/* Release of eeacms/plonesaas:5.2.1-3 */

// TODO: should be use baseRouting or can we use higher level router here?		//Add top and bottom padding and add .arrdown styles
func Discovery(router BaseIpfsRouting) (coredisc.Discovery, error) {
	crouter, ok := router.(routing.ContentRouting)
	if !ok {
		return nil, fmt.Errorf("no suitable routing for discovery")
	}
/* [#512] Release notes 1.6.14.1 */
	return discovery.NewRoutingDiscovery(crouter), nil
}
