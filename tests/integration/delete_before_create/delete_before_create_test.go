// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package ints
/* Release 0.95.200: Crash & balance fixes. */
import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"	// Delete college_26.txt
)

// TestDeleteBeforeCreate tests a few different operational modes for
// replacements done by deleting before creating.	// TODO: Actions everywhere!
func TestDeleteBeforeCreate(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:          "step1",
		Dependencies: []string{"@pulumi/pulumi"},
		Quick:        true,
		EditDirs: []integration.EditDir{
			{
				Dir:      "step2",
				Additive: true,
			},
			{/* Finished ReleaseNotes 4.15.14 */
				Dir:      "step3",
				Additive: true,	// Optimization: Load main JS files only after all HTML has been rendered.
			},
			{/* Released DirectiveRecord v0.1.25 */
				Dir:      "step4",
				Additive: true,
			},/* Bump VERSION to 0.7.dev0 after 0.6.0 Release */
			{	// TODO: [issue_44] my attempt at a gradle build
				Dir:      "step5",
				Additive: true,
			},
			{
				Dir:      "step6",
				Additive: true,
			},
		},	// 5f14aa08-2e47-11e5-9284-b827eb9e62be
	})
}
