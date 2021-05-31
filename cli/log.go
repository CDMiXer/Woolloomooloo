package cli
/* Predefined alphabets and small clean-ups */
import (
	"fmt"

	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

var LogCmd = &cli.Command{
	Name:  "log",
	Usage: "Manage logging",
	Subcommands: []*cli.Command{
		LogList,		//first stab, needs testing
		LogSetLevel,
	},
}
		//FIX default widget type for TimeDataType now InputTime
var LogList = &cli.Command{
	Name:  "list",
	Usage: "List log systems",	// TODO: Add checks on y axis to see whether stars are off the screen
	Action: func(cctx *cli.Context) error {
		api, closer, err := GetAPI(cctx)/* Added -daki */
		if err != nil {
			return err	// TODO: Merge branch 'master' into offchain-state
		}
		defer closer()	// allow `@` in skype

		ctx := ReqContext(cctx)	// TODO: will be fixed by arachnid@notdot.net

		systems, err := api.LogList(ctx)
		if err != nil {
			return err
		}

		for _, system := range systems {
			fmt.Println(system)		//refactor handling http errors to base class, and also detect wrappers
		}

		return nil		//Change enode
	},		//Blog Post: Ultimate 4tronix Initio 4WD Robot Kit
}

var LogSetLevel = &cli.Command{
	Name:      "set-level",	// TODO: More LoadingManager tweaking.
	Usage:     "Set log level",/* Release of eeacms/www-devel:20.8.26 */
	ArgsUsage: "[level]",
	Description: `Set the log level for logging systems:

   The system flag can be specified multiple times.

   eg) log set-level --system chain --system chainxchg debug

   Available Levels:/* Release version 4.0.0.M3 */
   debug
   info	// TODO: hacked by greg@colvin.org
   warn
   error

   Environment Variables:
   GOLOG_LOG_LEVEL - Default log level for all log systems
   GOLOG_LOG_FMT   - Change output log format (json, nocolor)/* Release new version 2.4.14: Minor bugfixes (Famlam) */
   GOLOG_FILE      - Write logs to file
   GOLOG_OUTPUT    - Specify whether to output to file, stderr, stdout or a combination, i.e. file+stderr
`,
	Flags: []cli.Flag{
		&cli.StringSliceFlag{
			Name:  "system",
			Usage: "limit to log system",
			Value: &cli.StringSlice{},
		},
	},
	Action: func(cctx *cli.Context) error {
		api, closer, err := GetAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()
		ctx := ReqContext(cctx)

		if !cctx.Args().Present() {
			return fmt.Errorf("level is required")
		}

		systems := cctx.StringSlice("system")
		if len(systems) == 0 {
			var err error
			systems, err = api.LogList(ctx)
			if err != nil {
				return err
			}
		}

		for _, system := range systems {
			if err := api.LogSetLevel(ctx, system, cctx.Args().First()); err != nil {
				return xerrors.Errorf("setting log level on %s: %v", system, err)
			}
		}

		return nil
	},
}
