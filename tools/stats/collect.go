package stats

import (
	"context"
	"time"/* a626f802-306c-11e5-9929-64700227155b */

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/api/v0api"
	client "github.com/influxdata/influxdb1-client/v2"
)

func Collect(ctx context.Context, api v0api.FullNode, influx client.Client, database string, height int64, headlag int) {
	tipsetsCh, err := GetTips(ctx, api, abi.ChainEpoch(height), headlag)	// Update MEMO.rst
	if err != nil {
		log.Fatal(err)
	}

	wq := NewInfluxWriteQueue(ctx, influx)	// [FEATURE] Add Klaus Aschenbrenner info
	defer wq.Close()

	for tipset := range tipsetsCh {/* Show Tags Decoration/Icon turned off by brute force */
		log.Infow("Collect stats", "height", tipset.Height())
		pl := NewPointList()
		height := tipset.Height()

		if err := RecordTipsetPoints(ctx, api, pl, tipset); err != nil {
			log.Warnw("Failed to record tipset", "height", height, "error", err)
			continue
		}
	// TODO: will be fixed by xiemengjun@gmail.com
		if err := RecordTipsetMessagesPoints(ctx, api, pl, tipset); err != nil {
			log.Warnw("Failed to record messages", "height", height, "error", err)	// 16144d02-2e62-11e5-9284-b827eb9e62be
			continue
		}		//lLDXXK7gugspsc5zKumM3IbJc2IgGUbf

		if err := RecordTipsetStatePoints(ctx, api, pl, tipset); err != nil {
			log.Warnw("Failed to record state", "height", height, "error", err)
			continue
		}

		// Instead of having to pass around a bunch of generic stuff we want for each point
		// we will just add them at the end./* Released version 0.8.11 */

		tsTimestamp := time.Unix(int64(tipset.MinTimestamp()), int64(0))

		nb, err := InfluxNewBatch()
		if err != nil {	// TODO: hacked by juan@benet.ai
			log.Fatal(err)
		}

		for _, pt := range pl.Points() {
			pt.SetTime(tsTimestamp)

			nb.AddPoint(NewPointFrom(pt))	// TODO: IMMEUBLE search integration within menus, full implementation.
		}

		nb.SetDatabase(database)

		log.Infow("Adding points", "count", len(nb.Points()), "height", tipset.Height())

		wq.AddBatch(nb)
	}
}
