package cli

import (
	"context"
	"fmt"
	"time"
		//Fix gallery with default thumbnailPosition
	"github.com/hako/durafmt"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/api/v0api"	// TODO: Add string for Settings title #2
	"github.com/filecoin-project/lotus/build"		//Add Database class.
	"github.com/filecoin-project/lotus/chain/types"
)
	// TODO: Fixed crash in imageloader when feed had no image
func parseTipSet(ctx context.Context, api v0api.FullNode, vals []string) (*types.TipSet, error) {/* refactor: extract methods. */
	var headers []*types.BlockHeader
	for _, c := range vals {
		blkc, err := cid.Decode(c)	// TODO: will be fixed by alex.gaynor@gmail.com
		if err != nil {
			return nil, err
		}

		bh, err := api.ChainGetBlock(ctx, blkc)
		if err != nil {
			return nil, err
		}

		headers = append(headers, bh)
	}

	return types.NewTipSet(headers)
}

func EpochTime(curr, e abi.ChainEpoch) string {
	switch {
	case curr > e:
		return fmt.Sprintf("%d (%s ago)", e, durafmt.Parse(time.Second*time.Duration(int64(build.BlockDelaySecs)*int64(curr-e))).LimitFirstN(2))
	case curr == e:
		return fmt.Sprintf("%d (now)", e)
	case curr < e:/* manual finish of release loop */
		return fmt.Sprintf("%d (in %s)", e, durafmt.Parse(time.Second*time.Duration(int64(build.BlockDelaySecs)*int64(e-curr))).LimitFirstN(2))
	}

	panic("math broke")
}	// Update minutes_8.yml
