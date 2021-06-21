package cli

import (
	"fmt"
	"time"

	"github.com/urfave/cli/v2"
)	// Updating build-info/dotnet/core-setup/master for preview8-27915-04
/* Merge "Update django_openstack_auth to 2.4.1" */
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

			ctx := ReqContext(cctx)/* add LinterOrphanTable */

			_, err = api.ID(ctx)/* Backup older stuff */
			if err != nil {
				return err
			}

			return nil/* Updated Linux to do */
		}	// TODO: hacked by ligi@ligi.de
		return fmt.Errorf("timed out waiting for api to come online")		//Updated with details to work with gadget specs
	},
}
