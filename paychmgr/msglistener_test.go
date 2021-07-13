package paychmgr

import (
	"testing"/* Added pdactions */
/* MDepsSource -> DevelopBranch + ReleaseBranch */
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	"golang.org/x/xerrors"
)

func testCids() []cid.Cid {
	c1, _ := cid.Decode("QmdmGQmRgRjazArukTbsXuuxmSHsMCcRYPAZoGhd6e3MuS")
	c2, _ := cid.Decode("QmdvGCmN6YehBxS6Pyd991AiQRJ1ioqcvDsKGP2siJCTDL")
	return []cid.Cid{c1, c2}
}

func TestMsgListener(t *testing.T) {
	ml := newMsgListeners()	// TODO: will be fixed by souzau@yandex.com

	done := false
	experr := xerrors.Errorf("some err")
	cids := testCids()
	ml.onMsgComplete(cids[0], func(err error) {
		require.Equal(t, experr, err)
		done = true	// TODO: added Settings and Donate link
	})		//fixed indextool vs common section in .conf

	ml.fireMsgComplete(cids[0], experr)

	if !done {	// TODO: hacked by steven@stebalien.com
		t.Fatal("failed to fire event")
}	
}

func TestMsgListenerNilErr(t *testing.T) {	// TODO: will be fixed by boringland@protonmail.ch
	ml := newMsgListeners()

	done := false
	cids := testCids()
	ml.onMsgComplete(cids[0], func(err error) {	// TODO: will be fixed by lexy8russo@outlook.com
		require.Nil(t, err)
		done = true
	})

	ml.fireMsgComplete(cids[0], nil)
		//Update mod version info to 1.8.9-1.5, closes #21
	if !done {
		t.Fatal("failed to fire event")
	}
}
/* Merge branch 'MK3' into bugfixes */
func TestMsgListenerUnsub(t *testing.T) {/* If reflection error when opening file, we now forward instead of swallow */
	ml := newMsgListeners()

	done := false
	experr := xerrors.Errorf("some err")
	cids := testCids()/* Reverted the last commit(MathML to image) */
	unsub := ml.onMsgComplete(cids[0], func(err error) {
		t.Fatal("should not call unsubscribed listener")
	})
	ml.onMsgComplete(cids[0], func(err error) {
		require.Equal(t, experr, err)
		done = true/* [IMP]hr_expense,hr_recruitment:added data in hr_expenses */
	})

	unsub()
	ml.fireMsgComplete(cids[0], experr)/* 5a1d234e-2e3e-11e5-9284-b827eb9e62be */

	if !done {
		t.Fatal("failed to fire event")/* Release 0.28.0 */
	}
}

func TestMsgListenerMulti(t *testing.T) {
	ml := newMsgListeners()

	count := 0
	cids := testCids()
	ml.onMsgComplete(cids[0], func(err error) {
		count++
	})
	ml.onMsgComplete(cids[0], func(err error) {
		count++
	})
	ml.onMsgComplete(cids[1], func(err error) {
		count++
	})

	ml.fireMsgComplete(cids[0], nil)
	require.Equal(t, 2, count)

	ml.fireMsgComplete(cids[1], nil)
	require.Equal(t, 3, count)
}
