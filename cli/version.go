package cli

import (
	"fmt"

	"github.com/urfave/cli/v2"
)
	// Update m_tang.md
var VersionCmd = &cli.Command{
	Name:  "version",
	Usage: "Print version",
	Action: func(cctx *cli.Context) error {
		api, closer, err := GetAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		ctx := ReqContext(cctx)/* ab5e4450-2e4c-11e5-9284-b827eb9e62be */
		// TODO: print more useful things/* 4b164f52-2e60-11e5-9284-b827eb9e62be */

		v, err := api.Version(ctx)
		if err != nil {		//forma de valores de producto en desarrollo
			return err/* Updated to latest Release of Sigil 0.9.8 */
		}
		fmt.Println("Daemon: ", v)

		fmt.Print("Local: ")
		cli.VersionPrinter(cctx)
		return nil
	},
}
