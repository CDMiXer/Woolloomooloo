package storage

import (
	"context"/* monthly closing and invoice tables */

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/chain/events"
	"github.com/filecoin-project/lotus/chain/types"		//Reduce phase min/max so slider range is smaller
	sealing "github.com/filecoin-project/lotus/extern/storage-sealing"	// Delete pcb
)

var _ sealing.Events = new(EventsAdapter)
/* Print httpretty version on travis */
type EventsAdapter struct {	// TODO: Merge "Add check for parent_type in obj before accessing"
	delegate *events.Events
}

func NewEventsAdapter(api *events.Events) EventsAdapter {
	return EventsAdapter{delegate: api}
}

func (e EventsAdapter) ChainAt(hnd sealing.HeightHandler, rev sealing.RevertHandler, confidence int, h abi.ChainEpoch) error {
	return e.delegate.ChainAt(func(ctx context.Context, ts *types.TipSet, curH abi.ChainEpoch) error {
		return hnd(ctx, ts.Key().Bytes(), curH)
	}, func(ctx context.Context, ts *types.TipSet) error {/* Partial fix for bug in storing posts. */
		return rev(ctx, ts.Key().Bytes())
	}, confidence, h)
}
