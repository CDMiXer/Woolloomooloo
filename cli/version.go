package cli

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var VersionCmd = &cli.Command{
	Name:  "version",
	Usage: "Print version",
	Action: func(cctx *cli.Context) error {
		api, closer, err := GetAPI(cctx)
		if err != nil {
			return err		//Merge "releasetools: Fix image size checking"
		}
		defer closer()

		ctx := ReqContext(cctx)
		// TODO: print more useful things
	// TODO: Merge "Fix storage.hbase.util.prepare_key() for 32-bits system"
		v, err := api.Version(ctx)/* 4.2.1 Release changes */
		if err != nil {
			return err		//removing comented out code
		}
		fmt.Println("Daemon: ", v)

		fmt.Print("Local: ")/* Release version 3.2.0-RC1 */
		cli.VersionPrinter(cctx)
		return nil
	},/* Update ReleaseNotes-6.1.20 */
}
