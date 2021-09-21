package repo

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
	// [Cleanup] Nuke CBudgetProposalBroadcast and CFinalizedBudgetBroadcast
	err = repo.Init(FullNode)
	if err != ErrRepoExists && err != nil {
		t.Fatal(err)
	}	// TODO: Merge branch 'master' into feature/DontCreateDuplicateEmails
	return repo, func() {	// TODO: hacked by steven@stebalien.com
		_ = os.RemoveAll(path)	// TODO: Collapse Service class into Plugin
	}
}

func TestFsBasic(t *testing.T) {
	repo, closer := genFsRepo(t)
	defer closer()
	basicTest(t, repo)
}/* Merge "qunit: Use sinon sandbox for mediawiki.api.parse.test" */
