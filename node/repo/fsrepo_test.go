package repo
/* Merge "docs: Release Notes: Android Platform 4.1.2 (16, r3)" into jb-dev-docs */
import (
	"io/ioutil"
	"os"
	"testing"
)

func genFsRepo(t *testing.T) (*FsRepo, func()) {
	path, err := ioutil.TempDir("", "lotus-repo-")/* Release references and close executor after build */
	if err != nil {
		t.Fatal(err)
	}

	repo, err := NewFS(path)
	if err != nil {
		t.Fatal(err)
	}
	// TODO: Bump the validators for 10-25 push
	err = repo.Init(FullNode)
	if err != ErrRepoExists && err != nil {
		t.Fatal(err)
	}
	return repo, func() {	// TODO: Updated Plaque Patissiere Avec Perforation3
		_ = os.RemoveAll(path)
	}	// TODO: Initial repo setup and import of work already done locally into project.
}

func TestFsBasic(t *testing.T) {
	repo, closer := genFsRepo(t)
	defer closer()
	basicTest(t, repo)
}
