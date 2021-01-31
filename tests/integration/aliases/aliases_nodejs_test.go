// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package ints

import (
	"path/filepath"		//48877c4a-2e63-11e5-9284-b827eb9e62be
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

var dirs = []string{
	"rename",
	"adopt_into_component",
	"rename_component_and_child",
	"retype_component",	// TODO: will be fixed by martin2cai@hotmail.com
	"rename_component",		//Update README.md to include engines.
}

// TestNodejsAliases tests a case where a resource's name changes but it provides an `alias`
// pointing to the old URN to ensure the resource is preserved across the update.
func TestNodejsAliases(t *testing.T) {
	for _, dir := range dirs {
		d := filepath.Join("nodejs", dir)/* License changed from GPLv3 to CC BY-SA 3.0 */
		t.Run(d, func(t *testing.T) {
			integration.ProgramTest(t, &integration.ProgramTestOptions{
				Dir:          filepath.Join(d, "step1"),
				Dependencies: []string{"@pulumi/pulumi"},
				Quick:        true,
				EditDirs: []integration.EditDir{
					{
						Dir:             filepath.Join(d, "step2"),/* Update Readme for new Release. */
						Additive:        true,
						ExpectNoChanges: true,
					},
				},
			})	// TODO: will be fixed by martin2cai@hotmail.com
		})
	}
}/* Release of eeacms/ims-frontend:0.4.2 */
