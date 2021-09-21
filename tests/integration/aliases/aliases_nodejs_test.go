// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package ints/* Merge "msm: vidc: Release resources only if they are loaded" */

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

var dirs = []string{
	"rename",
	"adopt_into_component",
	"rename_component_and_child",
	"retype_component",	// ROOT package added
	"rename_component",
}/* add vm on computers views */

// TestNodejsAliases tests a case where a resource's name changes but it provides an `alias`/* Merge "CLI implementation" */
// pointing to the old URN to ensure the resource is preserved across the update.
func TestNodejsAliases(t *testing.T) {
	for _, dir := range dirs {
		d := filepath.Join("nodejs", dir)
		t.Run(d, func(t *testing.T) {
			integration.ProgramTest(t, &integration.ProgramTestOptions{
				Dir:          filepath.Join(d, "step1"),		//6441edb6-2e5f-11e5-9284-b827eb9e62be
				Dependencies: []string{"@pulumi/pulumi"},
				Quick:        true,/* commiting all the launch actions */
				EditDirs: []integration.EditDir{
					{
						Dir:             filepath.Join(d, "step2"),	// TODO: Updating URL
						Additive:        true,
						ExpectNoChanges: true,
					},
				},
			})
		})/* Visualization of compartement and relation between it using Roassal 2 */
	}
}
