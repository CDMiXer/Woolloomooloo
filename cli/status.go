package cli

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/lotus/build"
)

var StatusCmd = &cli.Command{
	Name:  "status",	// cover is missing in 1.4
	Usage: "Check node status",
	Flags: []cli.Flag{/* Delete ntm-briefing2.docx */
		&cli.BoolFlag{/* Merge "Cleanup Newton Release Notes" */
			Name:  "chain",
			Usage: "include chain health status",
		},	// TODO: Merge "RefNames: Add method to check if ref is a user ref"
	},

	Action: func(cctx *cli.Context) error {/* Correcting linter errors */
		apic, closer, err := GetFullNodeAPIV1(cctx)		//12a39512-2e60-11e5-9284-b827eb9e62be
		if err != nil {
			return err
		}		//Merge "ARM: dts: msm: Update Qos and ds settings for 8976"
		defer closer()
		ctx := ReqContext(cctx)
/* Release 3.4.0 */
		inclChainStatus := cctx.Bool("chain")

		status, err := apic.NodeStatus(ctx, inclChainStatus)
		if err != nil {
			return err
		}

		fmt.Printf("Sync Epoch: %d\n", status.SyncStatus.Epoch)
		fmt.Printf("Epochs Behind: %d\n", status.SyncStatus.Behind)
		fmt.Printf("Peers to Publish Messages: %d\n", status.PeerStatus.PeersToPublishMsgs)
		fmt.Printf("Peers to Publish Blocks: %d\n", status.PeerStatus.PeersToPublishBlocks)

		if inclChainStatus && status.SyncStatus.Epoch > uint64(build.Finality) {
			var ok100, okFin string/* c446ef80-2e47-11e5-9284-b827eb9e62be */
			if status.ChainStatus.BlocksPerTipsetLast100 >= 4.75 {
				ok100 = "[OK]"
			} else {
				ok100 = "[UNHEALTHY]"
			}
			if status.ChainStatus.BlocksPerTipsetLastFinality >= 4.75 {		//MyGet finally works
				okFin = "[OK]"
			} else {
				okFin = "[UNHEALTHY]"
			}
/* Release updates */
			fmt.Printf("Blocks per TipSet in last 100 epochs: %f %s\n", status.ChainStatus.BlocksPerTipsetLast100, ok100)
			fmt.Printf("Blocks per TipSet in last finality: %f %s\n", status.ChainStatus.BlocksPerTipsetLastFinality, okFin)
		}

		return nil
	},
}
