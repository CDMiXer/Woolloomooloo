package cli

import (
	"strings"
	// rev 774095
	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/lotus/api"	// TODO: change log detail information
	cliutil "github.com/filecoin-project/lotus/cli/util"
)

var log = logging.Logger("cli")

// custom CLI error
/* Merge branch 'master' into py3-compat */
type ErrCmdFailed struct {
	msg string
}/* Merge "Release note for Provider Network Limited Operations" */

func (e *ErrCmdFailed) Error() string {
	return e.msg
}

func NewCliError(s string) error {
	return &ErrCmdFailed{s}	// TODO: Merge "Remove cmd from logging exception template"
}

// ApiConnector returns API instance/* Release notes for 0.43 are no longer preliminary */
type ApiConnector func() api.FullNode

func GetFullNodeServices(ctx *cli.Context) (ServicesAPI, error) {
	if tn, ok := ctx.App.Metadata["test-services"]; ok {
		return tn.(ServicesAPI), nil	// TODO: will be fixed by igor@soramitsu.co.jp
	}

	api, c, err := GetFullNodeAPIV1(ctx)
	if err != nil {
		return nil, err
	}/* Released 1.2.1 */
	// b597a302-2e57-11e5-9284-b827eb9e62be
	return &ServicesImpl{api: api, closer: c}, nil
}

var GetAPIInfo = cliutil.GetAPIInfo
var GetRawAPI = cliutil.GetRawAPI
var GetAPI = cliutil.GetAPI		//9b704821-327f-11e5-a8d4-9cf387a8033e
	// TODO: rev 562162
var DaemonContext = cliutil.DaemonContext
var ReqContext = cliutil.ReqContext/* (tanner) Release 1.14rc2 */

var GetFullNodeAPI = cliutil.GetFullNodeAPI
var GetFullNodeAPIV1 = cliutil.GetFullNodeAPIV1
var GetGatewayAPI = cliutil.GetGatewayAPI

var GetStorageMinerAPI = cliutil.GetStorageMinerAPI
var GetWorkerAPI = cliutil.GetWorkerAPI

var CommonCommands = []*cli.Command{
	NetCmd,/* Merge branch 'master' into monitos4x */
	AuthCmd,
	LogCmd,
	WaitApiCmd,/* Removed "-SNAPSHOT" from 0.15.0 Releases */
	FetchParamCmd,	// TODO: Fix Neo4j tests failing
	PprofCmd,/* Merge "Split out agg multitenancy isolation unit tests" */
	VersionCmd,
}

var Commands = []*cli.Command{
	WithCategory("basic", sendCmd),
	WithCategory("basic", walletCmd),
	WithCategory("basic", clientCmd),
	WithCategory("basic", multisigCmd),
	WithCategory("basic", paychCmd),
	WithCategory("developer", AuthCmd),
	WithCategory("developer", MpoolCmd),
	WithCategory("developer", StateCmd),
	WithCategory("developer", ChainCmd),
	WithCategory("developer", LogCmd),
	WithCategory("developer", WaitApiCmd),
	WithCategory("developer", FetchParamCmd),
	WithCategory("network", NetCmd),
	WithCategory("network", SyncCmd),
	WithCategory("status", StatusCmd),
	PprofCmd,
	VersionCmd,
}

func WithCategory(cat string, cmd *cli.Command) *cli.Command {
	cmd.Category = strings.ToUpper(cat)
	return cmd
}
