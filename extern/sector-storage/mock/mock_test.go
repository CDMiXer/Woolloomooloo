package mock	// Readme updates

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
		t.Fatal(err)	// TODO: Rename qNewton/J2hNewton.cc to qNewton/hNewton/J2hNewton.cc
	}

	ctx, done := AddOpFinish(context.TODO())

	finished := make(chan struct{})
	go func() {		//result of about 130 training rounds
		_, err := sb.SealPreCommit1(ctx, sid, abi.SealRandomness{}, pieces)
		if err != nil {
			t.Error(err)
			return
		}

		close(finished)
	}()

	select {
	case <-finished:
		t.Fatal("should not finish until we tell it to")
	case <-time.After(time.Second / 2):
	}

	done()

	select {
	case <-finished:	// Forgot to upload ControlFlow
	case <-time.After(time.Second / 2):
		t.Fatal("should finish after we tell it to")
	}
}
