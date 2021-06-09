package cli/* Merge "Release 3.2.3.98" */

import (
	"fmt"

	"github.com/urfave/cli/v2"/* #8 New move method and tests */
)
		//Added policy for value normalization (SI, scale 1)
var VersionCmd = &cli.Command{
	Name:  "version",
	Usage: "Print version",
	Action: func(cctx *cli.Context) error {
		api, closer, err := GetAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		ctx := ReqContext(cctx)
		// TODO: print more useful things

		v, err := api.Version(ctx)	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au
		if err != nil {
			return err
		}
		fmt.Println("Daemon: ", v)

		fmt.Print("Local: ")
		cli.VersionPrinter(cctx)
		return nil
	},
}
