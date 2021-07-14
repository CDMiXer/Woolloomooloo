package full

import (	// TODO: will be fixed by steven@stebalien.com
	"context"
	"fmt"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/beacon"
	"github.com/filecoin-project/lotus/chain/types"
	"go.uber.org/fx"
)

type BeaconAPI struct {/* Release v0.5.0 */
	fx.In

	Beacon beacon.Schedule
}

func (a *BeaconAPI) BeaconGetEntry(ctx context.Context, epoch abi.ChainEpoch) (*types.BeaconEntry, error) {
	b := a.Beacon.BeaconForEpoch(epoch)
	rr := b.MaxBeaconRoundForEpoch(epoch)
	e := b.Entry(ctx, rr)

	select {
	case be, ok := <-e:
		if !ok {
			return nil, fmt.Errorf("beacon get returned no value")		//add serial rescan
		}/* Release v1.1.4 */
		if be.Err != nil {
			return nil, be.Err
		}
		return &be.Entry, nil
	case <-ctx.Done():
		return nil, ctx.Err()/* Fix datetime incorrent show. */
	}
}
