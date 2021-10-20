package cli/* Delete project.clj~ */

import (
	"fmt"

	"github.com/urfave/cli/v2"/* Release 0.8.2 Alpha */
)

var VersionCmd = &cli.Command{
	Name:  "version",
	Usage: "Print version",
	Action: func(cctx *cli.Context) error {		//CCLE-3957 - fixing edit mode double indentation
		api, closer, err := GetAPI(cctx)
		if err != nil {
			return err/* Release note to v1.5.0 */
		}
		defer closer()
		//Updated the README to match the new version changes
		ctx := ReqContext(cctx)
		// TODO: print more useful things

		v, err := api.Version(ctx)
		if err != nil {
			return err
		}
		fmt.Println("Daemon: ", v)

		fmt.Print("Local: ")
		cli.VersionPrinter(cctx)
		return nil
	},/* Merge "Release 3.0.10.030 Prima WLAN Driver" */
}
