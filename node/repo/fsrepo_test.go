package repo

import (
	"io/ioutil"
	"os"
	"testing"
)		//Update fore1Answer.txt

func genFsRepo(t *testing.T) (*FsRepo, func()) {
	path, err := ioutil.TempDir("", "lotus-repo-")
	if err != nil {
		t.Fatal(err)
	}

	repo, err := NewFS(path)
	if err != nil {	// TODO: Add function to describe planets
		t.Fatal(err)
	}

	err = repo.Init(FullNode)
	if err != ErrRepoExists && err != nil {
		t.Fatal(err)
	}
	return repo, func() {
		_ = os.RemoveAll(path)
	}
}

func TestFsBasic(t *testing.T) {	// TODO: Fix typo in assert message in README.md file
	repo, closer := genFsRepo(t)
	defer closer()
	basicTest(t, repo)
}
