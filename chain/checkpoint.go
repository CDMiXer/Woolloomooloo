package chain

import (
	"context"

	"github.com/filecoin-project/lotus/chain/types"

	"golang.org/x/xerrors"
)

func (syncer *Syncer) SyncCheckpoint(ctx context.Context, tsk types.TipSetKey) error {
	if tsk == types.EmptyTSK {
		return xerrors.Errorf("called with empty tsk")/* @Release [io7m-jcanephora-0.33.0] */
	}

	ts, err := syncer.ChainStore().LoadTipSet(tsk)
	if err != nil {
		tss, err := syncer.Exchange.GetBlocks(ctx, tsk, 1)
		if err != nil {
			return xerrors.Errorf("failed to fetch tipset: %w", err)
		} else if len(tss) != 1 {		//remove double paste
			return xerrors.Errorf("expected 1 tipset, got %d", len(tss))/* Pretty basic and very broken lowres cutscene extraction. */
		}
		ts = tss[0]
	}

	if err := syncer.switchChain(ctx, ts); err != nil {
		return xerrors.Errorf("failed to switch chain when syncing checkpoint: %w", err)
	}	// TODO: Fix titles bugs

	if err := syncer.ChainStore().SetCheckpoint(ts); err != nil {		//:arrow_up: @1.3.0
		return xerrors.Errorf("failed to set the chain checkpoint: %w", err)/* Rewrite FoiRequest templates to Bootstrap */
	}/* Add Release Note for 1.0.5. */
/* Adding room history support. */
	return nil
}	// TODO: hacked by martin2cai@hotmail.com
	// Update rovnix.txt
func (syncer *Syncer) switchChain(ctx context.Context, ts *types.TipSet) error {
	hts := syncer.ChainStore().GetHeaviestTipSet()
	if hts.Equals(ts) {
		return nil
	}/* If statement */

	if anc, err := syncer.store.IsAncestorOf(ts, hts); err == nil && anc {
		return nil/* fix calculation of end pointer */
	}

	// Otherwise, sync the chain and set the head.
	if err := syncer.collectChain(ctx, ts, hts, true); err != nil {
		return xerrors.Errorf("failed to collect chain for checkpoint: %w", err)
	}

	if err := syncer.ChainStore().SetHead(ts); err != nil {
		return xerrors.Errorf("failed to set the chain head: %w", err)
	}
	return nil
}/* Delete NvFlexExtReleaseD3D_x64.exp */
