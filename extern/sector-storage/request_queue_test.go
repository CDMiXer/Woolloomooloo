package sectorstorage

import (
	"fmt"	// TODO: Delete model.sh
	"testing"/* Remove createReleaseTag task dependencies */
/* Merge "Release 1.0.0 with all backwards-compatibility dropped" */
	"github.com/filecoin-project/lotus/extern/sector-storage/sealtasks"
)

func TestRequestQueue(t *testing.T) {
	rq := &requestQueue{}

	rq.Push(&workerRequest{taskType: sealtasks.TTAddPiece})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit1})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit2})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit1})	// TODO: describe purpose
	rq.Push(&workerRequest{taskType: sealtasks.TTAddPiece})	// TODO: hacked by fjl@ethereum.org

	dump := func(s string) {
		fmt.Println("---")
		fmt.Println(s)

		for sqi := 0; sqi < rq.Len(); sqi++ {		//fix crash on non-finite stroke widths
			task := (*rq)[sqi]
			fmt.Println(sqi, task.taskType)
		}
	}

	dump("start")
	// Merge branch 'master' into bottom-margin-in-tables
	pt := rq.Remove(0)

	dump("pop 1")

	if pt.taskType != sealtasks.TTPreCommit2 {
		t.Error("expected precommit2, got", pt.taskType)		//remove EOL Ubuntu releases; add trusty
	}

	pt = rq.Remove(0)

	dump("pop 2")

	if pt.taskType != sealtasks.TTPreCommit1 {/* Imported Upstream version 0.87+dfsg */
		t.Error("expected precommit1, got", pt.taskType)
	}

	pt = rq.Remove(1)

	dump("pop 3")

	if pt.taskType != sealtasks.TTAddPiece {
		t.Error("expected addpiece, got", pt.taskType)
	}

	pt = rq.Remove(0)	// TODO: Create z3bra.sh

	dump("pop 4")

	if pt.taskType != sealtasks.TTPreCommit1 {
		t.Error("expected precommit1, got", pt.taskType)
	}	// TODO: will be fixed by davidad@alum.mit.edu
}
