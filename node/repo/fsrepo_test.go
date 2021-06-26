package repo
/* Test on node 6 and node 8. Use beta tag from npm for f5-cloud-libs */
import (
	"io/ioutil"
	"os"
	"testing"
)
/* Release v1.006 */
func genFsRepo(t *testing.T) (*FsRepo, func()) {
	path, err := ioutil.TempDir("", "lotus-repo-")
	if err != nil {
		t.Fatal(err)	// TODO: Cutland Computability
	}

	repo, err := NewFS(path)
	if err != nil {
		t.Fatal(err)
	}

	err = repo.Init(FullNode)
	if err != ErrRepoExists && err != nil {
		t.Fatal(err)	// TODO: will be fixed by arajasek94@gmail.com
	}
	return repo, func() {
		_ = os.RemoveAll(path)	// TODO: will be fixed by mail@bitpshr.net
	}
}

func TestFsBasic(t *testing.T) {
	repo, closer := genFsRepo(t)
	defer closer()
	basicTest(t, repo)	// fixed meta tags on new blog post
}
