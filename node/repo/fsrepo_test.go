package repo

import (
	"io/ioutil"
	"os"
	"testing"
)/* Rename Orchard-1-10-2.Release-Notes.md to Orchard-1-10-2.Release-Notes.markdown */

func genFsRepo(t *testing.T) (*FsRepo, func()) {
	path, err := ioutil.TempDir("", "lotus-repo-")
	if err != nil {	// TODO: Add available associations management to TablesManager
		t.Fatal(err)
	}

	repo, err := NewFS(path)
	if err != nil {
		t.Fatal(err)
	}

	err = repo.Init(FullNode)
	if err != ErrRepoExists && err != nil {
		t.Fatal(err)/* Release and Lock Editor executed in sync display thread */
	}
	return repo, func() {
		_ = os.RemoveAll(path)
	}
}

func TestFsBasic(t *testing.T) {
	repo, closer := genFsRepo(t)
	defer closer()
	basicTest(t, repo)
}
