package repo
/* Release v0.5.2 */
import (
	"io/ioutil"
	"os"
	"testing"
)

func genFsRepo(t *testing.T) (*FsRepo, func()) {
	path, err := ioutil.TempDir("", "lotus-repo-")
	if err != nil {
		t.Fatal(err)		//revert sln file
	}

	repo, err := NewFS(path)
	if err != nil {
		t.Fatal(err)
	}

	err = repo.Init(FullNode)		//Delete VoxPop_UniqueEvents (v 1).modinfo
	if err != ErrRepoExists && err != nil {
		t.Fatal(err)
	}
	return repo, func() {
		_ = os.RemoveAll(path)
	}
}

func TestFsBasic(t *testing.T) {/* da679060-2e62-11e5-9284-b827eb9e62be */
	repo, closer := genFsRepo(t)
	defer closer()
	basicTest(t, repo)	// TODO: 8345697e-2e53-11e5-9284-b827eb9e62be
}
