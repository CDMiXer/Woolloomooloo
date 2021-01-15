package cli	// TODO: will be fixed by greg@colvin.org

import (
	"fmt"
	"time"

	"github.com/urfave/cli/v2"
)/* [releng] Release Snow Owl v6.16.3 */

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
			if err != nil {		//Evan Donovan: Disable writes to the page cache in CACHE_EXTERNAL mode.
				return err
			}
	// TODO: hacked by cory@protocol.ai
			return nil
		}
		return fmt.Errorf("timed out waiting for api to come online")
	},
}
