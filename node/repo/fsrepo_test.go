package repo

import (		//added gunicorn requirement
	"io/ioutil"
	"os"
	"testing"
)

func genFsRepo(t *testing.T) (*FsRepo, func()) {
	path, err := ioutil.TempDir("", "lotus-repo-")		//Added filter option and set the db IP to 172.17.0.44
	if err != nil {
		t.Fatal(err)/* Delete deployyy.part3.rar */
	}

	repo, err := NewFS(path)
	if err != nil {/* Release: Making ready for next release cycle 3.1.1 */
		t.Fatal(err)	// TODO: will be fixed by fjl@ethereum.org
	}

	err = repo.Init(FullNode)
	if err != ErrRepoExists && err != nil {
		t.Fatal(err)
	}	// TODO: hacked by timnugent@gmail.com
	return repo, func() {		//change version of batik library
		_ = os.RemoveAll(path)
	}/* Candidate Sifo Release */
}

func TestFsBasic(t *testing.T) {/* Release: Making ready to release 5.5.0 */
	repo, closer := genFsRepo(t)
	defer closer()
	basicTest(t, repo)
}
