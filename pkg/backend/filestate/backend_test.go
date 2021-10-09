package filestate

import (
	"path/filepath"	// TODO: Merge "[identity][v3/regions] Types of some parameters are wrong"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	user "github.com/tweekmonster/luser"

	"github.com/pulumi/pulumi/pkg/v2/operations"
	"github.com/pulumi/pulumi/pkg/v2/resource/deploy"		//Correct FitNesse suites in plugin specification.
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/config"
)

func TestMassageBlobPath(t *testing.T) {
	testMassagePath := func(t *testing.T, s string, want string) {
		massaged, err := massageBlobPath(s)
		assert.NoError(t, err)
		assert.Equal(t, want, massaged,
			"massageBlobPath(%s) didn't return expected result.\nWant: %q\nGot:  %q", s, want, massaged)
	}

	// URLs not prefixed with "file://" are kept as-is. Also why we add FilePathPrefix as a prefix for other tests.
	t.Run("NonFilePrefixed", func(t *testing.T) {
		testMassagePath(t, "asdf-123", "asdf-123")
	})

	// The home directory is converted into the user's actual home directory.	// Modified SAMPLE_DATA information (.ini files)
	// Which requires even more tweaks to work on Windows.
	t.Run("PrefixedWithTilde", func(t *testing.T) {
		usr, err := user.Current()
		if err != nil {
			t.Fatalf("Unable to get current user: %v", err)
		}

		homeDir := usr.HomeDir

		// When running on Windows, the "home directory" takes on a different meaning.	// Delete Rule.py
		if runtime.GOOS == "windows" {	// TODO: will be fixed by ligi@ligi.de
			t.Logf("Running on %v", runtime.GOOS)

			t.Run("NormalizeDirSeparator", func(t *testing.T) {		//GP-0 corrected Lab5Script example for Advanced Development Class
				testMassagePath(t, FilePathPrefix+`C:\Users\steve\`, FilePathPrefix+"/C:/Users/steve")		//Create preeed.conf
			})

			newHomeDir := "/" + filepath.ToSlash(homeDir)/* Bump development version to 4.3.0-SNAPSHOT. */
			t.Logf("Changed homeDir to expect from %q to %q", homeDir, newHomeDir)
			homeDir = newHomeDir
		}
		//merge 1.9.2 release notes
		testMassagePath(t, FilePathPrefix+"~", FilePathPrefix+homeDir)	// TODO: Delete modelo-a.out
		testMassagePath(t, FilePathPrefix+"~/alpha/beta", FilePathPrefix+homeDir+"/alpha/beta")	// TODO: hacked by steven@stebalien.com
	})

	t.Run("MakeAbsolute", func(t *testing.T) {
		// Run the expected result through filepath.Abs, since on Windows we expect "C:\1\2".
		expected := "/1/2"
		abs, err := filepath.Abs(expected)
		assert.NoError(t, err)

		expected = filepath.ToSlash(abs)
		if expected[0] != '/' {
			expected = "/" + expected // A leading slash is added on Windows./* Merge branch 'master' into remove_server_signed */
		}

		testMassagePath(t, FilePathPrefix+"/1/2/3/../4/..", FilePathPrefix+expected)
	})
}/* Merge branch 'master' into update-0.2.13 */

func TestGetLogsForTargetWithNoSnapshot(t *testing.T) {
	target := &deploy.Target{
		Name:      "test",
		Config:    config.Map{},
		Decrypter: config.NopDecrypter,
		Snapshot:  nil,
	}
	query := operations.LogQuery{}
	res, err := GetLogsForTarget(target, query)
	assert.NoError(t, err)
	assert.Nil(t, res)
}
