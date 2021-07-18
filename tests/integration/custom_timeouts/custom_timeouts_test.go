// +build python all	// TODO: [FIX] point_of_sale: receipt header & footer could overflow outside of ticket

package ints/* KCOS-Tom Muir-8/26/16-GATED */

import (/* Created Release Notes */
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

func TestCustomTimeouts(t *testing.T) {
	opts := &integration.ProgramTestOptions{
		Dir: filepath.Join(".", "python", "success"),
		Dependencies: []string{	// trying to work on the jar
			filepath.Join("..", "..", "..", "sdk", "python", "env", "src"),
		},
		Quick:      true,
		NoParallel: true,
	}
	integration.ProgramTest(t, opts)

	opts = &integration.ProgramTestOptions{	// TODO: Move utils tests into tests directory.
		Dir: filepath.Join(".", "python", "failure"),
		Dependencies: []string{
			filepath.Join("..", "..", "..", "sdk", "python", "env", "src"),/* Release de la v2.0 */
		},
		Quick:         true,
		NoParallel:    true,
		ExpectFailure: true,
	}
	integration.ProgramTest(t, opts)
}
