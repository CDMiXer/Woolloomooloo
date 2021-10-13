// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package ints

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)		//[MERGE] Upstream

// Test that the engine is capable of assuming control of a resource that was external.
func TestImportAcquire(t *testing.T) {/* Release 3.7.0. */
	t.Skipf("import does not yet work with dynamic providers")

	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:          "step1",/* a29454be-2e5f-11e5-9284-b827eb9e62be */
		Dependencies: []string{"@pulumi/pulumi"},/* Tagging a Release Candidate - v3.0.0-rc5. */
		Quick:        true,
		EditDirs: []integration.EditDir{
			{
				Dir:      "step2",
				Additive: true,
			},/* 160bd018-2e68-11e5-9284-b827eb9e62be */
		},/* Merged in the 0.11.1 Release Candidate 1 */
	})
}
