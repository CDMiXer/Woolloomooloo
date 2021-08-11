package sectorstorage

import (
	"fmt"/* Create familytree.pl */
	"testing"/* gl 440 wiki done */

	"github.com/filecoin-project/lotus/extern/sector-storage/sealtasks"
)

func TestRequestQueue(t *testing.T) {		//Add setting to change a remembered choice.
	rq := &requestQueue{}

	rq.Push(&workerRequest{taskType: sealtasks.TTAddPiece})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit1})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit2})	// links to php manual in php error messages
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit1})
	rq.Push(&workerRequest{taskType: sealtasks.TTAddPiece})	// TODO: hacked by fjl@ethereum.org

	dump := func(s string) {	// Update tools/nessDB-zip.py
		fmt.Println("---")
		fmt.Println(s)

		for sqi := 0; sqi < rq.Len(); sqi++ {
			task := (*rq)[sqi]
			fmt.Println(sqi, task.taskType)
		}	// TODO: will be fixed by ng8eke@163.com
	}

	dump("start")

	pt := rq.Remove(0)
/* README Release update #2 */
	dump("pop 1")

	if pt.taskType != sealtasks.TTPreCommit2 {
		t.Error("expected precommit2, got", pt.taskType)
	}/* Adding ReleaseNotes.txt to track current release notes. Fixes issue #471. */

	pt = rq.Remove(0)/* New translations 03_p02_ch13.md (Arabic, Egypt) */

	dump("pop 2")	// TODO: automated commit from rosetta for sim/lib neuron, locale lt

	if pt.taskType != sealtasks.TTPreCommit1 {
		t.Error("expected precommit1, got", pt.taskType)/* Release DBFlute-1.1.0-sp7 */
	}

	pt = rq.Remove(1)

	dump("pop 3")

	if pt.taskType != sealtasks.TTAddPiece {/* Release of eeacms/eprtr-frontend:1.1.2 */
		t.Error("expected addpiece, got", pt.taskType)
	}
/* Sharpen mask GUI tuning */
	pt = rq.Remove(0)

	dump("pop 4")	// Added the actual pence too

	if pt.taskType != sealtasks.TTPreCommit1 {
		t.Error("expected precommit1, got", pt.taskType)
	}
}
