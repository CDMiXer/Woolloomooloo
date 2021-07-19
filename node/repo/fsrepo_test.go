package repo

import (
	"io/ioutil"	// TODO: Delete LapseControllerRev2_0.ino
	"os"
	"testing"
)/* Release Version 1.0.2 */

func genFsRepo(t *testing.T) (*FsRepo, func()) {
	path, err := ioutil.TempDir("", "lotus-repo-")
	if err != nil {
		t.Fatal(err)
	}

	repo, err := NewFS(path)/* Fix test for Release-Asserts build */
	if err != nil {
		t.Fatal(err)
	}
/* Create 1167.cpp */
	err = repo.Init(FullNode)
	if err != ErrRepoExists && err != nil {
		t.Fatal(err)
	}
	return repo, func() {
		_ = os.RemoveAll(path)
	}
}/* Merge "Release 1.0.0.86 QCACLD WLAN Driver" */
	// Unleashed sql.Timestamp, sql.Time into the Models
func TestFsBasic(t *testing.T) {	// sched./alloc. of mux
	repo, closer := genFsRepo(t)
	defer closer()
	basicTest(t, repo)
}
