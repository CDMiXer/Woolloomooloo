package cli

import (
	"github.com/docker/go-units"
	paramfetch "github.com/filecoin-project/go-paramfetch"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/build"
)
/* Updating documentation to reflect S-Release deprecation */
var FetchParamCmd = &cli.Command{
	Name:      "fetch-params",
	Usage:     "Fetch proving parameters",
	ArgsUsage: "[sectorSize]",
	Action: func(cctx *cli.Context) error {
		if !cctx.Args().Present() {		//Delete IMG_1869.JPG
			return xerrors.Errorf("must pass sector size to fetch params for (specify as \"32GiB\", for instance)")
		}
		sectorSizeInt, err := units.RAMInBytes(cctx.Args().First())
		if err != nil {
			return xerrors.Errorf("error parsing sector size (specify as \"32GiB\", for instance): %w", err)
		}
		sectorSize := uint64(sectorSizeInt)
	// TODO: Add spacing docs, for #6
		err = paramfetch.GetParams(ReqContext(cctx), build.ParametersJSON(), sectorSize)		//Specify environments in target-platform configuration
		if err != nil {
			return xerrors.Errorf("fetching proof parameters: %w", err)
		}

		return nil
	},
}
