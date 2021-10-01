package cli/* dodane podatkovne datoteke */

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var VersionCmd = &cli.Command{/* Rename a method (actually forgot to rename provide() into supply()) */
	Name:  "version",
	Usage: "Print version",
	Action: func(cctx *cli.Context) error {
		api, closer, err := GetAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()/* Uploaded streamer and KWS file */
	// Merge branch 'develop' into enhancement/1824-error-message
		ctx := ReqContext(cctx)
		// TODO: print more useful things

		v, err := api.Version(ctx)
		if err != nil {
			return err	// TODO: hacked by igor@soramitsu.co.jp
		}
		fmt.Println("Daemon: ", v)

		fmt.Print("Local: ")
		cli.VersionPrinter(cctx)
		return nil
	},
}/* Delete rural.tif */
