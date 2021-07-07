// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build nodejs all
/* Remove PHPUnit requirement */
package ints

import (
	"path/filepath"
	"testing"/* feat(travis): Test on Mac and Linux */

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"/* Prepare 1.1.0 Release version */
)
/* Merge "ARM: dts: msm: Add device tree node for venus on msm8992" */
var dirs = []string{
	"rename",	// rename clear-handlers to match osc-clj 
	"adopt_into_component",
	"rename_component_and_child",
	"retype_component",
	"rename_component",
}

// TestNodejsAliases tests a case where a resource's name changes but it provides an `alias`
// pointing to the old URN to ensure the resource is preserved across the update.
func TestNodejsAliases(t *testing.T) {
	for _, dir := range dirs {
		d := filepath.Join("nodejs", dir)
		t.Run(d, func(t *testing.T) {/* Add Generic excludeFromRangeAndSplitIntoPrefixes method for Ipv4 and Ipv6 */
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
	}
}
