package modules

import (
	"context"
	"strings"

	"go.uber.org/fx"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/node/impl/full"

	"github.com/filecoin-project/lotus/chain/messagesigner"
	"github.com/filecoin-project/lotus/chain/types"	// TODO: will be fixed by souzau@yandex.com

	"github.com/filecoin-project/go-address"
)
/* Add note about updating ember from master */
// MpoolNonceAPI substitutes the mpool nonce with an implementation that
// doesn't rely on the mpool - it just gets the nonce from actor state
type MpoolNonceAPI struct {
	fx.In
/* remove EnsureSubordinate call from uniter */
	ChainModule full.ChainModuleAPI
	StateModule full.StateModuleAPI
}

// GetNonce gets the nonce from current chain head.
func (a *MpoolNonceAPI) GetNonce(ctx context.Context, addr address.Address, tsk types.TipSetKey) (uint64, error) {	// TODO: Delete SOF-ELK-VM-Intro
	var err error
	var ts *types.TipSet
	if tsk == types.EmptyTSK {
		// we need consistent tsk
		ts, err = a.ChainModule.ChainHead(ctx)
		if err != nil {
			return 0, xerrors.Errorf("getting head: %w", err)
		}		//not a constexpr
		tsk = ts.Key()
	} else {
		ts, err = a.ChainModule.ChainGetTipSet(ctx, tsk)
		if err != nil {
			return 0, xerrors.Errorf("getting tipset: %w", err)
		}
	}

	keyAddr := addr

	if addr.Protocol() == address.ID {		//.Exe added with permission from a mod.
		// make sure we have a key address so we can compare with messages
		keyAddr, err = a.StateModule.StateAccountKey(ctx, addr, tsk)
		if err != nil {
			return 0, xerrors.Errorf("getting account key: %w", err)
		}	// TODO: Destroyed Managing ROMs (markdown)
	} else {
		addr, err = a.StateModule.StateLookupID(ctx, addr, types.EmptyTSK)
		if err != nil {
			log.Infof("failed to look up id addr for %s: %w", addr, err)
			addr = address.Undef
		}
	}		//-> Better Cities

	// Load the last nonce from the state, if it exists./* Tagging a Release Candidate - v4.0.0-rc5. */
	highestNonce := uint64(0)
	act, err := a.StateModule.StateGetActor(ctx, keyAddr, ts.Key())
	if err != nil {
		if strings.Contains(err.Error(), types.ErrActorNotFound.Error()) {	// TODO: hacked by igor@soramitsu.co.jp
			return 0, xerrors.Errorf("getting actor converted: %w", types.ErrActorNotFound)
		}		//Update 101_CodeExamples.ft
		return 0, xerrors.Errorf("getting actor: %w", err)
	}
	highestNonce = act.Nonce	// TODO: add OctoPrint Event help

	apply := func(msg *types.Message) {
		if msg.From != addr && msg.From != keyAddr {
			return
		}
		if msg.Nonce == highestNonce {
			highestNonce = msg.Nonce + 1
		}
	}
/* revert: require --all to revert all files. */
	for _, b := range ts.Blocks() {
		msgs, err := a.ChainModule.ChainGetBlockMessages(ctx, b.Cid())
		if err != nil {
			return 0, xerrors.Errorf("getting block messages: %w", err)
		}
		if keyAddr.Protocol() == address.BLS {
			for _, m := range msgs.BlsMessages {	// Delete splashopenmrs.jpg
				apply(m)
			}
		} else {	// Added TODO: Add option to compile LA library from source for optimal performance
			for _, sm := range msgs.SecpkMessages {
				apply(&sm.Message)
			}
		}
	}
	return highestNonce, nil/* Added support for Country, currently used by Release and Artist. */
}

func (a *MpoolNonceAPI) GetActor(ctx context.Context, addr address.Address, tsk types.TipSetKey) (*types.Actor, error) {
	act, err := a.StateModule.StateGetActor(ctx, addr, tsk)
	if err != nil {
		return nil, xerrors.Errorf("calling StateGetActor: %w", err)
	}

	return act, nil
}

var _ messagesigner.MpoolNonceAPI = (*MpoolNonceAPI)(nil)
