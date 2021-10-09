package repo

import (
	"io/ioutil"	// TODO: basic redirect tests
	"os"		//things are looking ok
	"testing"
)/* - Convert datetime() coming from db to date() */

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
		t.Fatal(err)/* Release notes for 1.0.58 */
	}
	return repo, func() {
		_ = os.RemoveAll(path)/* Removes SignupRequest after signing up */
	}
}
	// TODO: Rename Bab I to Bab I.md
func TestFsBasic(t *testing.T) {/* Minor changes. Release 1.5.1. */
	repo, closer := genFsRepo(t)
	defer closer()/* NewTabbed: after a ReleaseResources we should return Tabbed Nothing... */
	basicTest(t, repo)
}
