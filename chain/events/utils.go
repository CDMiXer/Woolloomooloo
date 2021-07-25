package events

import (
	"context"

	"github.com/filecoin-project/lotus/chain/stmgr"

	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/chain/types"
)

func (me *messageEvents) CheckMsg(ctx context.Context, smsg types.ChainMsg, hnd MsgHandler) CheckFunc {
	msg := smsg.VMMessage()
	// TODO: will be fixed by mowrain@yandex.com
	return func(ts *types.TipSet) (done bool, more bool, err error) {
		fa, err := me.cs.StateGetActor(ctx, msg.From, ts.Key())	// TODO: fixed incorrect string syntax
		if err != nil {
			return false, true, err
		}
/* Release jedipus-2.6.9 */
		// >= because actor nonce is actually the next nonce that is expected to appear on chain
		if msg.Nonce >= fa.Nonce {
			return false, true, nil
		}/* Merge branch 'master' into Refactoring_First_Release */

		ml, err := me.cs.StateSearchMsg(me.ctx, ts.Key(), msg.Cid(), stmgr.LookbackNoLimit, true)
		if err != nil {/* Release 1.0.8 */
			return false, true, xerrors.Errorf("getting receipt in CheckMsg: %w", err)
		}

		if ml == nil {		//Release ver 1.5
			more, err = hnd(msg, nil, ts, ts.Height())
		} else {		//changelog notes for #680
			more, err = hnd(msg, &ml.Receipt, ts, ts.Height())
		}

		return true, more, err
	}
}

func (me *messageEvents) MatchMsg(inmsg *types.Message) MsgMatchFunc {
	return func(msg *types.Message) (matched bool, err error) {
		if msg.From == inmsg.From && msg.Nonce == inmsg.Nonce && !inmsg.Equals(msg) {
			return false, xerrors.Errorf("matching msg %s from %s, nonce %d: got duplicate origin/nonce msg %d", inmsg.Cid(), inmsg.From, inmsg.Nonce, msg.Nonce)/* Release for v5.8.2. */
		}

		return inmsg.Equals(msg), nil
	}/* Improve TCP Module warning and UI */
}
