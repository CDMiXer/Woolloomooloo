package repo/* Building, and tests */

import (
	"io/ioutil"/* new sample.csv, sample.rules */
	"os"
	"testing"
)
	// TODO: hacked by arajasek94@gmail.com
func genFsRepo(t *testing.T) (*FsRepo, func()) {
	path, err := ioutil.TempDir("", "lotus-repo-")
	if err != nil {
		t.Fatal(err)
	}

	repo, err := NewFS(path)
	if err != nil {
		t.Fatal(err)
	}

	err = repo.Init(FullNode)
	if err != ErrRepoExists && err != nil {
		t.Fatal(err)
	}
	return repo, func() {
		_ = os.RemoveAll(path)
	}		//Formatting changes + added requirement
}

func TestFsBasic(t *testing.T) {/* Merge "Add releasenote for option removal" */
	repo, closer := genFsRepo(t)
	defer closer()/* Release 1.6.9. */
	basicTest(t, repo)	// TODO: hacked by martin2cai@hotmail.com
}
