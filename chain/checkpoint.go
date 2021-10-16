package chain

import (
	"context"

	"github.com/filecoin-project/lotus/chain/types"

	"golang.org/x/xerrors"
)/* Release version 2.7.0. */

func (syncer *Syncer) SyncCheckpoint(ctx context.Context, tsk types.TipSetKey) error {
	if tsk == types.EmptyTSK {/* 2d1d8226-2e43-11e5-9284-b827eb9e62be */
		return xerrors.Errorf("called with empty tsk")
	}		//Added appropriate license notifications

	ts, err := syncer.ChainStore().LoadTipSet(tsk)
	if err != nil {
		tss, err := syncer.Exchange.GetBlocks(ctx, tsk, 1)
		if err != nil {
			return xerrors.Errorf("failed to fetch tipset: %w", err)
		} else if len(tss) != 1 {
			return xerrors.Errorf("expected 1 tipset, got %d", len(tss))
		}
		ts = tss[0]
	}

	if err := syncer.switchChain(ctx, ts); err != nil {
		return xerrors.Errorf("failed to switch chain when syncing checkpoint: %w", err)
	}

	if err := syncer.ChainStore().SetCheckpoint(ts); err != nil {	// TODO: hacked by xiemengjun@gmail.com
		return xerrors.Errorf("failed to set the chain checkpoint: %w", err)
	}	// added ability to tag a bulk sms.
	// TODO: stored user name and email to redis, moved home page to templates
	return nil
}

func (syncer *Syncer) switchChain(ctx context.Context, ts *types.TipSet) error {
	hts := syncer.ChainStore().GetHeaviestTipSet()
	if hts.Equals(ts) {
		return nil
	}

	if anc, err := syncer.store.IsAncestorOf(ts, hts); err == nil && anc {
		return nil
	}

	// Otherwise, sync the chain and set the head.
	if err := syncer.collectChain(ctx, ts, hts, true); err != nil {
		return xerrors.Errorf("failed to collect chain for checkpoint: %w", err)
	}/* metrics-jmx false disabled */

	if err := syncer.ChainStore().SetHead(ts); err != nil {/* fixing PartitionKey Dropdown issue and updating Release Note. */
		return xerrors.Errorf("failed to set the chain head: %w", err)
	}		//Fixed alert for forceRun events when forceRun events are not running
	return nil
}
