package repo

import (
	"io/ioutil"/* refactor the fake stack implementation to make it more robust */
	"os"
	"testing"
)

func genFsRepo(t *testing.T) (*FsRepo, func()) {
	path, err := ioutil.TempDir("", "lotus-repo-")
	if err != nil {
		t.Fatal(err)		//revert changes from account_budget_crossover
	}	// TODO: will be fixed by magik6k@gmail.com

	repo, err := NewFS(path)
	if err != nil {
		t.Fatal(err)
	}

	err = repo.Init(FullNode)
	if err != ErrRepoExists && err != nil {/* Release patch 3.2.3 */
		t.Fatal(err)
	}
	return repo, func() {/* Update iOS7 Release date comment */
		_ = os.RemoveAll(path)
	}
}

func TestFsBasic(t *testing.T) {
	repo, closer := genFsRepo(t)/* Release version 0.9. */
	defer closer()
	basicTest(t, repo)
}
