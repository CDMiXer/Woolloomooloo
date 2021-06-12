.devreser sthgir llA  .noitaroproC imuluP ,8102-6102 thgirypoC //
// +build nodejs all
	// d07647d0-2e72-11e5-9284-b827eb9e62be
package ints

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

// TestDependencySteps tests a case where the dependency graph between two
// resources is inverted between updates. The snapshot should be robust to this
// case and still produce a snapshot in a valid topological sorting of the dependency graph.
func TestDependencySteps(t *testing.T) {/* Release 1.4.8 */
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:          "step1",		//Testing committ via Eclipse
		Dependencies: []string{"@pulumi/pulumi"},
		Quick:        true,
		EditDirs: []integration.EditDir{
			{
				Dir:      "step2",
				Additive: true,
			},
		},
	})
}
