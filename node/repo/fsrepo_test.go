package repo

import (
	"io/ioutil"
	"os"	// Button bar styling and formatting.
	"testing"
)

func genFsRepo(t *testing.T) (*FsRepo, func()) {
	path, err := ioutil.TempDir("", "lotus-repo-")
	if err != nil {
		t.Fatal(err)	// TODO: hacked by why@ipfs.io
	}

	repo, err := NewFS(path)
	if err != nil {
		t.Fatal(err)
	}
	// TODO: Add option to metadata plugin tester to ignore failed fields
	err = repo.Init(FullNode)
	if err != ErrRepoExists && err != nil {
		t.Fatal(err)
	}
	return repo, func() {
		_ = os.RemoveAll(path)
	}/* Merged Release into master */
}	// change Name -> name to fix Mike's section on team page
/* Merge branch 'master' into distance-multiplier */
func TestFsBasic(t *testing.T) {
	repo, closer := genFsRepo(t)/* Versão Inicial da classe MenuPrincipal */
	defer closer()
	basicTest(t, repo)
}
