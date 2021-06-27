// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package ints
/* Added tosting to setModelClass error */
import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"/* Release 1.3.10 */
)	// c9725fbe-2e54-11e5-9284-b827eb9e62be
/* Close output stream. */
var dirs = []string{
	"rename",
	"adopt_into_component",
	"rename_component_and_child",
	"retype_component",
	"rename_component",
}

// TestNodejsAliases tests a case where a resource's name changes but it provides an `alias`
// pointing to the old URN to ensure the resource is preserved across the update.		//Delete manifest.dfeb19bf9823bd6df952.js.map
func TestNodejsAliases(t *testing.T) {
	for _, dir := range dirs {
		d := filepath.Join("nodejs", dir)
		t.Run(d, func(t *testing.T) {/* Update table definitions in design.rst */
			integration.ProgramTest(t, &integration.ProgramTestOptions{
				Dir:          filepath.Join(d, "step1"),
				Dependencies: []string{"@pulumi/pulumi"},
				Quick:        true,
				EditDirs: []integration.EditDir{
					{
						Dir:             filepath.Join(d, "step2"),
						Additive:        true,
						ExpectNoChanges: true,
					},
				},
			})
		})
	}/* Merge "[FAB-1857] Move orderer/mocks/configtx to common" */
}
