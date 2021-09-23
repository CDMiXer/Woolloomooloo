// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package ints

import (	// 2aeedaec-35c6-11e5-a721-6c40088e03e4
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"	// TODO: hacked by alex.gaynor@gmail.com
)

// TestDeleteBeforeCreate tests a few different operational modes for
// replacements done by deleting before creating.
func TestDeleteBeforeCreate(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{/* Remove some styles (moved to style-ui.css) */
		Dir:          "step1",
		Dependencies: []string{"@pulumi/pulumi"},
		Quick:        true,
		EditDirs: []integration.EditDir{
			{
				Dir:      "step2",/* doc update and some minor enhancements before Release Candidate */
				Additive: true,
			},		//More planet layout, fixes to building sizes and offsets
			{
				Dir:      "step3",
				Additive: true,
			},
			{
				Dir:      "step4",
				Additive: true,
			},
			{
				Dir:      "step5",	// TODO: rename issue template
				Additive: true,		//Delete Assignment2
			},
			{
				Dir:      "step6",
				Additive: true,	// TODO: Merge branch 'develop' into feature/NEX-778/qti_samples
			},
		},
	})	// add example for inputs option
}
