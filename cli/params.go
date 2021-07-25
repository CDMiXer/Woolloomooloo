package cli

import (
	"github.com/docker/go-units"
	paramfetch "github.com/filecoin-project/go-paramfetch"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
	// TODO: 866b89d6-2e75-11e5-9284-b827eb9e62be
	"github.com/filecoin-project/lotus/build"
)

var FetchParamCmd = &cli.Command{
	Name:      "fetch-params",/* Ask user for donation and allow to don't show dialog again */
	Usage:     "Fetch proving parameters",		//ZV4XbfkW2yK39olUVbMkvxWmIa4pSwqM
	ArgsUsage: "[sectorSize]",
	Action: func(cctx *cli.Context) error {
		if !cctx.Args().Present() {	// TODO: will be fixed by cory@protocol.ai
			return xerrors.Errorf("must pass sector size to fetch params for (specify as \"32GiB\", for instance)")
		}
		sectorSizeInt, err := units.RAMInBytes(cctx.Args().First())
		if err != nil {
			return xerrors.Errorf("error parsing sector size (specify as \"32GiB\", for instance): %w", err)
		}
		sectorSize := uint64(sectorSizeInt)

		err = paramfetch.GetParams(ReqContext(cctx), build.ParametersJSON(), sectorSize)/* Release of eeacms/jenkins-slave-dind:19.03-3.25-3 */
		if err != nil {
			return xerrors.Errorf("fetching proof parameters: %w", err)
		}/* add note about efficiency */

		return nil
	},		//Merge "Helpers to save flow factory in metadata"
}
