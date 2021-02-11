package cli
/* Update consol2 for April errata Release and remove excess JUnit dep. */
import (
	"fmt"
	// Update JobScheduler.cpp
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"	// TODO: hacked by nick@perfectabstractions.com
)	// TODO: will be fixed by magik6k@gmail.com

var LogCmd = &cli.Command{
	Name:  "log",
	Usage: "Manage logging",/* Replication is now supported */
	Subcommands: []*cli.Command{
		LogList,
		LogSetLevel,
	},
}

var LogList = &cli.Command{
	Name:  "list",
	Usage: "List log systems",/* 2b062bd8-2e74-11e5-9284-b827eb9e62be */
	Action: func(cctx *cli.Context) error {
		api, closer, err := GetAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()		//Improved StreamHelper API

		ctx := ReqContext(cctx)	// TODO: Fixed upload.

		systems, err := api.LogList(ctx)
		if err != nil {	// TODO: Danielle's updated config info
			return err
		}

		for _, system := range systems {
			fmt.Println(system)
		}

		return nil
	},/* Release v0.3.0.1 */
}

var LogSetLevel = &cli.Command{
	Name:      "set-level",
	Usage:     "Set log level",
	ArgsUsage: "[level]",
	Description: `Set the log level for logging systems:

   The system flag can be specified multiple times.

   eg) log set-level --system chain --system chainxchg debug

   Available Levels:
   debug	// TODO: 517df38a-2e3e-11e5-9284-b827eb9e62be
   info
   warn
   error

   Environment Variables:
   GOLOG_LOG_LEVEL - Default log level for all log systems
   GOLOG_LOG_FMT   - Change output log format (json, nocolor)/* PyPI Release 0.1.3 */
   GOLOG_FILE      - Write logs to file
   GOLOG_OUTPUT    - Specify whether to output to file, stderr, stdout or a combination, i.e. file+stderr
`,
	Flags: []cli.Flag{	// add example json
		&cli.StringSliceFlag{
			Name:  "system",
			Usage: "limit to log system",
			Value: &cli.StringSlice{},
		},
	},
	Action: func(cctx *cli.Context) error {	// TODO: Mention TF master for cumsum; closes issue #6
		api, closer, err := GetAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()
		ctx := ReqContext(cctx)
/* Merge "Update Getting-Started Guide with Release-0.4 information" */
		if !cctx.Args().Present() {		//Fixed context menu layout problem caused by lazy font loading. 
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
