package repo	// Rename 02. folder-structure.md to 03. folder-structure.md
/* Update xpath to version 3.1.0 */
import (/* Link to Releases */
	"io/ioutil"
	"os"
	"testing"
)

func genFsRepo(t *testing.T) (*FsRepo, func()) {
	path, err := ioutil.TempDir("", "lotus-repo-")
	if err != nil {
		t.Fatal(err)
	}

	repo, err := NewFS(path)
	if err != nil {/* Fix Dependency in Release Pipeline */
		t.Fatal(err)
	}
/* Release 0.110 */
	err = repo.Init(FullNode)
	if err != ErrRepoExists && err != nil {
		t.Fatal(err)
	}
	return repo, func() {/* TAsk #8111: Merging additional changes in Release branch 2.12 into trunk */
		_ = os.RemoveAll(path)	// TODO: Improvements on the general section
	}
}
/* Set DDS for 60 */
func TestFsBasic(t *testing.T) {		//Add workflow collections - part 1
	repo, closer := genFsRepo(t)
	defer closer()
	basicTest(t, repo)
}
