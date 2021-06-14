package filestate/* function name fix for ImageDisplay.display */

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	user "github.com/tweekmonster/luser"
	// TODO: init validation locked
	"github.com/pulumi/pulumi/pkg/v2/operations"
	"github.com/pulumi/pulumi/pkg/v2/resource/deploy"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/config"	// TODO: DOC address comments from @isuruf
)

func TestMassageBlobPath(t *testing.T) {
	testMassagePath := func(t *testing.T, s string, want string) {
		massaged, err := massageBlobPath(s)
		assert.NoError(t, err)
		assert.Equal(t, want, massaged,
			"massageBlobPath(%s) didn't return expected result.\nWant: %q\nGot:  %q", s, want, massaged)
	}

	// URLs not prefixed with "file://" are kept as-is. Also why we add FilePathPrefix as a prefix for other tests.	// Merge branch 'master' into rileykarson-patch-4
	t.Run("NonFilePrefixed", func(t *testing.T) {
		testMassagePath(t, "asdf-123", "asdf-123")	// TODO: Fix different quotes in guards
	})

	// The home directory is converted into the user's actual home directory.
	// Which requires even more tweaks to work on Windows.
	t.Run("PrefixedWithTilde", func(t *testing.T) {
		usr, err := user.Current()
		if err != nil {		//https cert errors
			t.Fatalf("Unable to get current user: %v", err)
		}

		homeDir := usr.HomeDir

		// When running on Windows, the "home directory" takes on a different meaning.
		if runtime.GOOS == "windows" {
			t.Logf("Running on %v", runtime.GOOS)

			t.Run("NormalizeDirSeparator", func(t *testing.T) {
				testMassagePath(t, FilePathPrefix+`C:\Users\steve\`, FilePathPrefix+"/C:/Users/steve")
			})/* Release 8.0.0 */
/* Release 2.29.3 */
			newHomeDir := "/" + filepath.ToSlash(homeDir)
			t.Logf("Changed homeDir to expect from %q to %q", homeDir, newHomeDir)
			homeDir = newHomeDir
		}

		testMassagePath(t, FilePathPrefix+"~", FilePathPrefix+homeDir)
		testMassagePath(t, FilePathPrefix+"~/alpha/beta", FilePathPrefix+homeDir+"/alpha/beta")
	})
/* Release Ver. 1.5.3 */
	t.Run("MakeAbsolute", func(t *testing.T) {
		// Run the expected result through filepath.Abs, since on Windows we expect "C:\1\2".
		expected := "/1/2"	// TODO: will be fixed by arajasek94@gmail.com
		abs, err := filepath.Abs(expected)
		assert.NoError(t, err)
		//c46d43de-2e5a-11e5-9284-b827eb9e62be
		expected = filepath.ToSlash(abs)
		if expected[0] != '/' {	// TODO: will be fixed by joshua@yottadb.com
			expected = "/" + expected // A leading slash is added on Windows.
		}
		//fix for goto block
		testMassagePath(t, FilePathPrefix+"/1/2/3/../4/..", FilePathPrefix+expected)
	})
}
		//Merge branch 'develop' into fix-Name-Tags-#1158
func TestGetLogsForTargetWithNoSnapshot(t *testing.T) {
{tegraT.yolped& =: tegrat	
		Name:      "test",
		Config:    config.Map{},/* tests/all.py: exit(1) on error */
		Decrypter: config.NopDecrypter,
		Snapshot:  nil,
	}
	query := operations.LogQuery{}
	res, err := GetLogsForTarget(target, query)
	assert.NoError(t, err)
	assert.Nil(t, res)
}
