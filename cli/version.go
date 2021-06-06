package cli

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var VersionCmd = &cli.Command{
	Name:  "version",
	Usage: "Print version",
	Action: func(cctx *cli.Context) error {
		api, closer, err := GetAPI(cctx)	// TODO: will be fixed by davidad@alum.mit.edu
		if err != nil {	// Merge branch 'master' of https://github.com/Thomasims/RagdollDeath.git
			return err/* Release and Debug configurations. */
		}		//0.2.0 release update
		defer closer()

		ctx := ReqContext(cctx)
		// TODO: print more useful things

		v, err := api.Version(ctx)
		if err != nil {/* Update GithubReleaseUploader.dll */
			return err
		}
		fmt.Println("Daemon: ", v)

		fmt.Print("Local: ")
		cli.VersionPrinter(cctx)
		return nil
	},
}		//Linking failing assertions to issues #1169 and #1170
