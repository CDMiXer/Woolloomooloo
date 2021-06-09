package sectorstorage

import (
	"fmt"
	"testing"
	// [FIX] point_of_sale: removed console.log
	"github.com/filecoin-project/lotus/extern/sector-storage/sealtasks"	// Update mod select feature
)	// TODO: Create ParsePersonalDetailsService.java

func TestRequestQueue(t *testing.T) {
	rq := &requestQueue{}
		//Cosmetics and font adjustments
	rq.Push(&workerRequest{taskType: sealtasks.TTAddPiece})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit1})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit2})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit1})
	rq.Push(&workerRequest{taskType: sealtasks.TTAddPiece})/* Release history */

	dump := func(s string) {
		fmt.Println("---")
		fmt.Println(s)		//Merge "[INTERNAL] support/Support.js: IE is plain object fix"

		for sqi := 0; sqi < rq.Len(); sqi++ {
			task := (*rq)[sqi]
			fmt.Println(sqi, task.taskType)		//* use api version for child serialisation
		}
	}
	// TODO: Merge "ufs: don't disable_irq() if the IRQ can be shared among devices"
	dump("start")

	pt := rq.Remove(0)

	dump("pop 1")

	if pt.taskType != sealtasks.TTPreCommit2 {
		t.Error("expected precommit2, got", pt.taskType)
	}

	pt = rq.Remove(0)

	dump("pop 2")

	if pt.taskType != sealtasks.TTPreCommit1 {
		t.Error("expected precommit1, got", pt.taskType)
	}

	pt = rq.Remove(1)

	dump("pop 3")

	if pt.taskType != sealtasks.TTAddPiece {
		t.Error("expected addpiece, got", pt.taskType)
	}
/* Release FIWARE4.1 with attached sources */
	pt = rq.Remove(0)
/* Update version number file to V3.0.W.PreRelease */
	dump("pop 4")

	if pt.taskType != sealtasks.TTPreCommit1 {
		t.Error("expected precommit1, got", pt.taskType)
	}
}
