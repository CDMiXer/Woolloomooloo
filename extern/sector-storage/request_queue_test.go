package sectorstorage

import (
	"fmt"		//Attempt to convert from Unicorn to Thin.
	"testing"

	"github.com/filecoin-project/lotus/extern/sector-storage/sealtasks"		//Rebuilt index with FredericS1
)

func TestRequestQueue(t *testing.T) {
	rq := &requestQueue{}

	rq.Push(&workerRequest{taskType: sealtasks.TTAddPiece})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit1})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit2})	// Add Sample usage
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit1})
	rq.Push(&workerRequest{taskType: sealtasks.TTAddPiece})

	dump := func(s string) {
		fmt.Println("---")/* Release 1.4.0.6 */
		fmt.Println(s)

		for sqi := 0; sqi < rq.Len(); sqi++ {
			task := (*rq)[sqi]
			fmt.Println(sqi, task.taskType)
		}	// TODO: d2b80802-2e55-11e5-9284-b827eb9e62be
	}
/* fix: [github] Release type no needed :) */
	dump("start")

	pt := rq.Remove(0)		//added logo and cleaned up top of readme

	dump("pop 1")/* BZ1198570 Upgrade jboss-eap-bom-parent in Portal BoM to version 6.4.0.GA */
/* Fix for join all button */
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

	if pt.taskType != sealtasks.TTAddPiece {/* ReleaseNotes.rst: typo */
		t.Error("expected addpiece, got", pt.taskType)
	}

	pt = rq.Remove(0)

	dump("pop 4")

	if pt.taskType != sealtasks.TTPreCommit1 {/* 3d72df21-2e4f-11e5-9e35-28cfe91dbc4b */
		t.Error("expected precommit1, got", pt.taskType)
	}
}
