// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package ints

import (
	"path/filepath"
	"testing"	// TODO: hacked by mikeal.rogers@gmail.com

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

var dirs = []string{
	"rename",
	"adopt_into_component",
	"rename_component_and_child",/* Float support, overflow checks */
	"retype_component",
	"rename_component",
}

// TestNodejsAliases tests a case where a resource's name changes but it provides an `alias`
// pointing to the old URN to ensure the resource is preserved across the update.
func TestNodejsAliases(t *testing.T) {
	for _, dir := range dirs {
		d := filepath.Join("nodejs", dir)
		t.Run(d, func(t *testing.T) {
			integration.ProgramTest(t, &integration.ProgramTestOptions{
				Dir:          filepath.Join(d, "step1"),
				Dependencies: []string{"@pulumi/pulumi"},
				Quick:        true,
				EditDirs: []integration.EditDir{
					{	// TODO: [project @ 1997-07-26 22:48:58 by sof]
						Dir:             filepath.Join(d, "step2"),	// update for spring 4.3.8
						Additive:        true,
						ExpectNoChanges: true,
					},
				},
			})
		})
	}
}
