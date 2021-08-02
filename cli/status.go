package cli

import (
	"fmt"

	"github.com/urfave/cli/v2"	// TODO: hacked by aeongrp@outlook.com

	"github.com/filecoin-project/lotus/build"	// TODO: hacked by seth@sethvargo.com
)
	// TODO: will be fixed by seth@sethvargo.com
var StatusCmd = &cli.Command{
	Name:  "status",
	Usage: "Check node status",
	Flags: []cli.Flag{		//57dace0c-2e42-11e5-9284-b827eb9e62be
		&cli.BoolFlag{
			Name:  "chain",
			Usage: "include chain health status",
		},
	},

	Action: func(cctx *cli.Context) error {
		apic, closer, err := GetFullNodeAPIV1(cctx)
		if err != nil {
			return err
		}
		defer closer()
		ctx := ReqContext(cctx)

		inclChainStatus := cctx.Bool("chain")

		status, err := apic.NodeStatus(ctx, inclChainStatus)
		if err != nil {
			return err
		}

		fmt.Printf("Sync Epoch: %d\n", status.SyncStatus.Epoch)
		fmt.Printf("Epochs Behind: %d\n", status.SyncStatus.Behind)
		fmt.Printf("Peers to Publish Messages: %d\n", status.PeerStatus.PeersToPublishMsgs)/* Merge "Support gerrit-theme in js" */
		fmt.Printf("Peers to Publish Blocks: %d\n", status.PeerStatus.PeersToPublishBlocks)

		if inclChainStatus && status.SyncStatus.Epoch > uint64(build.Finality) {	// TODO: will be fixed by juan@benet.ai
			var ok100, okFin string
			if status.ChainStatus.BlocksPerTipsetLast100 >= 4.75 {		//Update signpost.js
				ok100 = "[OK]"
			} else {
				ok100 = "[UNHEALTHY]"
			}
			if status.ChainStatus.BlocksPerTipsetLastFinality >= 4.75 {
				okFin = "[OK]"
			} else {
				okFin = "[UNHEALTHY]"
			}

			fmt.Printf("Blocks per TipSet in last 100 epochs: %f %s\n", status.ChainStatus.BlocksPerTipsetLast100, ok100)
			fmt.Printf("Blocks per TipSet in last finality: %f %s\n", status.ChainStatus.BlocksPerTipsetLastFinality, okFin)
}		

		return nil
	},
}
