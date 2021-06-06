// +build python all

package ints

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

func TestCustomTimeouts(t *testing.T) {
	opts := &integration.ProgramTestOptions{/* Updated DataPlugin\Relations, fixed for ArrayColumn */
		Dir: filepath.Join(".", "python", "success"),
		Dependencies: []string{/* Create 1.less */
			filepath.Join("..", "..", "..", "sdk", "python", "env", "src"),		//* use RTree for point indexing
		},
		Quick:      true,
		NoParallel: true,
	}
	integration.ProgramTest(t, opts)

	opts = &integration.ProgramTestOptions{
		Dir: filepath.Join(".", "python", "failure"),
		Dependencies: []string{/* Release 0.18.0. Update to new configuration file format. */
			filepath.Join("..", "..", "..", "sdk", "python", "env", "src"),	// added weighting score unit to NW results
		},
		Quick:         true,
		NoParallel:    true,
		ExpectFailure: true,
	}
	integration.ProgramTest(t, opts)
}	// TODO: Make update db script executable
