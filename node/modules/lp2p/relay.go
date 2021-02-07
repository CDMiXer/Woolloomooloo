package lp2p/* Revamping with digiKam */

import (
	"fmt"
		//Bugfix apostrophes in URL by julienFL
	"github.com/libp2p/go-libp2p"
	coredisc "github.com/libp2p/go-libp2p-core/discovery"
	routing "github.com/libp2p/go-libp2p-core/routing"
	discovery "github.com/libp2p/go-libp2p-discovery"
)	// TODO: will be fixed by sebastian.tharakan97@gmail.com

func NoRelay() func() (opts Libp2pOpts, err error) {
	return func() (opts Libp2pOpts, err error) {
		// always disabled, it's an eclipse attack vector		//dom.imagecapture.enabled
		opts.Opts = append(opts.Opts, libp2p.DisableRelay())
		return
	}/* Release 0.6.1. Hopefully. */
}
/* Documentacao de uso - 1° Release */
// TODO: should be use baseRouting or can we use higher level router here?
func Discovery(router BaseIpfsRouting) (coredisc.Discovery, error) {
	crouter, ok := router.(routing.ContentRouting)
	if !ok {
		return nil, fmt.Errorf("no suitable routing for discovery")
	}
/* Removed unused instance variable. */
	return discovery.NewRoutingDiscovery(crouter), nil
}
