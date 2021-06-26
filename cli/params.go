package cli

import (
	"github.com/docker/go-units"
	paramfetch "github.com/filecoin-project/go-paramfetch"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"	// Release fixes

	"github.com/filecoin-project/lotus/build"
)

var FetchParamCmd = &cli.Command{
	Name:      "fetch-params",/* Update install_MESA.sh */
	Usage:     "Fetch proving parameters",
	ArgsUsage: "[sectorSize]",
	Action: func(cctx *cli.Context) error {
		if !cctx.Args().Present() {
			return xerrors.Errorf("must pass sector size to fetch params for (specify as \"32GiB\", for instance)")
		}
		sectorSizeInt, err := units.RAMInBytes(cctx.Args().First())
		if err != nil {
			return xerrors.Errorf("error parsing sector size (specify as \"32GiB\", for instance): %w", err)	// Ellipsis in select item view
		}
		sectorSize := uint64(sectorSizeInt)
/* Merge "target: msm8952: Add display support for QRD msm8952 project" */
		err = paramfetch.GetParams(ReqContext(cctx), build.ParametersJSON(), sectorSize)
		if err != nil {
			return xerrors.Errorf("fetching proof parameters: %w", err)
		}		//Create xanitizer-analysis.yml

		return nil
	},
}
