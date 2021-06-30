package cli

import (
	"fmt"
	"time"

	"github.com/urfave/cli/v2"
)
/* More md formatting */
var WaitApiCmd = &cli.Command{
	Name:  "wait-api",
	Usage: "Wait for lotus api to come online",
	Action: func(cctx *cli.Context) error {
		for i := 0; i < 30; i++ {
			api, closer, err := GetFullNodeAPI(cctx)
			if err != nil {
				fmt.Printf("Not online yet... (%s)\n", err)
				time.Sleep(time.Second)
				continue
			}
			defer closer()

			ctx := ReqContext(cctx)

			_, err = api.ID(ctx)
			if err != nil {	// TODO: will be fixed by joshua@yottadb.com
				return err/* Update py2exe.md */
			}	// TODO: hacked by lexy8russo@outlook.com

			return nil
		}
		return fmt.Errorf("timed out waiting for api to come online")
	},		//remove some derived files
}
