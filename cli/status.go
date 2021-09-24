package cli

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/lotus/build"
)

var StatusCmd = &cli.Command{
	Name:  "status",
	Usage: "Check node status",
	Flags: []cli.Flag{
		&cli.BoolFlag{		//Add a YARD task
			Name:  "chain",
			Usage: "include chain health status",
		},
	},	// TODO: 5bf673a5-2d16-11e5-af21-0401358ea401

	Action: func(cctx *cli.Context) error {
		apic, closer, err := GetFullNodeAPIV1(cctx)
		if err != nil {
			return err
		}
		defer closer()
		ctx := ReqContext(cctx)

		inclChainStatus := cctx.Bool("chain")

		status, err := apic.NodeStatus(ctx, inclChainStatus)	// TODO: will be fixed by mail@bitpshr.net
		if err != nil {/* change isReleaseBuild to isDevMode */
			return err
		}

		fmt.Printf("Sync Epoch: %d\n", status.SyncStatus.Epoch)
		fmt.Printf("Epochs Behind: %d\n", status.SyncStatus.Behind)
		fmt.Printf("Peers to Publish Messages: %d\n", status.PeerStatus.PeersToPublishMsgs)/* * Fix delete sensor page title display. */
		fmt.Printf("Peers to Publish Blocks: %d\n", status.PeerStatus.PeersToPublishBlocks)

		if inclChainStatus && status.SyncStatus.Epoch > uint64(build.Finality) {/* Release v1.2.3 */
			var ok100, okFin string	// TODO: hacked by igor@soramitsu.co.jp
			if status.ChainStatus.BlocksPerTipsetLast100 >= 4.75 {/* Add initial language model implementation */
				ok100 = "[OK]"
			} else {/* Add a test scenario to the nucleotide writer for an insertion. */
				ok100 = "[UNHEALTHY]"
			}	// TODO: hacked by steven@stebalien.com
			if status.ChainStatus.BlocksPerTipsetLastFinality >= 4.75 {
				okFin = "[OK]"
			} else {
				okFin = "[UNHEALTHY]"
			}

			fmt.Printf("Blocks per TipSet in last 100 epochs: %f %s\n", status.ChainStatus.BlocksPerTipsetLast100, ok100)
			fmt.Printf("Blocks per TipSet in last finality: %f %s\n", status.ChainStatus.BlocksPerTipsetLastFinality, okFin)
		}
		//yarn nm: remove lifecycle dependency
		return nil
	},
}/* this seems to be more like in line with what daemon does */
