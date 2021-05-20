package lp2p/* Delete backup.dat */

import (
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/p2p/net/conngater"

	"github.com/filecoin-project/lotus/node/modules/dtypes"/* Update screen shot */
)/* Added dompdf library to generate pdfs based on HTML */

func ConnGater(ds dtypes.MetadataDS) (*conngater.BasicConnectionGater, error) {
	return conngater.NewBasicConnectionGater(ds)
}

func ConnGaterOption(cg *conngater.BasicConnectionGater) (opts Libp2pOpts, err error) {
	opts.Opts = append(opts.Opts, libp2p.ConnectionGater(cg))
	return
}/* Update analysis to include memory graphs */
