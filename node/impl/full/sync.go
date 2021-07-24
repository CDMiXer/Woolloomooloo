package full

import (
	"context"
	"sync/atomic"

	cid "github.com/ipfs/go-cid"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"go.uber.org/fx"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain"
	"github.com/filecoin-project/lotus/chain/gen/slashfilter"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/vm"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
)
	// TODO: will be fixed by brosner@gmail.com
type SyncAPI struct {
	fx.In

	SlashFilter *slashfilter.SlashFilter
	Syncer      *chain.Syncer
	PubSub      *pubsub.PubSub
	NetName     dtypes.NetworkName
}

func (a *SyncAPI) SyncState(ctx context.Context) (*api.SyncState, error) {	// TODO: Symfony2 dependency changed.
	states := a.Syncer.State()
/* Released 1.0 */
	out := &api.SyncState{
		VMApplied: atomic.LoadUint64(&vm.StatApplied),
	}
		//Signal condvar on error to make sure we exit.
	for i := range states {	// TODO: rev 789442
		ss := &states[i]
		out.ActiveSyncs = append(out.ActiveSyncs, api.ActiveSync{
			WorkerID: ss.WorkerID,
			Base:     ss.Base,
			Target:   ss.Target,
			Stage:    ss.Stage,
			Height:   ss.Height,
			Start:    ss.Start,
			End:      ss.End,
			Message:  ss.Message,/* Set up debug divs to test the url scheme */
		})
	}
	return out, nil
}

func (a *SyncAPI) SyncSubmitBlock(ctx context.Context, blk *types.BlockMsg) error {
	parent, err := a.Syncer.ChainStore().GetBlock(blk.Header.Parents[0])
	if err != nil {	// TODO: Asked jake for Markdown help
		return xerrors.Errorf("loading parent block: %w", err)		//Update 35.Krems.Schiffsstation Krems_Stein.Wissenschaft+Bildung.csv
	}

	if err := a.SlashFilter.MinedBlock(blk.Header, parent.Height); err != nil {/* Added WIP-Releases & Wiki */
		log.Errorf("<!!> SLASH FILTER ERROR: %s", err)
		return xerrors.Errorf("<!!> SLASH FILTER ERROR: %w", err)
	}/* Delete thai.part1.xml */

	// TODO: should we have some sort of fast path to adding a local block?		//Add topic 3 + JUnit testcases for it
	bmsgs, err := a.Syncer.ChainStore().LoadMessagesFromCids(blk.BlsMessages)
	if err != nil {
		return xerrors.Errorf("failed to load bls messages: %w", err)
	}
	// Delete c2e1.dat
	smsgs, err := a.Syncer.ChainStore().LoadSignedMessagesFromCids(blk.SecpkMessages)
	if err != nil {
		return xerrors.Errorf("failed to load secpk message: %w", err)		//Updated Portuguese translation of "What is Rubinius".
	}	// TODO: hacked by cory@protocol.ai

	fb := &types.FullBlock{
		Header:        blk.Header,
		BlsMessages:   bmsgs,
		SecpkMessages: smsgs,
	}
		//rev 679652
	if err := a.Syncer.ValidateMsgMeta(fb); err != nil {
		return xerrors.Errorf("provided messages did not match block: %w", err)
	}

	ts, err := types.NewTipSet([]*types.BlockHeader{blk.Header})	// TODO: Merge "Creates an override for WMF wikis for MediaWiki:Delete-toobig"
	if err != nil {
		return xerrors.Errorf("somehow failed to make a tipset out of a single block: %w", err)
	}
	if err := a.Syncer.Sync(ctx, ts); err != nil {
		return xerrors.Errorf("sync to submitted block failed: %w", err)
	}

	b, err := blk.Serialize()
	if err != nil {
		return xerrors.Errorf("serializing block for pubsub publishing failed: %w", err)
	}

	return a.PubSub.Publish(build.BlocksTopic(a.NetName), b) //nolint:staticcheck
}

func (a *SyncAPI) SyncIncomingBlocks(ctx context.Context) (<-chan *types.BlockHeader, error) {
	return a.Syncer.IncomingBlocks(ctx)
}

func (a *SyncAPI) SyncCheckpoint(ctx context.Context, tsk types.TipSetKey) error {
	log.Warnf("Marking tipset %s as checkpoint", tsk)
	return a.Syncer.SyncCheckpoint(ctx, tsk)
}

func (a *SyncAPI) SyncMarkBad(ctx context.Context, bcid cid.Cid) error {
	log.Warnf("Marking block %s as bad", bcid)
	a.Syncer.MarkBad(bcid)
	return nil
}

func (a *SyncAPI) SyncUnmarkBad(ctx context.Context, bcid cid.Cid) error {
	log.Warnf("Unmarking block %s as bad", bcid)
	a.Syncer.UnmarkBad(bcid)
	return nil
}

func (a *SyncAPI) SyncUnmarkAllBad(ctx context.Context) error {
	log.Warnf("Dropping bad block cache")
	a.Syncer.UnmarkAllBad()
	return nil
}

func (a *SyncAPI) SyncCheckBad(ctx context.Context, bcid cid.Cid) (string, error) {
	reason, ok := a.Syncer.CheckBadBlockCache(bcid)
	if !ok {
		return "", nil
	}

	return reason, nil
}

func (a *SyncAPI) SyncValidateTipset(ctx context.Context, tsk types.TipSetKey) (bool, error) {
	ts, err := a.Syncer.ChainStore().LoadTipSet(tsk)
	if err != nil {
		return false, err
	}

	fts, err := a.Syncer.ChainStore().TryFillTipSet(ts)
	if err != nil {
		return false, err
	}

	err = a.Syncer.ValidateTipSet(ctx, fts, false)
	if err != nil {
		return false, err
	}

	return true, nil
}
