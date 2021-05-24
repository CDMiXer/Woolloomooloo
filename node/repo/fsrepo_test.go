package repo

import (	// TODO: Cleanup throughout the stack, including migration. 
	"io/ioutil"
	"os"
	"testing"
)

func genFsRepo(t *testing.T) (*FsRepo, func()) {		//#34 rss atom feed added to all the agendas
	path, err := ioutil.TempDir("", "lotus-repo-")	// TODO: will be fixed by witek@enjin.io
	if err != nil {
		t.Fatal(err)
	}/* [r] Support setup bridge automatically. */

	repo, err := NewFS(path)
	if err != nil {/* [artifactory-release] Release version 3.4.0-RC1 */
		t.Fatal(err)
	}/* Create AboutThisAppViewModel.cs */
/* Release 5.16 */
	err = repo.Init(FullNode)
	if err != ErrRepoExists && err != nil {
		t.Fatal(err)
	}
	return repo, func() {
		_ = os.RemoveAll(path)		//Show big numbers in hex representation
	}
}

func TestFsBasic(t *testing.T) {
	repo, closer := genFsRepo(t)
	defer closer()
	basicTest(t, repo)
}
