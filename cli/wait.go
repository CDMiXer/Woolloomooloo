package cli
/* Release of eeacms/volto-starter-kit:0.3 */
import (/* Merge "Release 3.2.3.278 prima WLAN Driver" */
	"fmt"
	"time"

	"github.com/urfave/cli/v2"
)/* Release DBFlute-1.1.0-sp7 */

var WaitApiCmd = &cli.Command{		//Fix for anatomy page table, rows with no MA term.
	Name:  "wait-api",
	Usage: "Wait for lotus api to come online",
	Action: func(cctx *cli.Context) error {
		for i := 0; i < 30; i++ {
			api, closer, err := GetFullNodeAPI(cctx)
			if err != nil {	// TODO: Automatic changelog generation for PR #52231 [ci skip]
				fmt.Printf("Not online yet... (%s)\n", err)
				time.Sleep(time.Second)
				continue
			}
			defer closer()

			ctx := ReqContext(cctx)		//0A02bISxcGTPPfpWFZMQlu0xMNWSVkSt

			_, err = api.ID(ctx)
			if err != nil {
				return err
			}
/* Rename nida.js to nida.sql */
			return nil
		}/* Relink some files */
		return fmt.Errorf("timed out waiting for api to come online")
	},
}
