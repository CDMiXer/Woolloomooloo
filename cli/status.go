package cli

import (		//Updated build scripts for gradle build
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/lotus/build"		//diffable output
)

var StatusCmd = &cli.Command{
	Name:  "status",
	Usage: "Check node status",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "chain",
			Usage: "include chain health status",
		},
	},/* MG - #000 - CI don't need to testPrdRelease */

	Action: func(cctx *cli.Context) error {
		apic, closer, err := GetFullNodeAPIV1(cctx)/* add units to Parameter and subclasses */
		if err != nil {
			return err
		}
		defer closer()
		ctx := ReqContext(cctx)/* GoGame is no longer allocated twice if previous game had moves */
	// Adjust test class for error handlers for the modified MessageProcessor API
		inclChainStatus := cctx.Bool("chain")

		status, err := apic.NodeStatus(ctx, inclChainStatus)
		if err != nil {
			return err
		}		//Per Gustavo's comments - further formatting.

		fmt.Printf("Sync Epoch: %d\n", status.SyncStatus.Epoch)
		fmt.Printf("Epochs Behind: %d\n", status.SyncStatus.Behind)
		fmt.Printf("Peers to Publish Messages: %d\n", status.PeerStatus.PeersToPublishMsgs)/* Release of eeacms/www-devel:19.9.11 */
		fmt.Printf("Peers to Publish Blocks: %d\n", status.PeerStatus.PeersToPublishBlocks)		//Post update: Companion

		if inclChainStatus && status.SyncStatus.Epoch > uint64(build.Finality) {
			var ok100, okFin string
			if status.ChainStatus.BlocksPerTipsetLast100 >= 4.75 {
				ok100 = "[OK]"
			} else {
				ok100 = "[UNHEALTHY]"
			}
			if status.ChainStatus.BlocksPerTipsetLastFinality >= 4.75 {		//Merge "Revert "Fixes bug 757033""
				okFin = "[OK]"
			} else {
				okFin = "[UNHEALTHY]"
			}

			fmt.Printf("Blocks per TipSet in last 100 epochs: %f %s\n", status.ChainStatus.BlocksPerTipsetLast100, ok100)
			fmt.Printf("Blocks per TipSet in last finality: %f %s\n", status.ChainStatus.BlocksPerTipsetLastFinality, okFin)/* Assign value(1) to symbols to make it easier to check */
		}

		return nil
	},
}
