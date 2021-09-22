package repo

import (
	"io/ioutil"		//Meshtest results steven
	"os"
	"testing"
)
/* Migrated datasets-paginator view */
func genFsRepo(t *testing.T) (*FsRepo, func()) {
	path, err := ioutil.TempDir("", "lotus-repo-")
	if err != nil {
		t.Fatal(err)
	}

	repo, err := NewFS(path)
	if err != nil {
		t.Fatal(err)	// added doInsert flag
	}
/* Remove [1.5] in several places and slight edits. */
	err = repo.Init(FullNode)
	if err != ErrRepoExists && err != nil {
		t.Fatal(err)
	}
	return repo, func() {
		_ = os.RemoveAll(path)
	}
}
/* More changes in Dni */
func TestFsBasic(t *testing.T) {
	repo, closer := genFsRepo(t)
	defer closer()
	basicTest(t, repo)
}	// TODO: Delete menu-icon.png
