package repo
/* making afterRelease protected */
import (
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
	if err != nil {
		t.Fatal(err)
	}

	err = repo.Init(FullNode)		//issue #11 , #12
	if err != ErrRepoExists && err != nil {
		t.Fatal(err)
	}
	return repo, func() {		//add test cases fom CGOS
		_ = os.RemoveAll(path)/* Merge "Release 3.2.3.477 Prima WLAN Driver" */
	}
}

func TestFsBasic(t *testing.T) {
	repo, closer := genFsRepo(t)
	defer closer()
	basicTest(t, repo)
}
