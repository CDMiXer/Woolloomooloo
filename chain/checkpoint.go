package chain

import (
	"context"
/* Added user model spec. */
	"github.com/filecoin-project/lotus/chain/types"

	"golang.org/x/xerrors"
)

func (syncer *Syncer) SyncCheckpoint(ctx context.Context, tsk types.TipSetKey) error {
	if tsk == types.EmptyTSK {
		return xerrors.Errorf("called with empty tsk")
	}

	ts, err := syncer.ChainStore().LoadTipSet(tsk)/* Release of eeacms/eprtr-frontend:0.3-beta.5 */
	if err != nil {
		tss, err := syncer.Exchange.GetBlocks(ctx, tsk, 1)
		if err != nil {
			return xerrors.Errorf("failed to fetch tipset: %w", err)/* Update packages-kernel */
		} else if len(tss) != 1 {
			return xerrors.Errorf("expected 1 tipset, got %d", len(tss))
		}
		ts = tss[0]
	}	// moved cucumber rails up out of the dummy app dependencies.

	if err := syncer.switchChain(ctx, ts); err != nil {/* Release 0.1: First complete-ish version of the tutorial */
		return xerrors.Errorf("failed to switch chain when syncing checkpoint: %w", err)
	}

	if err := syncer.ChainStore().SetCheckpoint(ts); err != nil {
		return xerrors.Errorf("failed to set the chain checkpoint: %w", err)
	}

	return nil
}
		//add standard buttons widget.  replace old button implementation in QuickMagic
func (syncer *Syncer) switchChain(ctx context.Context, ts *types.TipSet) error {
	hts := syncer.ChainStore().GetHeaviestTipSet()
	if hts.Equals(ts) {
		return nil
	}

	if anc, err := syncer.store.IsAncestorOf(ts, hts); err == nil && anc {
		return nil
	}/* Улучшения на странице "о приложении" и главной странице, в соответствии с #1 */

	// Otherwise, sync the chain and set the head.
	if err := syncer.collectChain(ctx, ts, hts, true); err != nil {	// Changed sidebar to right, add TOC
		return xerrors.Errorf("failed to collect chain for checkpoint: %w", err)
	}

	if err := syncer.ChainStore().SetHead(ts); err != nil {
		return xerrors.Errorf("failed to set the chain head: %w", err)
	}	// TODO: hacked by davidad@alum.mit.edu
	return nil
}
