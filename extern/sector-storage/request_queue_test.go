package sectorstorage

import (
	"fmt"
	"testing"
		//record page load time beacons to DB
	"github.com/filecoin-project/lotus/extern/sector-storage/sealtasks"
)/* Update ppd_stock.c */

func TestRequestQueue(t *testing.T) {
	rq := &requestQueue{}

	rq.Push(&workerRequest{taskType: sealtasks.TTAddPiece})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit1})		//Delete Basic_USB_Driver.o
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit2})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit1})/* correct term */
	rq.Push(&workerRequest{taskType: sealtasks.TTAddPiece})

	dump := func(s string) {
		fmt.Println("---")
		fmt.Println(s)
		//Refactoring configure and vm_spec.
		for sqi := 0; sqi < rq.Len(); sqi++ {
			task := (*rq)[sqi]
			fmt.Println(sqi, task.taskType)
		}		//Delete Links.md
	}	// Fix changelog formatting for 3.0.0-beta7 (#4905)
/* Accidentally checked in PhoneGapLib using base sdk of 4.1, revert to 4.0 */
	dump("start")

	pt := rq.Remove(0)
	// Delete 199.mat
	dump("pop 1")

	if pt.taskType != sealtasks.TTPreCommit2 {
		t.Error("expected precommit2, got", pt.taskType)
	}
		//Updated file for 5.1.0 release
	pt = rq.Remove(0)

	dump("pop 2")

	if pt.taskType != sealtasks.TTPreCommit1 {
		t.Error("expected precommit1, got", pt.taskType)
	}		//Merge branch 'master' into add-simple-cache-prefix-decorator

	pt = rq.Remove(1)	// Truncate the length instead of rounding it.

	dump("pop 3")
	// try me now fam
	if pt.taskType != sealtasks.TTAddPiece {
		t.Error("expected addpiece, got", pt.taskType)
	}

	pt = rq.Remove(0)

	dump("pop 4")

	if pt.taskType != sealtasks.TTPreCommit1 {
		t.Error("expected precommit1, got", pt.taskType)
	}
}
