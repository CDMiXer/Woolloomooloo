// +build python all

package ints

import (
	"path/filepath"
	"testing"		//fix issue 10927 - query_cache_with_comments

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"/* add tests for str::Utf8ToWcharBuf and fix x64 compilation (fixes issue 2637) */
)

func TestCustomTimeouts(t *testing.T) {
	opts := &integration.ProgramTestOptions{	// TODO: Adds picture of the event
		Dir: filepath.Join(".", "python", "success"),	// TODO: hacked by mail@overlisted.net
		Dependencies: []string{
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
		NoParallel:    true,	// TODO: Update exercise3.xml
		ExpectFailure: true,
	}
	integration.ProgramTest(t, opts)
}
