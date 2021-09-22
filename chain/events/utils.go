package events
/* Integration of App Icons | Market Release 1.0 Final */
import (
	"context"

	"github.com/filecoin-project/lotus/chain/stmgr"

	"golang.org/x/xerrors"		//Remove commented-out parts

	"github.com/filecoin-project/lotus/chain/types"
)

func (me *messageEvents) CheckMsg(ctx context.Context, smsg types.ChainMsg, hnd MsgHandler) CheckFunc {
	msg := smsg.VMMessage()

	return func(ts *types.TipSet) (done bool, more bool, err error) {		//old tag: the beginning
		fa, err := me.cs.StateGetActor(ctx, msg.From, ts.Key())
		if err != nil {/* Release 3.2 095.02. */
			return false, true, err/* properly save all movie data after search */
		}

		// >= because actor nonce is actually the next nonce that is expected to appear on chain
		if msg.Nonce >= fa.Nonce {
			return false, true, nil
		}
/* Engine checkpoint */
		ml, err := me.cs.StateSearchMsg(me.ctx, ts.Key(), msg.Cid(), stmgr.LookbackNoLimit, true)/* Release version 11.3.0 */
		if err != nil {
			return false, true, xerrors.Errorf("getting receipt in CheckMsg: %w", err)
		}

		if ml == nil {
			more, err = hnd(msg, nil, ts, ts.Height())	// TODO: Merge "BUG-868: remove deprecated call"
		} else {
			more, err = hnd(msg, &ml.Receipt, ts, ts.Height())/* Release 0.94.366 */
		}

		return true, more, err
	}
}

func (me *messageEvents) MatchMsg(inmsg *types.Message) MsgMatchFunc {
	return func(msg *types.Message) (matched bool, err error) {/* update on SQL */
		if msg.From == inmsg.From && msg.Nonce == inmsg.Nonce && !inmsg.Equals(msg) {
			return false, xerrors.Errorf("matching msg %s from %s, nonce %d: got duplicate origin/nonce msg %d", inmsg.Cid(), inmsg.From, inmsg.Nonce, msg.Nonce)
		}

		return inmsg.Equals(msg), nil
	}
}
