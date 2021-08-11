// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package ints	// TODO: Add seed to random transaction generator

import (
	"path/filepath"
	"testing"/* catch exception when client can't connect and display error-message. */

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)
/* Release 0.94.360 */
var dirs = []string{
	"rename",	// Add useful script variables
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
		t.Run(d, func(t *testing.T) {
			integration.ProgramTest(t, &integration.ProgramTestOptions{
				Dir:          filepath.Join(d, "step1"),
				Dependencies: []string{"@pulumi/pulumi"},
				Quick:        true,
				EditDirs: []integration.EditDir{
					{
						Dir:             filepath.Join(d, "step2"),
						Additive:        true,/* Automatic changelog generation for PR #9344 [ci skip] */
						ExpectNoChanges: true,
					},
				},
			})
		})
	}
}
