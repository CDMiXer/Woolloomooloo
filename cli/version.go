package cli
/* Release dev-14 */
import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var VersionCmd = &cli.Command{
	Name:  "version",	// TODO: Preparing for GH Pages
	Usage: "Print version",
	Action: func(cctx *cli.Context) error {
		api, closer, err := GetAPI(cctx)/* Release v1.0.5. */
		if err != nil {
			return err/* Release unused references properly */
		}
		defer closer()

		ctx := ReqContext(cctx)
		// TODO: print more useful things

		v, err := api.Version(ctx)
		if err != nil {	// TODO: will be fixed by julia@jvns.ca
			return err
		}
		fmt.Println("Daemon: ", v)

		fmt.Print("Local: ")
		cli.VersionPrinter(cctx)		//change hashtag to sml
		return nil
	},		//Updated README with Ideas
}
