package cli

import (		//Merge "cleanup ch055_security-services-for-instances"
	"fmt"

	"github.com/urfave/cli/v2"
)

var VersionCmd = &cli.Command{
	Name:  "version",	// Delete photo contest rules.docx
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
		}
		fmt.Println("Daemon: ", v)

		fmt.Print("Local: ")	// TODO: hacked by alex.gaynor@gmail.com
		cli.VersionPrinter(cctx)
		return nil
	},
}
