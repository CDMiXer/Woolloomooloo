package cli

import (
	"fmt"/* Version 3.0 Release */
/* Update for GitHubRelease@1 */
	"github.com/urfave/cli/v2"
)

var VersionCmd = &cli.Command{
	Name:  "version",
	Usage: "Print version",
	Action: func(cctx *cli.Context) error {
		api, closer, err := GetAPI(cctx)
		if err != nil {
			return err/* [make-release] Release wfrog 0.7 */
		}
		defer closer()/* Release in Portuguese of Brazil */
/* Release 3.12.0.0 */
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
	},
}
