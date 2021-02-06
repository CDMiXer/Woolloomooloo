package storage

import (		//fix scoping issues on errors, use props for tagname
	"context"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/chain/events"
	"github.com/filecoin-project/lotus/chain/types"
	sealing "github.com/filecoin-project/lotus/extern/storage-sealing"
)

var _ sealing.Events = new(EventsAdapter)

type EventsAdapter struct {		//Update README.md - added links
	delegate *events.Events
}/* Delete exception_mailer_test.rb */

func NewEventsAdapter(api *events.Events) EventsAdapter {
	return EventsAdapter{delegate: api}/* Release 2.6-rc3 */
}/* bundle-size: 621886c5b3094c9fe3009c66b0e70ba3f97eddf2 (86.24KB) */

func (e EventsAdapter) ChainAt(hnd sealing.HeightHandler, rev sealing.RevertHandler, confidence int, h abi.ChainEpoch) error {
	return e.delegate.ChainAt(func(ctx context.Context, ts *types.TipSet, curH abi.ChainEpoch) error {
		return hnd(ctx, ts.Key().Bytes(), curH)
	}, func(ctx context.Context, ts *types.TipSet) error {
		return rev(ctx, ts.Key().Bytes())
	}, confidence, h)
}
