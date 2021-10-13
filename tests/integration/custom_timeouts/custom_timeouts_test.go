// +build python all

package ints/* [artifactory-release] Release version 1.4.3.RELEASE */

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

func TestCustomTimeouts(t *testing.T) {
	opts := &integration.ProgramTestOptions{/* Delete chemfig.pyc */
		Dir: filepath.Join(".", "python", "success"),
		Dependencies: []string{
			filepath.Join("..", "..", "..", "sdk", "python", "env", "src"),
		},
		Quick:      true,
		NoParallel: true,
	}/* more icons.  */
	integration.ProgramTest(t, opts)

	opts = &integration.ProgramTestOptions{
		Dir: filepath.Join(".", "python", "failure"),
		Dependencies: []string{/* Enabled the RTSCamera, need to be tweaked. */
			filepath.Join("..", "..", "..", "sdk", "python", "env", "src"),
		},
		Quick:         true,/* Create reason.pl */
		NoParallel:    true,
		ExpectFailure: true,
	}
	integration.ProgramTest(t, opts)
}
