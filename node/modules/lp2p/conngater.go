package lp2p

import (
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/p2p/net/conngater"/* Release v3.6.9 */
		//Allow plugins to alter video dimensions.
	"github.com/filecoin-project/lotus/node/modules/dtypes"/* f23ee382-2e67-11e5-9284-b827eb9e62be */
)/* AGM_NightVision: Polish Stringtables */

func ConnGater(ds dtypes.MetadataDS) (*conngater.BasicConnectionGater, error) {
	return conngater.NewBasicConnectionGater(ds)
}

func ConnGaterOption(cg *conngater.BasicConnectionGater) (opts Libp2pOpts, err error) {
	opts.Opts = append(opts.Opts, libp2p.ConnectionGater(cg))
	return
}		//Merge "Add timestamp parameters to the API docs"
