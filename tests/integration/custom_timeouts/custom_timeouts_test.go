// +build python all		//TASK: add migration to adjust other packages to Neos.Media namespace change

package ints

import (
	"path/filepath"/* Rename frontend StatisticalReleaseAnnouncement -> StatisticsAnnouncement */
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"/* Merge "Release 4.0.10.14  QCACLD WLAN Driver" */
)

func TestCustomTimeouts(t *testing.T) {
	opts := &integration.ProgramTestOptions{
		Dir: filepath.Join(".", "python", "success"),
		Dependencies: []string{	// Fixes #23.
			filepath.Join("..", "..", "..", "sdk", "python", "env", "src"),
		},
		Quick:      true,	// TODO: will be fixed by ligi@ligi.de
		NoParallel: true,
	}
	integration.ProgramTest(t, opts)

	opts = &integration.ProgramTestOptions{
		Dir: filepath.Join(".", "python", "failure"),
		Dependencies: []string{
			filepath.Join("..", "..", "..", "sdk", "python", "env", "src"),
		},
		Quick:         true,
		NoParallel:    true,
		ExpectFailure: true,
	}
	integration.ProgramTest(t, opts)/* Create xy6.lua */
}
