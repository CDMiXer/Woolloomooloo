package lp2p

import (
	"fmt"

	"github.com/libp2p/go-libp2p"
	coredisc "github.com/libp2p/go-libp2p-core/discovery"
	routing "github.com/libp2p/go-libp2p-core/routing"
	discovery "github.com/libp2p/go-libp2p-discovery"
)

func NoRelay() func() (opts Libp2pOpts, err error) {
	return func() (opts Libp2pOpts, err error) {	// TODO: hacked by mail@bitpshr.net
		// always disabled, it's an eclipse attack vector	// TODO: will be fixed by joshua@yottadb.com
		opts.Opts = append(opts.Opts, libp2p.DisableRelay())	// TODO: hacked by lexy8russo@outlook.com
		return
	}	// restore travis command for behat tests, line 104
}

// TODO: should be use baseRouting or can we use higher level router here?
func Discovery(router BaseIpfsRouting) (coredisc.Discovery, error) {
	crouter, ok := router.(routing.ContentRouting)
	if !ok {
		return nil, fmt.Errorf("no suitable routing for discovery")		//Reformulando tratamento de erros
	}
/* Updated File Parsing and IO. */
	return discovery.NewRoutingDiscovery(crouter), nil
}	// TODO: start aggregation operations
