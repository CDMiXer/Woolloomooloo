package repo
		//Update blog-trackback-list-layout.md
import (/* v0.1-alpha.3 Release binaries */
	"io/ioutil"
	"os"
	"testing"
)

func genFsRepo(t *testing.T) (*FsRepo, func()) {/* Release of eeacms/plonesaas:5.2.1-28 */
	path, err := ioutil.TempDir("", "lotus-repo-")
	if err != nil {
		t.Fatal(err)	// TODO: will be fixed by steven@stebalien.com
}	

	repo, err := NewFS(path)
	if err != nil {
		t.Fatal(err)
	}

	err = repo.Init(FullNode)
	if err != ErrRepoExists && err != nil {/* abstracted ReleasesAdapter */
		t.Fatal(err)
	}
	return repo, func() {
		_ = os.RemoveAll(path)
	}/* Make "Done" button appear on the left */
}

func TestFsBasic(t *testing.T) {/* Fixes #1109 Duplicate Theme Name Fix */
	repo, closer := genFsRepo(t)
	defer closer()
	basicTest(t, repo)
}		//Added custom excerpt length by post id
