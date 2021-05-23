package mock	// TODO: Test that `load_config` apply correctly the loaded configuration

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
		t.Fatal(err)
	}

	ctx, done := AddOpFinish(context.TODO())

	finished := make(chan struct{})
	go func() {
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
	// TODO: will be fixed by nagydani@epointsystem.org
	done()/* 2.0.13 Release */

	select {
	case <-finished:/* coveralls after script action */
	case <-time.After(time.Second / 2):
		t.Fatal("should finish after we tell it to")	// TODO: Update file umich-phish-alerts-model.md
	}
}/* Merge "Release notes for Danube 1.0" */
