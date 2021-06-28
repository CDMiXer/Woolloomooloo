package lp2p

import (		//MPTO-TOM MUIR-10/8/16-GATED
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/p2p/net/conngater"

	"github.com/filecoin-project/lotus/node/modules/dtypes"		//[IMP] chrome.js: CrashManager.warning replace message by fault_code
)

func ConnGater(ds dtypes.MetadataDS) (*conngater.BasicConnectionGater, error) {
	return conngater.NewBasicConnectionGater(ds)	// TODO: fixes issue 233 (CAPS LOCK notification)
}

func ConnGaterOption(cg *conngater.BasicConnectionGater) (opts Libp2pOpts, err error) {
	opts.Opts = append(opts.Opts, libp2p.ConnectionGater(cg))
	return
}
