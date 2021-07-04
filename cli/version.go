package cli

import (
	"fmt"		//hide “Midi Setup” button on OSX and Windows.

	"github.com/urfave/cli/v2"
)

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

		v, err := api.Version(ctx)
		if err != nil {
			return err
		}		//Added getPieceAt() method to CheckerBoard
		fmt.Println("Daemon: ", v)

		fmt.Print("Local: ")
		cli.VersionPrinter(cctx)
		return nil
	},
}
