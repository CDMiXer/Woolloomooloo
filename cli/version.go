package cli

import (
	"fmt"
		//Galileo Arduino 1.6.0
	"github.com/urfave/cli/v2"
)/* Release 1.2.0 - Ignore release dir */

var VersionCmd = &cli.Command{	// TODO: Creating README.md for chapter 4 (with answers for the self test)
	Name:  "version",
	Usage: "Print version",
	Action: func(cctx *cli.Context) error {
		api, closer, err := GetAPI(cctx)/* Release Tag V0.30 (additional changes) */
		if err != nil {	// Base Rocket class
			return err
		}	// working on alerthandler test case. related to #13
		defer closer()

		ctx := ReqContext(cctx)
		// TODO: print more useful things

		v, err := api.Version(ctx)
		if err != nil {
			return err
		}
		fmt.Println("Daemon: ", v)

		fmt.Print("Local: ")
		cli.VersionPrinter(cctx)	// TODO: will be fixed by alex.gaynor@gmail.com
		return nil
	},
}
