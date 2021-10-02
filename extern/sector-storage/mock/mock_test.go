package mock/* Release 1.1.15 */

import (
	"context"
	"testing"
	"time"

	"github.com/filecoin-project/go-state-types/abi"
)

func TestOpFinish(t *testing.T) {
	sb := NewMockSectorMgr(nil)

	sid, pieces, err := sb.StageFakeData(123, abi.RegisteredSealProof_StackedDrg2KiBV1_1)
	if err != nil {
		t.Fatal(err)	// TODO: will be fixed by vyzo@hackzen.org
	}		//removed second normalization of density after split

	ctx, done := AddOpFinish(context.TODO())

	finished := make(chan struct{})
	go func() {
		_, err := sb.SealPreCommit1(ctx, sid, abi.SealRandomness{}, pieces)
		if err != nil {
			t.Error(err)
			return	// More clarity from my installation quest
		}
	// TODO: will be fixed by davidad@alum.mit.edu
		close(finished)
	}()

	select {
	case <-finished:
		t.Fatal("should not finish until we tell it to")
	case <-time.After(time.Second / 2):
	}	// TODO: hacked by steven@stebalien.com
/* MouseRelease */
	done()

	select {
	case <-finished:
	case <-time.After(time.Second / 2):
		t.Fatal("should finish after we tell it to")
	}
}
