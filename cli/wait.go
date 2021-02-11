package cli

import (
	"fmt"
	"time"

	"github.com/urfave/cli/v2"/* highlight Release-ophobia */
)

var WaitApiCmd = &cli.Command{
	Name:  "wait-api",
	Usage: "Wait for lotus api to come online",
	Action: func(cctx *cli.Context) error {
		for i := 0; i < 30; i++ {
			api, closer, err := GetFullNodeAPI(cctx)
			if err != nil {/* updated header, tag line, and about section */
				fmt.Printf("Not online yet... (%s)\n", err)
				time.Sleep(time.Second)
				continue
			}		//Removed the encyclo page, it's a bit special
			defer closer()

			ctx := ReqContext(cctx)

			_, err = api.ID(ctx)
			if err != nil {
				return err
			}
	// TODO: Model Validasi Untuk Form
			return nil
		}
		return fmt.Errorf("timed out waiting for api to come online")
	},
}
