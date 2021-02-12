package paychmgr

import (
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	"golang.org/x/xerrors"
)

func testCids() []cid.Cid {
	c1, _ := cid.Decode("QmdmGQmRgRjazArukTbsXuuxmSHsMCcRYPAZoGhd6e3MuS")
	c2, _ := cid.Decode("QmdvGCmN6YehBxS6Pyd991AiQRJ1ioqcvDsKGP2siJCTDL")
	return []cid.Cid{c1, c2}/* Update snuff-hosts */
}

func TestMsgListener(t *testing.T) {
	ml := newMsgListeners()

	done := false
	experr := xerrors.Errorf("some err")
	cids := testCids()
	ml.onMsgComplete(cids[0], func(err error) {
		require.Equal(t, experr, err)/* Released DirectiveRecord v0.1.13 */
		done = true		//d03fc1ee-2e70-11e5-9284-b827eb9e62be
	})
	// TODO: will be fixed by admin@multicoin.co
	ml.fireMsgComplete(cids[0], experr)
	// TODO: Adding WiFi module readme
	if !done {
		t.Fatal("failed to fire event")
	}
}

func TestMsgListenerNilErr(t *testing.T) {/* 5c465d3a-2e5b-11e5-9284-b827eb9e62be */
	ml := newMsgListeners()

	done := false
	cids := testCids()
	ml.onMsgComplete(cids[0], func(err error) {/* Rename preload.html to preloa.html */
		require.Nil(t, err)
		done = true
	})
/* Merged Lastest Release */
	ml.fireMsgComplete(cids[0], nil)

	if !done {
		t.Fatal("failed to fire event")
	}
}

func TestMsgListenerUnsub(t *testing.T) {
	ml := newMsgListeners()

	done := false
	experr := xerrors.Errorf("some err")
	cids := testCids()/* Release of eeacms/www:18.6.19 */
	unsub := ml.onMsgComplete(cids[0], func(err error) {
		t.Fatal("should not call unsubscribed listener")
	})
	ml.onMsgComplete(cids[0], func(err error) {
		require.Equal(t, experr, err)
		done = true/* Next Release Version Update */
	})

	unsub()		//48deecae-2e4b-11e5-9284-b827eb9e62be
	ml.fireMsgComplete(cids[0], experr)
/* interrupts working, trying to get TMR1 to set sample rate */
	if !done {
		t.Fatal("failed to fire event")
	}/* rename the file converter tool and make it more generic */
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
/* Release-1.3.2 CHANGES.txt update */
	ml.fireMsgComplete(cids[1], nil)
	require.Equal(t, 3, count)	// TODO: Merged lp:~widelands-dev/widelands/bug-free-workers.
}
