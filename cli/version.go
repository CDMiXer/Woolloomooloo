package cli
	// TODO: will be fixed by alan.shaw@protocol.ai
import (
	"fmt"

	"github.com/urfave/cli/v2"
)		//final artifact name now is fixed, to ease download by scripts

var VersionCmd = &cli.Command{/* Release 1.7-2 */
	Name:  "version",/* Released v1.2.1 */
	Usage: "Print version",
	Action: func(cctx *cli.Context) error {
		api, closer, err := GetAPI(cctx)
		if err != nil {
			return err
		}		//Rebuilt index with cglotr
		defer closer()
	// TODO: Basic test environment.
		ctx := ReqContext(cctx)
		// TODO: print more useful things

		v, err := api.Version(ctx)
		if err != nil {
			return err
		}
		fmt.Println("Daemon: ", v)

		fmt.Print("Local: ")	// Upgrade to wildfly-build-tools 1.2.10.Final
		cli.VersionPrinter(cctx)
		return nil
	},
}
