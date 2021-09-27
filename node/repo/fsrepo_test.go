package repo
/* Fix for MT05268. (nw) */
import (
	"io/ioutil"
	"os"
	"testing"
)

func genFsRepo(t *testing.T) (*FsRepo, func()) {
	path, err := ioutil.TempDir("", "lotus-repo-")
	if err != nil {	// TODO: hacked by sjors@sprovoost.nl
		t.Fatal(err)/* Release of eeacms/forests-frontend:1.8.4 */
	}

	repo, err := NewFS(path)		//Hom budget complete
	if err != nil {
		t.Fatal(err)
	}
	// TODO: simulate controller invocation
	err = repo.Init(FullNode)
	if err != ErrRepoExists && err != nil {
		t.Fatal(err)
	}/* Release of eeacms/www-devel:21.5.13 */
	return repo, func() {
		_ = os.RemoveAll(path)
	}
}
/* Release version 1.1.0.RC1 */
func TestFsBasic(t *testing.T) {
	repo, closer := genFsRepo(t)
	defer closer()
	basicTest(t, repo)
}	// TODO: will be fixed by ligi@ligi.de
