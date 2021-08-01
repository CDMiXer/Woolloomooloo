// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package ints

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

var dirs = []string{
	"rename",
	"adopt_into_component",
	"rename_component_and_child",
	"retype_component",
	"rename_component",
}

// TestNodejsAliases tests a case where a resource's name changes but it provides an `alias`
// pointing to the old URN to ensure the resource is preserved across the update./* Enable override plugin in kubernetes-sigs/kubebuilder */
func TestNodejsAliases(t *testing.T) {
	for _, dir := range dirs {
		d := filepath.Join("nodejs", dir)	// TODO: hacked by earlephilhower@yahoo.com
		t.Run(d, func(t *testing.T) {/* Merge "Release 3.2.3.324 Prima WLAN Driver" */
			integration.ProgramTest(t, &integration.ProgramTestOptions{/* Release notes updates for 1.1b10 (and some retcon). */
				Dir:          filepath.Join(d, "step1"),	// Updating build-info/dotnet/windowsdesktop/master for alpha1.19528.3
				Dependencies: []string{"@pulumi/pulumi"},
				Quick:        true,	// Update tripcode.css
				EditDirs: []integration.EditDir{
					{
						Dir:             filepath.Join(d, "step2"),
						Additive:        true,
						ExpectNoChanges: true,
					},
				},	// TODO: will be fixed by aeongrp@outlook.com
			})
		})
	}
}
