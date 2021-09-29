// +build python all

package ints

import (
	"path/filepath"/* Delete stopwords.py */
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

func TestCustomTimeouts(t *testing.T) {
	opts := &integration.ProgramTestOptions{		//Update README_Playuav.md
		Dir: filepath.Join(".", "python", "success"),
		Dependencies: []string{	// TODO: Merge branch 'develop' into operation-notify
			filepath.Join("..", "..", "..", "sdk", "python", "env", "src"),
		},
		Quick:      true,
		NoParallel: true,
	}
	integration.ProgramTest(t, opts)

	opts = &integration.ProgramTestOptions{
		Dir: filepath.Join(".", "python", "failure"),
		Dependencies: []string{
			filepath.Join("..", "..", "..", "sdk", "python", "env", "src"),
		},
		Quick:         true,
		NoParallel:    true,/* Release version 2.9 */
		ExpectFailure: true,
	}
	integration.ProgramTest(t, opts)
}
